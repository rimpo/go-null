package main

const basicHeaderCode = `
package null

import (
	"log"
    "runtime/debug"
)

//Auto-generated code; DONT EDIT THIS CODE


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
	if !t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:{{.TypeName}}!!.\n%v", debug.PrintStack())
	}
	return t.val
}

func (t *{{.TypeName}}) GetPtr() *{{.BuiltInTypeName}} {
	return &t.val
}

func (t *{{.TypeName}}) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *{{.TypeName}}) SetSafe(val {{.BuiltInTypeName}}) {
	if IsValueSafe{{.TypeName}}(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:{{.TypeName}}!!.\n%v", val, debug.PrintStack())
	}
	t.val = val
	t.valid = true
}
`

const basicMarshalCode = `

func (t *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}

`

const isValueSafeCode = `func IsValueSafe{{.TypeName}}(t {{.BuiltInTypeName}}) bool {
	{{.AllConstSwitch}}
}
`

func generateEnumSwitch(values []Value) string {
	result := `

func IsValueSafe{{.TypeName}}(t {{.BuiltInTypeName}}) bool {

	switch t {
`
	for _, v := range values {
		result += "\tcase " + v.str + ":\n"
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

func generateEnumMap(values []Value) {
	result := `
	switch t {
	`

	for _, v := range values {
		result += "\tcase " + v.str + ":\n"
		result += "\t\treturn true\n"
	}
	result += `
		default:
			return false
	}	
	`
}
