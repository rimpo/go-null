package main

import (
	"text/template"
)

const basicHeaderCode = `
package null

import (
	"log"
    "runtime/debug"
	"encoding/json"

	{{.ImportLib}}
)

//Auto-generated code; DONT EDIT THIS CODE

`

const coreHeaderCode = `
package null

import (
	"log"
    "runtime/debug"
	"encoding/json"
)

//Auto-generated code; DONT EDIT THIS CODE


`

const allTypesHeaderCode = `

package null

//Auto-generated code; DONT EDIT THIS CODE

var (
	AllTypes = map[string]bool {
	}
)

`

const basicBodyCode = `

type {{.TypeName}} struct {
	val {{.BuiltInTypeName}}
	valid bool
}

func New{{.TypeName}}(val {{.BuiltInTypeName}}) {{.TypeName}} {
	return {{.TypeName}}{val: val, valid:true }
}

func (t *{{.TypeName}}) Set(val {{.BuiltInTypeName}}) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *{{.TypeName}}) Get() {{.BuiltInTypeName}} {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:{{.TypeName}}!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *{{.TypeName}}) GetPtr() *{{.BuiltInTypeName}} {
	return &t.val
}

func (t *{{.TypeName}}) IsNull() bool {
	return !t.valid
}


`

const basicSetSafeCode = `

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *{{.TypeName}}) SetSafe(val {{.BuiltInTypeName}}) {
	t.val = val
	t.valid = true
}

`
const setSafeIsValueCode = `

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *{{.TypeName}}) SetSafe(val {{.BuiltInTypeName}}) {
	if !IsValue{{.TypeName}}(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:{{.TypeName}}!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

`

const isValueSwitchCode = `

func IsValue{{.TypeName}}(val {{.BuiltInTypeName}}) bool {
	switch val {
	{{.SwitchCode}}
	default:
			return false

	}
}

`
const isValueMapCode = `

var (
	map{{.TypeName}} = map[{{.BuiltInTypeName}}]bool {
	{{.MapValues}}
	}
)

func IsValue{{.TypeName}}(val {{.BuiltInTypeName}}) bool {
	_, ok := map{{.TypeName}}[val]
	if !ok {
		return false
	}
	return true
}

`

const basicMarshalCode = `

func (t *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}

`

const numericalMapMarshalCode = `

var (
	map{{.TypeName}}NumToText = map[{{.BuiltInTypeName}}] string {
	{{.MapValues}}
	}
)

func (t *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	v, ok := map{{.TypeName}}NumToText[t.val]
	if ok {
		return json.Marshal(v)
	}
	return json.Marshal(t.val)
}

`

var templateBasicHeaderCode = template.Must(template.New("tempalteBasicHeaderCode").Parse(basicHeaderCode))
var templateBasicBodyCode = template.Must(template.New("templateBasicBodyCode").Parse(basicBodyCode))
var templateBasicSetSafeCode = template.Must(template.New("templateBasicSetSafeCode").Parse(basicSetSafeCode))
var templateSetSafeIsValueCode = template.Must(template.New("templateSetSafeIsValueCode").Parse(setSafeIsValueCode))
var templateIsValueSwitchCode = template.Must(template.New("templateIsValueSwitchCode").Parse(isValueSwitchCode))
var templateIsValueMapCode = template.Must(template.New("templateIsValueMapCode").Parse(isValueMapCode))
var templateBasicMarshalCode = template.Must(template.New("templateBasicMarshalCode").Parse(basicMarshalCode))
var templateNumericalMapMarshalCode = template.Must(template.New("templateBasicIsSafeCode").Parse(numericalMapMarshalCode))

func generateStringEnumSwitch_IsValue(typ *Type) string {
	result := basicSetSafeCode + `

func IsValue{{.TypeName}}(val {{.BuiltInTypeName}}) bool {

	switch val {
`
	for _, v := range typ.values {
		result += "\tcase " + typ.pkgName + "." + v.name + ":\n"
		result += "\t\treturn true\n"
	}
	result += `
		default:
			return false
	}	
}
	`
	return result
}

func generateStringEnumMap_IsValue(typ *Type) string {
	result := basicSetSafeCode + `

var (
	map{{.TypeName}} = map[{{.BuiltInTypeName}}]bool {
	`
	for _, v := range typ.values {
		result += "\t" + typ.pkgName + "." + v.name + ": true,\n"
	}
	result += `
	}
)

func IsValue{{.TypeName}}(val {{.BuiltInTypeName}}) bool {
	_, ok := map{{.TypeName}}[val]
	if !ok {
		return false
	}
	return true
}
`
	return result
}

func generateIntMarshal(typ *Type) string {
	result := `
var (
	map{{.TypeName}}NumToText = map[{{.BuiltInTypeName}}] string {
`
	for _, v := range typ.values {
		result += "\t" + typ.pkgName + "." + v.name + ": \"" + v.comment + "\",\n"
	}
	result += `
	}
)
func (t *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	v, ok := map{{.TypeName}}NumToText[t.val]
	if ok {
		return json.Marshal(v)
	}
	return json.Marshal(t.val)
}

`
	return result
}
