package null

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/rimpo/go-null/examples/example1/enum"
)

//Auto-generated code; DONT EDIT THIS CODE

type TypeShowPhoto struct {
	val   enum.TypeShowPhoto
	valid bool
}

func NewTypeShowPhoto(val enum.TypeShowPhoto) TypeShowPhoto {
	return TypeShowPhoto{val: val, valid: true}
}

func (t *TypeShowPhoto) Set(val enum.TypeShowPhoto) {
	t.val = val
	t.valid = true
}

//Logs error message
func (t *TypeShowPhoto) Get() enum.TypeShowPhoto {
	if t.IsNull() {
		log.Printf("ERROR: Fetching a null value from type:TypeShowPhoto!!.\n")
		debug.PrintStack()
	}
	return t.val
}

func (t *TypeShowPhoto) GetPtr() *enum.TypeShowPhoto {
	return &t.val
}

func (t *TypeShowPhoto) IsNull() bool {
	return !t.valid
}

//Must for loading from external data (i.e. database, elastic, redis, etc.). logs error message
func (t *TypeShowPhoto) SetSafe(val enum.TypeShowPhoto) {
	if !IsValueTypeShowPhoto(val) {
		log.Printf("ERROR: Unknown value:%v assigned to type:TypeShowPhoto!!.\n", val)
		debug.PrintStack()
	}
	t.val = val
	t.valid = true
}

var (
	mapTypeShowPhotoIDToText = map[enum.TypeShowPhoto]string{
		enum.ShowPhoto:                    "show_photo",
		enum.ShowPhotoNotAvailable:        "show_photo_not_available",
		enum.ShowRequestPhoto:             "show_request_photo",
		enum.ShowRequestPhotoSent:         "show_request_photo_sent",
		enum.ShowRequestPhotoPassword:     "show_request_photo_password",
		enum.ShowRequestPhotoPasswordSent: "show_request_photo_password_sent",
		enum.ShowAddPhoto:                 "show_add_photo",
		enum.ShowPhotoComingSoon:          "show_comming_soon",
		enum.ShowMemberPhotoNotScreened:   "show_member_photo_not_screened",
	}
)

func _LookupTypeShowPhotoIDToText(val enum.TypeShowPhoto) (string, bool) {
	switch val {
	case enum.ShowPhoto:
		return "show_photo", true
	case enum.ShowPhotoNotAvailable:
		return "show_photo_not_available", true
	case enum.ShowRequestPhoto:
		return "show_request_photo", true
	case enum.ShowRequestPhotoSent:
		return "show_request_photo_sent", true
	case enum.ShowRequestPhotoPassword:
		return "show_request_photo_password", true
	case enum.ShowRequestPhotoPasswordSent:
		return "show_request_photo_password_sent", true
	case enum.ShowAddPhoto:
		return "show_add_photo", true
	case enum.ShowPhotoComingSoon:
		return "show_comming_soon", true
	case enum.ShowMemberPhotoNotScreened:
		return "show_member_photo_not_screened", true
	default:
		return "", false
	}
}

func IsValueTypeShowPhoto(val enum.TypeShowPhoto) bool {
	_, ok := _LookupTypeShowPhotoIDToText(val)
	return ok
}

func (t *TypeShowPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
