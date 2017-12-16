package null

import (
	"encoding/json"
	"log"
	"runtime/debug"
)

//Auto-generated code; DONT EDIT THIS CODE

type String struct {
	val   string
	valid bool
}

func NewString(val string) String {
	return String{val: val, valid: true}
}

func (t *String) Set(val string) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *String) Get() string {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:String!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *String) GetPtr() *string {
	return &t.val
}

func (t *String) IsNull() bool {
	return !t.valid
}

func (t *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
