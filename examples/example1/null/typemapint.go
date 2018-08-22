package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeMapInt struct {
	val   enum.TypeMapInt
	valid bool
}

func NewTypeMapInt(val enum.TypeMapInt) TypeMapInt {
	return TypeMapInt{val: val, valid: true}
}

func (t *TypeMapInt) Set(val enum.TypeMapInt) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeMapInt) Get() enum.TypeMapInt {
	if t.IsNull() {
		log.WithFields(log.Fields{"type": "TypeMapInt", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *TypeMapInt) GetUnsafe() enum.TypeMapInt {
	return t.val
}

func (t *TypeMapInt) GetPtr() *enum.TypeMapInt {
	return &t.val
}

func (t *TypeMapInt) IsNull() bool {
	return !t.valid
}

func (t *TypeMapInt) Reset() {
	t.valid = false
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *TypeMapInt) SetSafe(val enum.TypeMapInt) {
	t.val = val
	t.valid = true
}

func (t *TypeMapInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
