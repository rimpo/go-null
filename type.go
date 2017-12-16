package main

import (
	"bytes"
	"fmt"
	"log"
)

type Type struct {
	name      string     // name of the type
	baseName  string     // built in type ie. int, string, float64, bool
	values    ValueSlice // Accumulator for constant value of this type.
	pkgPath   string     // path of the package
	pkgName   string     //source package name; package where type is declared
	numerical bool       // numerical is true for float64 and int
}

var coreTypes = []Type{
	Type{
		name:      "Int",
		baseName:  "int",
		values:    []Value{},
		pkgName:   "",
		numerical: true,
	},
	Type{
		name:      "String",
		baseName:  "string",
		values:    []Value{},
		pkgName:   "",
		numerical: false,
	},
	Type{
		name:      "Float64",
		baseName:  "float64",
		values:    []Value{},
		pkgName:   "",
		numerical: true,
	},
	Type{
		name:      "Bool",
		baseName:  "bool",
		values:    []Value{},
		pkgName:   "",
		numerical: false,
	},
}

func (typ *Type) IsNumerical() bool {
	return typ.baseName == "int" || typ.baseName == "float64"
}

func (typ *Type) IsCore() bool {
	for i, _ := range coreTypes {
		if typ.name == coreTypes[i].name {
			return true
		}
	}
	return false
}

func (typ *Type) IsEnum() bool {
	return len(typ.values) > 0
}

func (typ *Type) Check() bool {
	dupVal := make(map[string]bool)
	//for all types check for duplicate declaration
	for _, v := range typ.values {
		_, ok := dupVal[v.str]
		if !ok {
			dupVal[v.str] = true
		} else {
			fmt.Println("Error: Duplicate enum found in type:", typ.name, " name:", v.name, dupVal, typ.values)
			return true
		}
	}
	//for numerical - we need text representation in comments
	if typ.numerical {
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

func (typ *Type) generateHeaderCode() string {
	var buf bytes.Buffer
	var err error
	if typ.IsCore() {
		return coreHeaderCode
	} else {
		err = templateBasicHeaderCode.Execute(&buf, struct {
			ImportLib string
		}{
			"\"" + typ.pkgPath + "/" + typ.pkgName + "\"",
		})
	}

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return ""
	}
	return buf.String()
}

func (typ *Type) generateBodyCode() string {
	var buf bytes.Buffer
	var builtInTypeName string

	if typ.IsCore() {
		builtInTypeName = typ.baseName
	} else {
		builtInTypeName = typ.pkgName + "." + typ.name
	}

	err := templateBasicBodyCode.Execute(&buf, struct {
		TypeName        string
		BuiltInTypeName string
	}{
		typ.name,
		builtInTypeName,
	})

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return ""
	}
	return buf.String()
}

func (typ *Type) generateIsValueCode() string {
	if !typ.IsEnum() || typ.IsCore() {
		return ""
	}
	var buf bytes.Buffer
	var err error
	if len(typ.values) > 10 {
		err = templateIsValueMapCode.Execute(&buf, struct {
			TypeName        string
			BuiltInTypeName string
			MapValues       string
		}{
			typ.name,
			typ.pkgName + "." + typ.name,
			typ.values.generateMapValue(typ.pkgName),
		})
	} else {
		err = templateIsValueSwitchCode.Execute(&buf, struct {
			TypeName        string
			BuiltInTypeName string
			SwitchCode      string
		}{
			typ.name,
			typ.pkgName + "." + typ.name,
			typ.values.generateSwitchCode(typ.pkgName),
		})
	}

	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return ""
	}
	return buf.String()
}

func (typ *Type) generateMarshalCode() string {
	var buf bytes.Buffer
	var err error
	if typ.IsEnum() && typ.IsNumerical() {
		//enum type
		err = templateNumericalMapMarshalCode.Execute(&buf, struct {
			TypeName        string
			BuiltInTypeName string
			MapValues       string
		}{
			typ.name,
			typ.pkgName + "." + typ.name,
			typ.values.generateMapKeyValue(typ.pkgName),
		})
	} else {
		err = templateBasicMarshalCode.Execute(&buf, struct {
			TypeName string
		}{
			typ.name,
		})
	}
	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return ""
	}

	return buf.String()
}

func (typ *Type) generateCode() bytes.Buffer {
	//writing in the null_types.go map
	//g.appendInAllType(typ.name)

	var buf bytes.Buffer
	buf.WriteString(typ.generateHeaderCode())
	buf.WriteString(typ.generateBodyCode())
	buf.WriteString(typ.generateIsValueCode())
	buf.WriteString(typ.generateMarshalCode())

	/*nullDir := fmt.Sprintf("%s/null", g.outputDir)
	typeFilePath := nullDir + "/" + strings.ToLower(typ.name) + ".go"

	if err := ioutil.WriteFile(typeFilePath, g.format(&buf), 0644); err != nil {
		fmt.Println("Warning: failed to write file:", typeFilePath, " err:", err)
	}*/
	return buf
}
