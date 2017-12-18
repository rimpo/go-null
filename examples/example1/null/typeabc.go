package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeABC struct {
	val   enum.TypeABC
	valid bool
}

func NewTypeABC(val enum.TypeABC) TypeABC {
	return TypeABC{val: val, valid: true}
}

func (t *TypeABC) Set(val enum.TypeABC) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeABC) Get() enum.TypeABC {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypeABC!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypeABC) GetPtr() *enum.TypeABC {
	return &t.val
}

func (t *TypeABC) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeABC) SetSafe(val enum.TypeABC) {
	if !IsValueTypeABC(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypeABC!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

func _LookupTypeABCIDToText(val enum.TypeABC) (string, bool) {
	res, ok := enum.MapTypeABCIDToText[int(val)]
	return res, ok

}

func IsValueTypeABC(val enum.TypeABC) bool {
	_, ok := _LookupTypeABCIDToText(val)
	return ok
}

func (t *TypeABC) MarshalJSON() ([]byte, error) {
	val, ok := _LookupTypeABCIDToText(t.val)
	if ok {
		return json.Marshal(val)
	}
	return json.Marshal(t.val)
}
