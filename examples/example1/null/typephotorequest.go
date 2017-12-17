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

func IsValueTypePhotoRequest(val enum.TypePhotoRequest) bool {
	switch val {
	case enum.PhotoRequestNotAvailable:
		return true
	case enum.PhotoRequestSent:
		return true
	case enum.PhotoRequestAccepted:
		return true
	case enum.PhotoRequestRejected:
		return true
	case enum.PhotoRequestDelete:
		return true

	default:
		return false

	}
}

var (
	mapTypePhotoRequestNumToText = map[enum.TypePhotoRequest]string{
		enum.PhotoRequestNotAvailable: "photo request not available",
		enum.PhotoRequestSent:         "photo request sent",
		enum.PhotoRequestAccepted:     "photo request rejected",
		enum.PhotoRequestRejected:     "photo rejected",
		enum.PhotoRequestDelete:       "deleted",
	}
)

func (t *TypePhotoRequest) MarshalJSON() ([]byte, error) {
	v, ok := mapTypePhotoRequestNumToText[t.val]
	if ok {
		return json.Marshal(v)
	}
	return json.Marshal(t.val)
}
