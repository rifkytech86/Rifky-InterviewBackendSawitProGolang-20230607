package commons

import "errors"

const (
	DuplicateKey              = "duplicate key value violates unique constraint"
	ErrorValidatorPassword    = "%s does not meet the password requirements"
	ErrorValidatorPhoneNumber = "%s does not meet the phone number requirements"
	ErrorValidatorFullName    = "%s does not meet the full name requirements"
	ErrorDefaultValidator     = "field: %s, error: %s"
)

var (
	ErrMissingAuthorizationHeader   = errors.New("missing authorization header")
	ErrInvalidTokenFormat           = errors.New("invalid token format")
	ErrInvalidToken                 = errors.New("invalid token")
	ErrInvalidLoadPrivateKey        = errors.New("error load private key")
	ErrInvalidLoadPublicKey         = errors.New("error load public key")
	ErrRegisterValidatorPhoneNumber = errors.New("error register validator phone number")
	ErrRegisterValidatorPassword    = errors.New("error register validator password")
	ErrRegisterValidatorFullName    = errors.New("error register validator full name")
	ErrorConnectionToDatabase       = errors.New("error connecting to database")
	ErrorGeneratePassword           = errors.New("error generate password")
	ErrorBindRequest                = errors.New("error request format")
	ErrorGetUserID                  = errors.New("error get user id")
	ErrorInvalidRequest             = errors.New("invalid request")
	ErrorInternalServer             = errors.New("internal server error")
	ErrorUserAlreadyExist           = errors.New("user already exist")
	ErrorUserNotFound               = errors.New("error user not found")
	ErrorPhoneNumberAlreadyExist    = errors.New("error phone number already exist")
)
