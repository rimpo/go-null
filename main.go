package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	exact "go/constant"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	packagePath = flag.String("package", "", "package path in the project; example: github.com/rimpo/profile-api/pkg")
	outputDir   = flag.String("output", "..", "output directory (relative path from go:generate command file place); default parent directory")
	trimprefix  = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "\thttp://godoc.org/golang.org/x/tools/cmd/make-null-type\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func mkDir(dir string) {
	if err := os.MkdirAll(dir, 0700); err != nil {
		fmt.Println("Warning: failed to create directory:", dir, err)
	}

}

func main() {
	log.SetFlags(0)
	log.SetPrefix("make-null-type: ")
	flag.Usage = Usage
	flag.Parse()

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	// Parse the package once.
	var dir string
	g := Generator{trimPrefix: *trimprefix}
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
		g.parsePackageDir(args[0])
	} else {
		dir = filepath.Dir(args[0])
		g.parsePackageFiles(args)
	}

	g.outputDir = *outputDir

	g.pkgPath = *packagePath

	mkDir(g.outputDir + "/jsonx")
	mkDir(g.outputDir + "/null")

	g.createJsonxLib()

	g.generate()

	fmt.Println(dir)
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.

	trimPrefix string
	pkgPath    string //pkg path where all the packages of projects are kept
	outputDir  string //Path where the null, jsonx directory will be created
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

type Type struct {
	name     string
	baseName string
	values   []Value // Accumulator for constant values of that type.
	pkgName  string
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	types      []Type // Accumulator for all alias types (which is a built in type)
	trimPrefix string
}

type Package struct {
	dir      string
	name     string
	defs     map[*ast.Ident]types.Object
	files    []*File
	typesPkg *types.Package
}

// parsePackageDir parses the package residing in the directory.
func (g *Generator) parsePackageDir(directory string) {
	pkg, err := build.Default.ImportDir(directory, 0)
	if err != nil {
		log.Fatalf("cannot process directory %s: %s", directory, err)
	}
	var names []string
	names = append(names, pkg.GoFiles...)
	names = append(names, pkg.CgoFiles...)
	// TODO: Need to think about constants in test files. Maybe write type_string_test.go
	// in a separate pass? For later.
	// names = append(names, pkg.TestGoFiles...) // These are also in the "foo" package.
	names = append(names, pkg.SFiles...)
	names = prefixDirectory(directory, names)
	g.parsePackage(directory, names, nil)
}

// parsePackageFiles parses the package occupying the named files.
func (g *Generator) parsePackageFiles(names []string) {
	g.parsePackage(".", names, nil)
}

// prefixDirectory places the directory name on the beginning of each name in the list.
func prefixDirectory(directory string, names []string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}

// parsePackage analyzes the single package constructed from the named files.
// If text is non-nil, it is a string to be used instead of the content of the file,
// to be used for testing. parsePackage exits if there is an error.
func (g *Generator) parsePackage(directory string, names []string, text interface{}) {
	var files []*File
	var astFiles []*ast.File
	g.pkg = new(Package)
	fs := token.NewFileSet()
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		parsedFile, err := parser.ParseFile(fs, name, text, parser.ParseComments)
		if err != nil {
			log.Fatalf("parsing package: %s: %s", name, err)
		}
		astFiles = append(astFiles, parsedFile)
		files = append(files, &File{
			file:       parsedFile,
			pkg:        g.pkg,
			trimPrefix: g.trimPrefix,
		})
	}
	if len(astFiles) == 0 {
		log.Fatalf("%s: no buildable Go files", directory)
	}
	g.pkg.name = astFiles[0].Name.Name
	g.pkg.files = files
	g.pkg.dir = directory
	// Type check the package.
	g.pkg.check(fs, astFiles)
}

