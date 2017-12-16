package null

import (
	"encoding/json"
	"log"
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
		log.Printf("ERROR: Fetching a null value from type:TypeSample3!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypeSample3) GetPtr() *enum.TypeSample3 {
	return &t.val
}

func (t *TypeSample3) IsNull() bool {
	return !t.valid
}

func IsValueTypeSample3(val enum.TypeSample3) bool {
	switch val {
	case enum.Sample3_1:
		return true
	case enum.Sample3_2:
		return true
	case enum.Sample3_3:
		return true

	default:
		return false

	}
}

var (
	mapTypeSample3NumToText = map[enum.TypeSample3]string{
		enum.Sample3_1: "One",
		enum.Sample3_2: "Two",
		enum.Sample3_3: "Three",
	}
)

func (t *TypeSample3) MarshalJSON() ([]byte, error) {
	v, ok := mapTypeSample3NumToText[t.val]
	if ok {
		return json.Marshal(v)
	}
	return json.Marshal(t.val)
}
