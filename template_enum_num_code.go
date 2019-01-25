package main

const templateEnumNumCodeString = `
package null

import (
	"encoding/json"
	"runtime/debug"
	log "github.com/Sirupsen/logrus"

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
		log.WithFields(log.Fields{"type":"{{.TypeName}}", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *{{.TypeName}}) GetUnsafe() {{.SourceType}} {
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
		log.WithFields(log.Fields{"type":"{{.TypeName}}", "value": val, "stack": string(debug.Stack()[:])}).Warn("unknown value assigned !!!.")
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

func (t *{{.TypeName}}) GetDisplay() string {
	if !t.valid {
		return ""
	}
	val, ok := _Lookup{{.TypeName}}IDToText(t.val)
	if ok {
		return val
	}
	return ""
}

func (t *{{.TypeName}}) MarshalJSON() ([]byte, error) {
	val, ok := _Lookup{{.TypeName}}IDToText(t.val)
	if ok {
		return json.Marshal(val)
	}
	return json.Marshal(t.val)
}
`
