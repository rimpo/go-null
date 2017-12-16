package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/typ"
)

//Auto-generated code; DONT EDIT THIS CODE

type Float struct {
	val   typ.Float
	valid bool
}

func NewFloat(val typ.Float) Float {
	return Float{val: val, valid: true}
}

func (t *Float) Set(val typ.Float) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *Float) Get() typ.Float {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:Float!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *Float) GetPtr() *typ.Float {
	return &t.val
}

func (t *Float) IsNull() bool {
	return !t.valid
}

func (t *Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
