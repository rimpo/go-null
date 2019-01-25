package null

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
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
		log.WithFields(log.Fields{"type": "TypePhonePrivacy", "stack": string(debug.Stack()[:])}).Warn("null value used !!!.")
	}
	return t.val
}

func (t *TypePhonePrivacy) GetUnsafe() enum.TypePhonePrivacy {
	return t.val
}

func (t *TypePhonePrivacy) GetPtr() *enum.TypePhonePrivacy {
	return &t.val
}

func (t *TypePhonePrivacy) IsNull() bool {
	return !t.valid
}

func (t *TypePhonePrivacy) Reset() {
	t.valid = false
}

func (t *TypePhonePrivacy) IsEmpty() bool {
	return t.IsNull() || len(string(t.val)) == 0
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypePhonePrivacy) SetSafe(val enum.TypePhonePrivacy) {
	if !IsValueTypePhonePrivacy(val) {
		log.WithFields(log.Fields{"type": "TypePhonePrivacy", "value": val, "stack": string(debug.Stack()[:])}).Warn("unknown value assigned !!!.")
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
