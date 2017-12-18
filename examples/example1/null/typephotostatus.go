package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypePhotoStatus struct {
	val   enum.TypePhotoStatus
	valid bool
}

func NewTypePhotoStatus(val enum.TypePhotoStatus) TypePhotoStatus {
	return TypePhotoStatus{val: val, valid: true}
}

func (t *TypePhotoStatus) Set(val enum.TypePhotoStatus) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypePhotoStatus) Get() enum.TypePhotoStatus {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypePhotoStatus!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypePhotoStatus) GetPtr() *enum.TypePhotoStatus {
	return &t.val
}

func (t *TypePhotoStatus) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypePhotoStatus) SetSafe(val enum.TypePhotoStatus) {
	if !IsValueTypePhotoStatus(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypePhotoStatus!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

var (
	mapTypePhotoStatusIDToText = map[enum.TypePhotoStatus]string{
		enum.PhotoNotAvailable: "not available",
		enum.PhotoComingSoon:   "coming soon",
		enum.PhotoAvailable:    "available",
	}
)

func _LookupTypePhotoStatusIDToText(val enum.TypePhotoStatus) (string, bool) {
	switch val {
	case enum.PhotoNotAvailable:
		return "not available", true
	case enum.PhotoComingSoon:
		return "coming soon", true
	case enum.PhotoAvailable:
		return "available", true
	default:
		return "", false
	}
}

func IsValueTypePhotoStatus(val enum.TypePhotoStatus) bool {
	_, ok := _LookupTypePhotoStatusIDToText(val)
	return ok
}

func (t *TypePhotoStatus) MarshalJSON() ([]byte, error) {
	val, ok := _LookupTypePhotoStatusIDToText(t.val)
	if ok {
		return json.Marshal(val)
	}
	return json.Marshal(t.val)
}
