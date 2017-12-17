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

func IsValueTypeSample1(val enum.TypeSample1) bool {
	switch val {
	case enum.Sample1_1:
		return true

	default:
		return false

	}
}

var (
	mapTypeSample1NumToText = map[enum.TypeSample1]string{
		enum.Sample1_1: "Hero",
	}
)

func (t *TypeSample1) MarshalJSON() ([]byte, error) {
	v, ok := mapTypeSample1NumToText[t.val]
	if ok {
		return json.Marshal(v)
	}
	return json.Marshal(t.val)
}
