package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
		log.WithFields(log.Fields{"type": "Int", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *Int) GetUnsafe() int {
	return t.val
}

func (t *Int) GetPtr() *int {
	return &t.val
}

func (t *Int) IsNull() bool {
	return !t.valid
}

func (t *Int) Reset() {
	t.valid = false
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). //dummy function (same as Set)
func (t *Int) SetSafe(val int) {
	t.val = val
	t.valid = true
}

func (t *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
