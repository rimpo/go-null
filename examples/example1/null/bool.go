package null

import (
	"encoding/json"
	"log"
	"runtime/debug"
)

//Auto-generated code; DONT EDIT THIS CODE

type Bool struct {
	val   bool
	valid bool
}

func NewBool(val bool) Bool {
	return Bool{val: val, valid: true}
}

func (t *Bool) Set(val bool) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *Bool) Get() bool {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:Bool!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *Bool) GetPtr() *bool {
	return &t.val
}

func (t *Bool) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *Bool) SetSafe(val bool) {
	t.val = val
	t.valid = true
}

func (t *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
