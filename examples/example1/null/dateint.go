package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/typ"
)

//Auto-generated code; DONT EDIT THIS CODE

type DateInt struct {
	val   typ.DateInt
	valid bool
}

func NewDateInt(val typ.DateInt) DateInt {
	return DateInt{val: val, valid: true}
}

func (t *DateInt) Set(val typ.DateInt) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *DateInt) Get() typ.DateInt {
	if t.IsNull() {
		log.WithFields(log.Fields{"type": "DateInt", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *DateInt) GetUnsafe() typ.DateInt {
	return t.val
}

func (t *DateInt) GetPtr() *typ.DateInt {
	return &t.val
}

func (t *DateInt) IsNull() bool {
	return !t.valid
}

func (t *DateInt) Reset() {
	t.valid = false
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *DateInt) SetSafe(val typ.DateInt) {
	t.val = val
	t.valid = true
}

func (t *DateInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
