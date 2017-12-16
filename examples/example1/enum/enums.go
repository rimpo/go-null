package enum

//go:generate go-null -package=github.com/rimpo/go-null/examples/example1 -output=..

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
	PhotoNotAvailable TypePhotoStatus = iota //not available
	PhotoComingSoon                          //coming soon
	PhotoAvailable                           //available
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
	PhotoRequestNotAvailable TypePhotoRequest = 10 //photo request not available
	PhotoRequestSent                          = 11 //photo request sent
	PhotoRequestAccepted                      = 12 //photo request rejected
	PhotoRequestRejected                      = 13 //photo rejected
	PhotoRequestDelete                        = 14 //deleted
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
