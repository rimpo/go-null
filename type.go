package main

import (
	"bytes"
	"fmt"
	exact "go/constant"
	"log"
	"text/template"
)

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

type Type struct {
	name         string  // name of the type
	baseName     string  // built in type ie. int, string, float64, bool
	values       []Value // Accumulator for constant value of this type.
	pkgPath      string  // path of the package
	pkgName      string  //source package name; package where type is declared
	templateType TemplateType
}

var coreTypes = []Type{
	Type{
		name:     "Int",
		baseName: "int",
		values:   []Value{},
		pkgName:  "",
	},
	Type{
		name:     "String",
		baseName: "string",
		values:   []Value{},
		pkgName:  "",
	},
	Type{
		name:     "Float64",
		baseName: "float64",
		values:   []Value{},
		pkgName:  "",
	},
	Type{
		name:     "Bool",
		baseName: "bool",
		values:   []Value{},
		pkgName:  "",
	},
}

//Special Type: TypeMapInt and TypeMapString
//Should be rarely used.
//New type declaration using special type should only used when we have very big enum set and the enum const are not needed for business logic.
//One example of this type would be Country Code. We only need the code string mapping of huge code list.

type typeID string

const (
	stringID        typeID = "string"
	intID                  = "int"           //numerical types
	float64ID              = "float64"       //numerical types
	typeMapIntID           = "TypeMapInt"    //special types (hack): enum types with no const values but with a map exported with name Map<type name>IDToText
	typeMapStringID        = "TypeMapString" //special types (hack): enum types with no const values but with a map exported with name Map<type name>IDToText
)

func (typ *Type) isSpecialMapType() bool {
	return typ.baseName == string(typeMapIntID) || typ.baseName == string(typeMapStringID)
}

func (typ *Type) isNumerical() bool {
	return typ.baseName == string(intID) || typ.baseName == string(float64ID)
}

func (typ *Type) isStringType() bool {
	return typ.baseName == string(stringID) || typ.baseName == string(typeMapStringID)
}

func (typ *Type) isCore() bool {
	for i, _ := range coreTypes {
		if typ.name == coreTypes[i].name {
			return true
		}
	}
	return false
}

func (typ *Type) isEnum() bool {
	return len(typ.values) > 0
}

func (typ *Type) assert() error {
	dupVal := make(map[string]bool)
	//for all types check for duplicate declaration
	for _, v := range typ.values {
		_, ok := dupVal[v.str]
		if !ok {
			dupVal[v.str] = true
		} else {
			return fmt.Errorf("Error: Duplicate enum found in type:%v name:%v dup value:%v values:%v", typ.name, v.name, dupVal, typ.values)
		}
	}
	//for numerical - we need text representation in comments
	if typ.isNumerical() {
		dupComment := make(map[string]bool)
		for _, v := range typ.values {
			if v.comment == "" {
				//return fmt.Errorf("Error: No comment found in int/float64 type enum type:%v name:%v", typ.name, v.name)
			}
			_, ok := dupComment[v.comment]
			if !ok {
				dupComment[v.comment] = true
			} else {
				return fmt.Errorf("Error: Duplicate enum comment found in type:%v name:%v", typ.name, v.name)
			}
		}
	}
	return nil
}

var templateCore = template.Must(template.New("templateCoreCodeString").Parse(templateCoreCodeString))
var templateNonEnum = template.Must(template.New("templateNonEnumCodeString").Parse(templateNonEnumCodeString))
var templateEnumNum = template.Must(template.New("templateEnumNumCodeString").Parse(templateEnumNumCodeString))
var templateEnumNonNum = template.Must(template.New("templateEnumNonNumCodeString").Parse(templateEnumNonNumCodeString))
var templateEnumMap = template.Must(template.New("templateEnumMapCodeString").Parse(templateEnumMapCodeString))

type TemplateType int

