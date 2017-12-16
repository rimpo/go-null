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

func IsValueTypeShowPhoto(val enum.TypeShowPhoto) bool {
	switch val {
	case enum.ShowPhoto:
		return true
	case enum.ShowPhotoNotAvailable:
		return true
	case enum.ShowRequestPhoto:
		return true
	case enum.ShowRequestPhotoSent:
		return true
	case enum.ShowRequestPhotoPassword:
		return true
	case enum.ShowRequestPhotoPasswordSent:
		return true
	case enum.ShowAddPhoto:
		return true
	case enum.ShowPhotoComingSoon:
		return true
	case enum.ShowMemberPhotoNotScreened:
		return true

	default:
		return false

	}
}

func (t *TypeShowPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.val)
}
