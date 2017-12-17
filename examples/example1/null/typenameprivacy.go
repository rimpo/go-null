package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeNamePrivacy struct {
	val   enum.TypeNamePrivacy
	valid bool
}

func NewTypeNamePrivacy(val enum.TypeNamePrivacy) TypeNamePrivacy {
	return TypeNamePrivacy{val: val, valid: true}
}

func (t *TypeNamePrivacy) Set(val enum.TypeNamePrivacy) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeNamePrivacy) Get() enum.TypeNamePrivacy {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypeNamePrivacy!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypeNamePrivacy) GetPtr() *enum.TypeNamePrivacy {
	return &t.val
}

func (t *TypeNamePrivacy) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeNamePrivacy) SetSafe(val enum.TypeNamePrivacy) {
	if !IsValueTypeNamePrivacy(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypeNamePrivacy!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

func IsValueTypeNamePrivacy(val enum.TypeNamePrivacy) bool {
	switch val {
	case enum.HideFirstName:
		return true
	case enum.HideLastName:
		return true
	case enum.DisplayFullName:
		return true
	case enum.DisplayProfileID:
		return true

	default:
		return false

	}
}

func (t *TypeNamePrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