// check type-checks the package. The package must be OK to proceed.
func (pkg *Package) check(fs *token.FileSet, astFiles []*ast.File) {
	pkg.defs = make(map[*ast.Ident]types.Object)
	config := types.Config{Importer: defaultImporter(), FakeImportC: true}
	info := &types.Info{
		Defs: pkg.defs,
	}
	typesPkg, err := config.Check(pkg.dir, fs, astFiles, info)
	if err != nil {
		log.Fatalf("checking package: %s", err)
	}
	pkg.typesPkg = typesPkg
}

func (g *Generator) checkError(typ *Type) bool {
	dupVal := make(map[string]bool)
	for _, v := range typ.values {
		_, ok := dupVal[v.str]
		if !ok {
			dupVal[v.str] = true
		} else {
			fmt.Println("Error: Duplicate enum found in type:", typ.name, " name:", v.name, dupVal, typ.values)

			return true
		}
	}
	if typ.baseName == "int" || typ.baseName == "float64" {
		dupComment := make(map[string]bool)
		for _, v := range typ.values {
			if v.comment == "" {
				fmt.Println("Error: No comment found in int/float64 type enum type:", typ.name, " name:", v.name)
				return true
			}
			_, ok := dupComment[v.comment]
			if !ok {
				dupComment[v.comment] = true
			} else {
				fmt.Println("Error: Duplicate enum comment found in type:", typ.name, " name:", v.name)
				return true
			}
		}
	}
	return false
}

// generate produces the String method for the named type.
func (g *Generator) generate() {
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.types = nil

		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			//fmt.Println(g.pkg.name, file)
		}

		for i, _ := range file.types {
			//fmt.Println(file.types[i], "\n\n\n")
			g.checkError(&file.types[i])
			if len(file.types[i].values) == 0 {
				//Not a enum type
				g.createNullForNonEnums(&file.types[i])
			} else {
				switch file.types[i].baseName {
				case "string":
					g.createNullForStringEnum(&file.types[i])
				case "int":
					g.createNullForIntEnum(&file.types[i])
				default:
				}
			}
		}
	}

	var coreTypes []Type
	coreTypes = append(coreTypes, Type{
		name:     "Int",
		baseName: "int",
		values:   []Value{},
		pkgName:  "",
	})
	coreTypes = append(coreTypes, Type{
		name:     "String",
		baseName: "string",
		values:   []Value{},
		pkgName:  "",
	})
	coreTypes = append(coreTypes, Type{
		name:     "Float64",
		baseName: "float64",
		values:   []Value{},
		pkgName:  "",
	})
	coreTypes = append(coreTypes, Type{
		name:     "Bool",
		baseName: "bool",
		values:   []Value{},
		pkgName:  "",
	})
	for i, _ := range coreTypes {
		g.createNullForCoreType(&coreTypes[i])
	}

}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format(buf *bytes.Buffer) []byte {
	src, err := format.Source(buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return buf.Bytes()
	}
	return src
}

// Value represents a declared constant.
type Value struct {
	typ  string //type name
	name string // The name of the constant.
	// The value is stored as a bit pattern alone. The boolean tells us
	// whether to interpret it as an int64 or a uint64; the only place
	// this matters is when sorting.
	// Much of the time the str field is all we need; it is printed
	// by Value.String.
	value   uint64 // Will be converted to int64 when needed.
	signed  bool   // Whether the constant is a signed type.
	str     string // The string representation given by the "go/exact" package.
	comment string
	kind    exact.Kind
}

func (v *Value) String() string {
	return v.str
}

// byValue lets us sort the constants into increasing order.
// We take care in the Less method to sort in signed or unsigned order,
// as appropriate.
type byValue []Value

