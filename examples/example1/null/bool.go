package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
		log.WithFields(log.Fields{"type": "Bool", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *Bool) GetUnsafe() bool {
	return t.val
}

func (t *Bool) GetPtr() *bool {
	return &t.val
}

func (t *Bool) IsNull() bool {
	return !t.valid
}

func (t *Bool) Reset() {
	t.valid = false
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *Bool) SetSafe(val bool) {
	t.val = val
	t.valid = true
}

func (t *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
