package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypePhotoRequest struct {
	val   enum.TypePhotoRequest
	valid bool
}

func NewTypePhotoRequest(val enum.TypePhotoRequest) TypePhotoRequest {
	return TypePhotoRequest{val: val, valid: true}
}

func (t *TypePhotoRequest) Set(val enum.TypePhotoRequest) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypePhotoRequest) Get() enum.TypePhotoRequest {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypePhotoRequest!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypePhotoRequest) GetPtr() *enum.TypePhotoRequest {
	return &t.val
}

func (t *TypePhotoRequest) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypePhotoRequest) SetSafe(val enum.TypePhotoRequest) {
	if !IsValueTypePhotoRequest(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypePhotoRequest!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

var (
	mapTypePhotoRequestIDToText = map[enum.TypePhotoRequest]string{
		enum.PhotoRequestNotAvailable: "photo request not available",
		enum.PhotoRequestSent:         "photo request sent",
		enum.PhotoRequestAccepted:     "photo request rejected",
		enum.PhotoRequestRejected:     "photo rejected",
		enum.PhotoRequestDelete:       "deleted",
	}
)

func _LookupTypePhotoRequestIDToText(val enum.TypePhotoRequest) (string, bool) {
	switch val {
	case enum.PhotoRequestNotAvailable:
		return "photo request not available", true
	case enum.PhotoRequestSent:
		return "photo request sent", true
	case enum.PhotoRequestAccepted:
		return "photo request rejected", true
	case enum.PhotoRequestRejected:
		return "photo rejected", true
	case enum.PhotoRequestDelete:
		return "deleted", true
	default:
		return "", false
	}
}

func IsValueTypePhotoRequest(val enum.TypePhotoRequest) bool {
	_, ok := _LookupTypePhotoRequestIDToText(val)
	return ok
}

func (t *TypePhotoRequest) MarshalJSON() ([]byte, error) {
	val, ok := _LookupTypePhotoRequestIDToText(t.val)
	if ok {
		return json.Marshal(val)
	}
	return json.Marshal(t.val)
}