func (b byValue) Len() int      { return len(b) }
func (b byValue) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byValue) Less(i, j int) bool {
	if b[i].signed {
		return int64(b[i].value) < int64(b[j].value)
	}
	return b[i].value < b[j].value
}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if ok {
		if decl.Tok == token.TYPE {
			for _, spec := range decl.Specs {
				tspec, ok := spec.(*ast.TypeSpec) // Guaranteed to succeed as this is CONST.
				if ok {
					tbase, ok := tspec.Type.(*ast.Ident)
					if ok {
						//Accumulate type declaration
						t := Type{
							name:     tspec.Name.String(),
							baseName: tbase.String(),
							pkgName:  f.pkg.name,
						}
						f.types = append(f.types, t)
					}
				}

			}
		}
	}
	if !ok || decl.Tok != token.CONST {
		// We only care about const declarations.
		return true
	}
	// The name of the type of the constants we are declaring.
	// Can change if this is a multi-element declaration.
	typ := ""
	// Loop over the elements of the declaration. Each element is a ValueSpec:
	// a list of names possibly followed by a type, possibly followed by values.
	// If the type and value are both missing, we carry down the type (and value,
	// but the "go/types" package takes care of that).
	for _, spec := range decl.Specs {
		vspec := spec.(*ast.ValueSpec) // Guaranteed to succeed as this is CONST.
		if typ == "" && vspec.Type == nil && len(vspec.Values) > 0 {
			// "X = 1". With no type but a value, the constant is untyped.
			// Skip this vspec and reset the remembered type.
			typ = ""
			//fmt.Println("Skip the type", vspec.Type, vspec.Values)
			continue
		}
		if vspec.Type != nil {
			// "X T". We have a type. Remember it.
			ident, ok := vspec.Type.(*ast.Ident)
			if !ok {
				//fmt.Println("Skip vspec.Type != nil", vspec.Type, vspec.Values)
				continue
			}
			typ = ident.Name
		}

		// We now have a list of names (from one line of source code) all being
		// declared with the desired type.
		// Grab their names and actual values and store them in f.types[i].values.
		for _, name := range vspec.Names {
			if name.Name == "_" {
				continue
			}
			// This dance lets the type checker find the values for us. It's a
			// bit tricky: look up the object declared by the name, find its
			// types.Const, and extract its value.
			obj, ok := f.pkg.defs[name]
			if !ok {
				log.Fatalf("no value for constant %s", name)
			}
			//info := obj.Type().Underlying().(*types.Basic).Info()
			//if info&types.IsInteger == 0 {
			//	log.Fatalf("can't handle non-integer constant type %s", typ)
			//}
			value := obj.(*types.Const).Val() // Guaranteed to succeed as this is CONST.
			/*var u64 uint64
			var isUint bool
			if value.Kind() == exact.Int {
				log.Fatalf("can't happen: constant is not an integer %s", name)
				i64, isInt := exact.Int64Val(value)
				u64, isUint = exact.Uint64Val(value)
				if !isInt && !isUint {
					log.Fatalf("internal error: value of %s is not an integer: %s", name, value.String())
				}
				if !isInt {
					u64 = uint64(i64)
				}
			}*/
			v := Value{
				typ:  typ,
				name: name.Name,
				//value: u64,
				//signed: info&types.IsUnsigned == 0,
				kind:    value.Kind(),
				str:     value.String(),
				comment: strings.TrimRight(vspec.Comment.Text(), "\n"),
			}
			v.name = strings.TrimPrefix(v.name, f.trimPrefix)
			f.types[len(f.types)-1].values = append(f.types[len(f.types)-1].values, v)
			//fmt.Println("inside gen:", v)
		}
	}
	return false
}

const stringHeader = `
package %s

`

const stringImport = `
import (
	"log"
	"runtime/debug"
	%s
)
`

func (g *Generator) appendInAllType(typeName string) error {
	nullDir := fmt.Sprintf("%s/null", g.outputDir)
	data, err := ioutil.ReadFile(nullDir + "/all_types.go")
	var code string
	if err != nil {
		code = allTypesHeaderCode
	} else {
		code = string(data)
	}

	if strings.Contains(code, typeName) {
		//already available
		return nil
	}

	code = strings.Replace(code, "}", fmt.Sprintf("\t\"%s\": true,\n}", typeName), -1)

	var buf bytes.Buffer
	buf.WriteString(code)
	if err := ioutil.WriteFile(nullDir+"/all_types.go", g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write all_types.go file:", err)
	}
	return nil
}

