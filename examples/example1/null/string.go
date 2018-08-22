package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
		log.WithFields(log.Fields{"type": "String", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *String) GetUnsafe() string {
	return t.val
}

func (t *String) GetPtr() *string {
	return &t.val
}

func (t *String) IsNull() bool {
	return !t.valid
}

func (t *String) Reset() {
	t.valid = false
}

func (t *String) IsEmpty() bool {
	return t.IsNull() || len(string(t.val)) == 0
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *String) SetSafe(val string) {
	t.val = val
	t.valid = true
}

func (t *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
