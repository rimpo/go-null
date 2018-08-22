package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeMapString struct {
	val   enum.TypeMapString
	valid bool
}

func NewTypeMapString(val enum.TypeMapString) TypeMapString {
	return TypeMapString{val: val, valid: true}
}

func (t *TypeMapString) Set(val enum.TypeMapString) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeMapString) Get() enum.TypeMapString {
	if t.IsNull() {
		log.WithFields(log.Fields{"type": "TypeMapString", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *TypeMapString) GetUnsafe() enum.TypeMapString {
	return t.val
}

func (t *TypeMapString) GetPtr() *enum.TypeMapString {
	return &t.val
}

func (t *TypeMapString) IsNull() bool {
	return !t.valid
}

func (t *TypeMapString) Reset() {
	t.valid = false
}

func (t *TypeMapString) IsEmpty() bool {
	return t.IsNull() || len(string(t.val)) == 0
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *TypeMapString) SetSafe(val enum.TypeMapString) {
	t.val = val
	t.valid = true
}

func (t *TypeMapString) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
