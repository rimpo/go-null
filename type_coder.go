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

var templateCore = template.Must(template.New("templateCoreCode").Parse(templateCoreCodeString))
var templateNonEnum = template.Must(template.New("templateNonEnumCode").Parse(templateNonEnumCodeString))
var templateEnumNum = template.Must(template.New("templateEnumNumCodeString").Parse(templateEnumNumCodeString))
var templateEnumNonNum = template.Must(template.New("templateEnumNonNumCode").Parse(templateEnumNonNumCodeString))

type TemplateType int

const (
	InvalidType    TemplateType = 0
	CoreType                    = 1
	NonEnumType                 = 2
	EnumNumType                 = 3
	EnumNonNumType              = 4
)

func (typ *Type) setTemplateType() {
	if typ.IsCore() {
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

func (typ *Type) getLookupCode() string {
	if len(typ.values) >= 10 {
		return typ.getLookupMapCode()
	} else {
		return typ.getLookupSwitchCode()
	}
}

func (typ *Type) generateCode() bytes.Buffer {
	var buf bytes.Buffer
	err := typ.getTemplate().Execute(&buf, struct {
		ImportLib  string
		TypeName   string
		SourceType string
		MapValues  string
		LookupCode string
	}{
		typ.getImportLib(),
		typ.name,
		typ.getSourceType(),
		typ.getMapValuesCode(),
		typ.getLookupCode(),
	})
	if err != nil {
		log.Fatalf("Execution failed:%s", err)
		return buf
	}
	return buf
}
