package main

import (
	"bytes"
	"log"
	"text/template"
)

//Type code generator functions

const allTypesHeaderCode = `
package null
//Auto-generated code; DONT EDIT THIS CODE
var (
	AllTypes = map[string]bool {
	}
)
`

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
	if typ.IsSpecialMapType() {
		typ.templateType = EnumMapType
	} else if typ.IsCore() {
		typ.templateType = CoreType
	} else if !typ.IsEnum() {
		typ.templateType = NonEnumType
	} else if typ.IsEnum() && typ.IsNumerical() {
		typ.templateType = EnumNumType
	} else if typ.IsEnum() && !typ.IsNumerical() {
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
	if typ.baseName == "TypeMapInt" {
		result += "[int(val)]\n"
	} else if typ.baseName == "TypeMapString" {
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
	if typ.IsStringType() {
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
