package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/typ"
)

//Auto-generated code; DONT EDIT THIS CODE

type DateTimeInt struct {
	val   typ.DateTimeInt
	valid bool
}

func NewDateTimeInt(val typ.DateTimeInt) DateTimeInt {
	return DateTimeInt{val: val, valid: true}
}

func (t *DateTimeInt) Set(val typ.DateTimeInt) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *DateTimeInt) Get() typ.DateTimeInt {
	if t.IsNull() {
		log.WithFields(log.Fields{"type": "DateTimeInt", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *DateTimeInt) GetUnsafe() typ.DateTimeInt {
	return t.val
}

func (t *DateTimeInt) GetPtr() *typ.DateTimeInt {
	return &t.val
}

func (t *DateTimeInt) IsNull() bool {
	return !t.valid
}

func (t *DateTimeInt) Reset() {
	t.valid = false
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *DateTimeInt) SetSafe(val typ.DateTimeInt) {
	t.val = val
	t.valid = true
}

func (t *DateTimeInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
