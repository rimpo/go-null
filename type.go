package main

import (
	"fmt"
)

type Type struct {
	name         string     // name of the type
	baseName     string     // built in type ie. int, string, float64, bool
	values       ValueSlice // Accumulator for constant value of this type.
	pkgPath      string     // path of the package
	pkgName      string     //source package name; package where type is declared
	numerical    bool       // numerical is true for float64 and int
	templateType TemplateType
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

func (typ *Type) IsSpecialMapType() bool {
	return typ.baseName == "TypeMapInt" || typ.baseName == "TypeMapString"
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
