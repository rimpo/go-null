package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeSample3 struct {
	val   enum.TypeSample3
	valid bool
}

func NewTypeSample3(val enum.TypeSample3) TypeSample3 {
	return TypeSample3{val: val, valid: true}
}

func (t *TypeSample3) Set(val enum.TypeSample3) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeSample3) Get() enum.TypeSample3 {
	if t.IsNull() {
		log.WithFields(log.Fields{"type": "TypeSample3", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *TypeSample3) GetUnsafe() enum.TypeSample3 {
	return t.val
}

func (t *TypeSample3) GetPtr() *enum.TypeSample3 {
	return &t.val
}

func (t *TypeSample3) IsNull() bool {
	return !t.valid
}

func (t *TypeSample3) Reset() {
	t.valid = false
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeSample3) SetSafe(val enum.TypeSample3) {
	if !IsValueTypeSample3(val) {
		log.WithFields(log.Fields{"type": "TypeSample3", "value": val, "stack": string(debug.Stack()[:])}).Warn("unknown value assigned !!!.")
	}
	t.val = val
	t.valid = true
}

var (
	mapTypeSample3IDToText = map[enum.TypeSample3]string{
		enum.Sample3_1: "One",
		enum.Sample3_2: "Two",
		enum.Sample3_3: "Three",
	}
)

func _LookupTypeSample3IDToText(val enum.TypeSample3) (string, bool) {
	switch val {
	case enum.Sample3_1:
		return "One", true
	case enum.Sample3_2:
		return "Two", true
	case enum.Sample3_3:
		return "Three", true
	default:
		return "", false
	}
}

func IsValueTypeSample3(val enum.TypeSample3) bool {
	_, ok := _LookupTypeSample3IDToText(val)
	return ok
}

func (t *TypeSample3) GetDisplay() string {
	if !t.valid {
		return ""
	}
	val, ok := _LookupTypeSample3IDToText(t.val)
	if ok {
		return val
	}
	return ""
}

func (t *TypeSample3) MarshalJSON() ([]byte, error) {
	val, ok := _LookupTypeSample3IDToText(t.val)
	if ok {
		return json.Marshal(val)
	}
	return json.Marshal(t.val)
}
