package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeSample1 struct {
	val   enum.TypeSample1
	valid bool
}

func NewTypeSample1(val enum.TypeSample1) TypeSample1 {
	return TypeSample1{val: val, valid: true}
}

func (t *TypeSample1) Set(val enum.TypeSample1) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeSample1) Get() enum.TypeSample1 {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypeSample1!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypeSample1) GetPtr() *enum.TypeSample1 {
	return &t.val
}

func (t *TypeSample1) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeSample1) SetSafe(val enum.TypeSample1) {
	if !IsValueTypeSample1(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypeSample1!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

var (
	mapTypeSample1IDToText = map[enum.TypeSample1]string{
		enum.Sample1_1: "Hero",
	}
)

func _LookupTypeSample1IDToText(val enum.TypeSample1) (string, bool) {
	switch val {
	case enum.Sample1_1:
		return "Hero", true
	default:
		return "", false
	}
}

func IsValueTypeSample1(val enum.TypeSample1) bool {
	_, ok := _LookupTypeSample1IDToText(val)
	return ok
}

func (t *TypeSample1) MarshalJSON() ([]byte, error) {
	val, ok := _LookupTypeSample1IDToText(t.val)
	if ok {
		return json.Marshal(val)
	}
	return json.Marshal(t.val)
}
