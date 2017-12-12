package enums

//go:generate make-null-type -type=TypeMemberStatus
type TypeMemberStatus string

const (
	MemberActivated    TypeMemberStatus = "Activated"      //hello world
	MemberDeactivated                   = "Deactivated"    //hiii
	MemberToBeScreened                  = "To Be Screened" //to alksdj
	MemberSuspended                     = "Suspended"
)

const (
	X = 1
	Y = "asdkfj"
)

type TypePhotoStatus int

const (
	PhotoNotAvailable TypePhotoStatus = iota
	PhotoComingSoon
	PhotoAvailable
)

type TypeShowPhoto string

const (
	ShowPhoto                    TypeShowPhoto = "show_photo"
	ShowPhotoNotAvailable                      = "show_photo_not_available"
	ShowRequestPhoto                           = "show_request_photo"
	ShowRequestPhotoSent                       = "show_request_photo_sent"
	ShowRequestPhotoPassword                   = "show_request_photo_password"
	ShowRequestPhotoPasswordSent               = "show_request_photo_password_sent"
	ShowAddPhoto                               = "show_add_photo"
	ShowPhotoComingSoon                        = "show_comming_soon"
	ShowMemberPhotoNotScreened                 = "show_member_photo_not_screened"
)

type TypePhotoRequest int

const (
	PhotoRequestNotAvailable TypePhotoRequest = iota
	PhotoRequestSent
	PhotoRequestAccepted
	PhotoRequestRejected
	PhotoRequestDelete
)

//go:generate make-null-type -type=TypeSample1
type TypeSample1 int

const (
	Sample1_1 TypeSample1 = iota
)

//go:generate make-null-type -type=TypeSample2
type TypeSample2 int

const (
	Sample2_1 TypeSample2 = iota
	Sample2_2
	Sample2_3
	Sample2_4
	Sample2_5
	Sample2_6
	Sample2_7
	Sample2_8
	Sample2_9
	Sample2_10
	Sample2_11
)

//go:generate make-null-type -type=TypeSample3
type TypeSample3 int

const (
	Sample3_1 TypeSample3 = -1
	Sample3_2             = -2
	Sample3_3             = -3
)

//go:generate make-null-type -type=TypePhotoPasswordRequest
type TypePhotoPasswordRequest int

const (
	PhotoPasswordRequestNotAvailable TypePhotoPasswordRequest = iota
	PhotoPasswordRequestSent
	PhotoPasswordRequestAccepted
	PhotoPasswordRequestRejected
	PhotoPasswordRequestDelete
)

type TypeNamePrivacy string

const (
	HideFirstName    TypeNamePrivacy = "partial_name"
	HideLastName                     = "partial_name_inverse"
	DisplayFullName                  = "full_name"
	DisplayProfileID                 = "profile_id"
)

type TypePhonePrivacy string

const (
	PhoneVisibleToPremium              TypePhonePrivacy = "Show All"
	PhoneVisibleToPreimumWishToConnect                  = "When I Contact"
	PhoneNumberHide                                     = "Hide My Number"
)
