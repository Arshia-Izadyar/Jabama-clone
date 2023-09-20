package service_errors

const (
	OtpExists     = "otp exits"
	OtpUsed       = "otp used"
	OtpInvalid    = "otp invalid"
	ClaimNotFound = "claim not found"

	// user
	EmailExists    = "email already exits"
	UsernameExists = "Username already exits"
	WrongPassword  = "WrongPassword"

	TokenNotPresent = "no token provided"
	TokenExpired    = "token is expired !"
	TokenInvalid    = "provided token is invalid"
	NotRefreshToken = "provided token is not a refresh token"
	InternalError   = "some thing happened"

	PermissionDenied = "Permission Denied"
)