const (
	InvalidType    TemplateType = 0
	CoreType                    = 1
	NonEnumType                 = 2
	EnumNumType                 = 3
	EnumNonNumType              = 4
	EnumMapType                 = 5
)

func (typ *Type) setTemplateType() {
	if typ.isSpecialMapType() {
		typ.templateType = EnumMapType
	} else if typ.isCore() {
		typ.templateType = CoreType
	} else if !typ.isEnum() {
		typ.templateType = NonEnumType
	} else if typ.isEnum() && typ.isNumerical() {
		typ.templateType = EnumNumType
	} else if typ.isEnum() && !typ.isNumerical() {
		typ.templateType = EnumNonNumType
	} else {
		typ.templateType = InvalidType
	}
}

func (typ *Type) getTemplate() *template.Template {
	switch typ.templateType {
	case CoreType:
		return templateCore
	case NonEnumType:
		return templateNonEnum
	case EnumNumType:
		return templateEnumNum
	case EnumNonNumType:
		return templateEnumNonNum
	case EnumMapType:
		return templateEnumMap
	}
	return nil
}

func (typ *Type) getImportLib() string {
	switch typ.templateType {
	case CoreType:
		return ""
	default:
		return "\"" + typ.pkgPath + "/" + typ.pkgName + "\""
	}
}

func (typ *Type) getSourceType() string {
	switch typ.templateType {
	case CoreType:
		return typ.baseName
	default:
		return typ.pkgName + "." + typ.name
	}
}

func (typ *Type) getMapValuesCode() string {
	if typ.templateType == EnumMapType {
		return ""
	}
	result := ""
	for _, v := range typ.values {
		result += "\t" + typ.pkgName + "." + v.name + ": "
		switch typ.templateType {
		case EnumNumType:
			result += "\"" + v.comment + "\",\n"
		default:
			result += v.str + ",\n"
		}
	}
	return result
}

func (typ *Type) getLookupSwitchCode() string {
	result := "switch val {\n"
	for _, v := range typ.values {
		result += "\tcase " + typ.pkgName + "." + v.name + ":\n"
		switch typ.templateType {
		case EnumNumType:
			result += "\t\treturn \"" + v.comment + "\", true\n"
		default:
			result += "\t\treturn " + v.str + ", true\n"
		}
	}
	result += "\tdefault:\n\treturn \"\", false}\t"
	return result
}

func (typ *Type) getLookupMapCode() string {
	result := "\tres, ok := map" + typ.name + "IDToText[val]\n"
	result += "\treturn res, ok\n"
	return result
}

func (typ *Type) getLookupSpecialMapCode() string {
	result := "\tres, ok := " + typ.pkgName + ".Map" + typ.name + "IDToText"
	if typ.baseName == string(typeMapIntID) {
		result += "[int(val)]\n"
	} else if typ.baseName == string(typeMapStringID) {
		result += "[string(val)]\n"
	}
	result += "\treturn res, ok\n"
	return result
}

func (typ *Type) getLookupCode() string {
	if typ.templateType == EnumMapType {
		return typ.getLookupSpecialMapCode()
	}

	if len(typ.values) >= 10 {
		return typ.getLookupMapCode()
	} else {
		return typ.getLookupSwitchCode()
	}
}

func (typ *Type) getIsEmptyCode() string {
	result := ""
	if typ.isStringType() {
		return "func (t *" + typ.name + ") IsEmpty() bool {\n\treturn t.IsNull() || len(string(t.val)) == 0\n}"
	}
	return result
}

func (typ *Type) generateCode() bytes.Buffer {
	var buf bytes.Buffer
	err := typ.getTemplate().Execute(&buf, struct {
		ImportLib   string
		TypeName    string
		SourceType  string
		MapValues   string
		LookupCode  string
		IsEmptyCode string
	}{
		typ.getImportLib(),
		typ.name,
		typ.getSourceType(),
		typ.getMapValuesCode(),
		typ.getLookupCode(),
		typ.getIsEmptyCode(),
	})
	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return buf
	}
	return buf
}
