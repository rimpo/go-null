package null

import (
	"encoding/json"
	"log"
	"runtime/debug"
)

//Auto-generated code; DONT EDIT THIS CODE

type Int struct {
	val   int
	valid bool
}

func NewInt(val int) Int {
	return Int{val: val, valid: true}
}

func (t *Int) Set(val int) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *Int) Get() int {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:Int!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *Int) GetPtr() *int {
	return &t.val
}

func (t *Int) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *Int) SetSafe(val int) {
	t.val = val
	t.valid = true
}

func (t *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