func (g *Generator) createJsonxLib() error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, GetJsonxCode(), g.pkgPath)

	jsonxDir := fmt.Sprintf("%s/jsonx", g.outputDir)
	if err := ioutil.WriteFile(jsonxDir+"/jsonx.go", g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write jsonx.go file:", err)
	}

	buf.Reset()
	fmt.Fprintf(&buf, GetJsonxTestCode(), g.pkgPath)

	if err := ioutil.WriteFile(jsonxDir+"/jsonx_test.go", g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write jsonx_test.go file:", err)
	}
	return nil
}

func (g *Generator) createNullForStringEnum(typ *Type) {
	//writing in the null_types.go map
	g.appendInAllType(typ.name)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, basicHeaderCode, g.pkgPath+"/"+typ.pkgName)

	t := template.Must(template.New("null_string").Parse(basicBodyCode))

	err := t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	var isValueCode string
	if len(typ.values) > 5 {
		isValueCode = generateStringEnumMap_IsValue(typ)
	} else {
		isValueCode = generateStringEnumSwitch_IsValue(typ)
	}

	t = template.Must(template.New("is_value").Parse(isValueCode))
	err = t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	t = template.Must(template.New("marshal").Parse(basicMarshalCode))
	err = t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	nullDir := fmt.Sprintf("%s/null", g.outputDir)
	typeFilePath := nullDir + "/" + strings.ToLower(typ.name) + ".go"

	if err := ioutil.WriteFile(typeFilePath, g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write file:", typeFilePath, " err:", err)
	}
}

func (g *Generator) createNullForIntEnum(typ *Type) {
	//writing in the null_types.go map
	g.appendInAllType(typ.name)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, basicHeaderCode, g.pkgPath+"/"+typ.pkgName)
	t := template.Must(template.New("null_basic_body").Parse(basicBodyCode))

	err := t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	var isValueCode string
	if len(typ.values) > 10 {
		isValueCode = generateStringEnumMap_IsValue(typ)
	} else {
		isValueCode = generateStringEnumSwitch_IsValue(typ)
	}

	t = template.Must(template.New("is_value").Parse(isValueCode))
	err = t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	t = template.Must(template.New("marshal").Parse(generateIntMarshal(typ)))
	err = t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	nullDir := fmt.Sprintf("%s/null", g.outputDir)
	typeFilePath := nullDir + "/" + strings.ToLower(typ.name) + ".go"

	if err := ioutil.WriteFile(typeFilePath, g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write file:", typeFilePath, " err:", err)
	}
}

func (g *Generator) createNullForNonEnums(typ *Type) {
	//writing in the null_types.go map
	g.appendInAllType(typ.name)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, basicHeaderCode, g.pkgPath+"/"+typ.pkgName)
	t := template.Must(template.New("null_basic_body").Parse(basicBodyCode))

	err := t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	t = template.Must(template.New("marshal").Parse(basicMarshalCode))
	err = t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.pkgName + "." + typ.name,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	nullDir := fmt.Sprintf("%s/null", g.outputDir)
	typeFilePath := nullDir + "/" + strings.ToLower(typ.name) + ".go"

	if err := ioutil.WriteFile(typeFilePath, g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write file:", typeFilePath, " err:", err)
	}
}

func (g *Generator) createNullForCoreType(typ *Type) {
	//writing in the null_types.go map
	g.appendInAllType(typ.name)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, coreHeaderCode)
	t := template.Must(template.New("null_basic_body").Parse(basicBodyCode))

	err := t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.baseName,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}
	t = template.Must(template.New("marshal").Parse(basicMarshalCode))
	err = t.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		typ.baseName,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
	}

	nullDir := fmt.Sprintf("%s/null", g.outputDir)
	typeFilePath := nullDir + "/" + strings.ToLower(typ.name) + ".go"

	if err := ioutil.WriteFile(typeFilePath, g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write file:", typeFilePath, " err:", err)
	}
}
