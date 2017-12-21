package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeMemberStatus struct {
	val   enum.TypeMemberStatus
	valid bool
}

func NewTypeMemberStatus(val enum.TypeMemberStatus) TypeMemberStatus {
	return TypeMemberStatus{val: val, valid: true}
}

func (t *TypeMemberStatus) Set(val enum.TypeMemberStatus) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeMemberStatus) Get() enum.TypeMemberStatus {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypeMemberStatus!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypeMemberStatus) GetPtr() *enum.TypeMemberStatus {
	return &t.val
}

func (t *TypeMemberStatus) IsNull() bool {
	return !t.valid
}

func (t *TypeMemberStatus) IsEmpty() bool {
	return t.IsNull() || len(string(t.val)) == 0
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeMemberStatus) SetSafe(val enum.TypeMemberStatus) {
	if !IsValueTypeMemberStatus(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypeMemberStatus!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

var (
	mapTypeMemberStatusIDToText = map[enum.TypeMemberStatus]string{
		enum.MemberActivated:    "Activated",
		enum.MemberDeactivated:  "Deactivated",
		enum.MemberToBeScreened: "To Be Screened",
		enum.MemberSuspended:    "Suspended",
	}
)

func _LookupTypeMemberStatusIDToText(val enum.TypeMemberStatus) (string, bool) {
	switch val {
	case enum.MemberActivated:
		return "Activated", true
	case enum.MemberDeactivated:
		return "Deactivated", true
	case enum.MemberToBeScreened:
		return "To Be Screened", true
	case enum.MemberSuspended:
		return "Suspended", true
	default:
		return "", false
	}
}

func IsValueTypeMemberStatus(val enum.TypeMemberStatus) bool {
	_, ok := _LookupTypeMemberStatusIDToText(val)
	return ok
}

func (t *TypeMemberStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
