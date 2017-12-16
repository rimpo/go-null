package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeSample2 struct {
	val   enum.TypeSample2
	valid bool
}

func NewTypeSample2(val enum.TypeSample2) TypeSample2 {
	return TypeSample2{val: val, valid: true}
}

func (t *TypeSample2) Set(val enum.TypeSample2) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeSample2) Get() enum.TypeSample2 {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypeSample2!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypeSample2) GetPtr() *enum.TypeSample2 {
	return &t.val
}

func (t *TypeSample2) IsNull() bool {
	return !t.valid
}

var (
	mapTypeSample2 = map[enum.TypeSample2]bool{
		enum.Sample2_1:  true,
		enum.Sample2_2:  true,
		enum.Sample2_3:  true,
		enum.Sample2_4:  true,
		enum.Sample2_5:  true,
		enum.Sample2_6:  true,
		enum.Sample2_7:  true,
		enum.Sample2_8:  true,
		enum.Sample2_9:  true,
		enum.Sample2_10: true,
		enum.Sample2_11: true,
	}
)

func IsValueTypeSample2(val enum.TypeSample2) bool {
	_, ok := mapTypeSample2[val]
	if !ok {
		return false
	}
	return true
}

var (
	mapTypeSample2NumToText = map[enum.TypeSample2]string{
		enum.Sample2_1:  "A",
		enum.Sample2_2:  "B",
		enum.Sample2_3:  "C",
		enum.Sample2_4:  "D",
		enum.Sample2_5:  "E",
		enum.Sample2_6:  "F",
		enum.Sample2_7:  "G",
		enum.Sample2_8:  "H",
		enum.Sample2_9:  "I",
		enum.Sample2_10: "J",
		enum.Sample2_11: "K",
	}
)

func (t *TypeSample2) MarshalJSON() ([]byte, error) {
	v, ok := mapTypeSample2NumToText[t.val]
	if ok {
		return json.Marshal(v)
	}
	return json.Marshal(t.val)
}
