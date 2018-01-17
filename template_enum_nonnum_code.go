package main

const templateEnumNonNumCodeString = `
package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	{{.ImportLib}}	
)

//Auto-generated code; DONT EDIT THIS CODE

type {{.TypeName}} struct {
	val   {{.SourceType}}
	valid bool
}

func New{{.TypeName}}(val {{.SourceType}}) {{.TypeName}} {
	return {{.TypeName}}{val: val, valid: true}
}

func (t *{{.TypeName}}) Set(val {{.SourceType}}) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *{{.TypeName}}) Get() {{.SourceType}} {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:{{.TypeName}}!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *{{.TypeName}}) GetPtr() *{{.SourceType}} {
	return &t.val
}

func (t *{{.TypeName}}) IsNull() bool {
	return !t.valid
}

func (t *{{.TypeName}}) Reset() {
	t.valid = false
}


{{.IsEmptyCode}}


//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *{{.TypeName}}) SetSafe(val {{.SourceType}}) {
	if !IsValue{{.TypeName}}(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:{{.TypeName}}!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

var (
	map{{.TypeName}}IDToText = map[{{.SourceType}}] string {
	{{.MapValues}}
	}
)


func _Lookup{{.TypeName}}IDToText(val {{.SourceType}}) (string, bool) {
	{{.LookupCode}}
}


func IsValue{{.TypeName}}(val {{.SourceType}}) bool {
	_, ok := _Lookup{{.TypeName}}IDToText(val)
	return ok
}

func (t *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
`
