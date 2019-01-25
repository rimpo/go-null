package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
		log.WithFields(log.Fields{"type": "TypeNamePrivacy", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *TypeNamePrivacy) GetUnsafe() enum.TypeNamePrivacy {
	return t.val
}

func (t *TypeNamePrivacy) GetPtr() *enum.TypeNamePrivacy {
	return &t.val
}

func (t *TypeNamePrivacy) IsNull() bool {
	return !t.valid
}

func (t *TypeNamePrivacy) Reset() {
	t.valid = false
}

func (t *TypeNamePrivacy) IsEmpty() bool {
	return t.IsNull() || len(string(t.val)) == 0
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeNamePrivacy) SetSafe(val enum.TypeNamePrivacy) {
	if !IsValueTypeNamePrivacy(val) {
		log.WithFields(log.Fields{"type": "TypeNamePrivacy", "value": val, "stack": string(debug.Stack()[:])}).Warn("unknown value assigned !!!.")
	}
	t.val = val
	t.valid = true
}

var (
	mapTypeNamePrivacyIDToText = map[enum.TypeNamePrivacy]string{
		enum.HideFirstName:    "partial_name",
		enum.HideLastName:     "partial_name_inverse",
		enum.DisplayFullName:  "full_name",
		enum.DisplayProfileID: "profile_id",
	}
)

func _LookupTypeNamePrivacyIDToText(val enum.TypeNamePrivacy) (string, bool) {
	switch val {
	case enum.HideFirstName:
		return "partial_name", true
	case enum.HideLastName:
		return "partial_name_inverse", true
	case enum.DisplayFullName:
		return "full_name", true
	case enum.DisplayProfileID:
		return "profile_id", true
	default:
		return "", false
	}
}

func IsValueTypeNamePrivacy(val enum.TypeNamePrivacy) bool {
	_, ok := _LookupTypeNamePrivacyIDToText(val)
	return ok
}

func (t *TypeNamePrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
