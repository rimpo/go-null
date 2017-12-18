package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypePhonePrivacy struct {
	val   enum.TypePhonePrivacy
	valid bool
}

func NewTypePhonePrivacy(val enum.TypePhonePrivacy) TypePhonePrivacy {
	return TypePhonePrivacy{val: val, valid: true}
}

func (t *TypePhonePrivacy) Set(val enum.TypePhonePrivacy) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypePhonePrivacy) Get() enum.TypePhonePrivacy {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypePhonePrivacy!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypePhonePrivacy) GetPtr() *enum.TypePhonePrivacy {
	return &t.val
}

func (t *TypePhonePrivacy) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypePhonePrivacy) SetSafe(val enum.TypePhonePrivacy) {
	if !IsValueTypePhonePrivacy(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypePhonePrivacy!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

var (
	mapTypePhonePrivacyIDToText = map[enum.TypePhonePrivacy]string{
		enum.PhoneVisibleToPremium:              "Show All",
		enum.PhoneVisibleToPreimumWishToConnect: "When I Contact",
		enum.PhoneNumberHide:                    "Hide My Number",
	}
)

func _LookupTypePhonePrivacyIDToText(val enum.TypePhonePrivacy) (string, bool) {
	switch val {
	case enum.PhoneVisibleToPremium:
		return "Show All", true
	case enum.PhoneVisibleToPreimumWishToConnect:
		return "When I Contact", true
	case enum.PhoneNumberHide:
		return "Hide My Number", true
	default:
		return "", false
	}
}

func IsValueTypePhonePrivacy(val enum.TypePhonePrivacy) bool {
	_, ok := _LookupTypePhonePrivacyIDToText(val)
	return ok
}

func (t *TypePhonePrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
