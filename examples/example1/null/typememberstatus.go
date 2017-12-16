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

func IsValueTypeMemberStatus(val enum.TypeMemberStatus) bool {
	switch val {
	case enum.MemberActivated:
		return true
	case enum.MemberDeactivated:
		return true
	case enum.MemberToBeScreened:
		return true
	case enum.MemberSuspended:
		return true

	default:
		return false

	}
}

func (t *TypeMemberStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
