package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
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
	fmt.Fprintf(os.Stderr, "\thttp://github.com/rimpo/go-null\n")
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
	log.SetPrefix("go-null: ")
	flag.Usage = Usage
	flag.Parse()

	// process whole package in current directory.
	args := []string{"."}

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

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	types      []Type // Accumulator for all alias types (which is a built in type)
	trimPrefix string
}

func (f *File) setPkgPath(pkgPath string) {
	for i, _ := range f.types {
		f.types[i].pkgPath = pkgPath
	}
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

// generate produces the String method for the named type.
func (g *Generator) generate() {
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.types = nil

		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)

			file.setPkgPath(g.pkgPath)

			for i, _ := range file.types {
				file.types[i].Check()
				g.generateCode(&file.types[i])

			}
		}
	}

	for i, _ := range coreTypes {
		g.generateCode(&coreTypes[i])
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

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if ok {
		if decl.Tok == token.TYPE {
			for _, spec := range decl.Specs {
				tspec, ok := spec.(*ast.TypeSpec)
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
	//fmt.Fprintf(&buf, GetJsonxCode(), g.pkgPath)

	var templateJsonxCode = template.Must(template.New("templateJsonxCode").Parse(GetJsonxCode()))

	err := templateJsonxCode.Execute(&buf, struct {
		PackagePath string
	}{
		g.pkgPath,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return err
	}

	jsonxDir := fmt.Sprintf("%s/jsonx", g.outputDir)
	if err := ioutil.WriteFile(jsonxDir+"/jsonx.go", g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write jsonx.go file:", err)
	}

	buf.Reset()

	var templateJsonxTestCode = template.Must(template.New("templateJsonxTestCode").Parse(GetJsonxTestCode()))

	err = templateJsonxTestCode.Execute(&buf, struct {
		PackagePath string
	}{
		g.pkgPath,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return err
	}

	if err := ioutil.WriteFile(jsonxDir+"/jsonx_test.go", g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write jsonx_test.go file:", err)
	}
	return nil
}

func (g *Generator) generateCode(typ *Type) {
	//writing in the null_types.go map
	g.appendInAllType(typ.name)

	buf := typ.generateCode()

	nullDir := fmt.Sprintf("%s/null", g.outputDir)
	typeFilePath := nullDir + "/" + strings.ToLower(typ.name) + ".go"

	if err := ioutil.WriteFile(typeFilePath, g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write file:", typeFilePath, " err:", err)
	}
}
