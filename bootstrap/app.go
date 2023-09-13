package bootstrap

import (
	"github.com/SawitProRecruitment/UserService/commons"
	"github.com/SawitProRecruitment/UserService/errors"
	"io/ioutil"
)

type Application struct {
	Env            *ENV
	PostgresClient *PostgresClient
	Validator      IValidator
	Harsher        IBcryptHasher
	Jwt            IJWTRepository
	Logger         ILogger
}

func NewApp() Application {
	app := &Application{}
	app.Env = NewENV()
	app.PostgresClient = NewPostgresClient(app.Env.DatabaseURL, app.Env.MaxOpenConnection, app.Env.MaxIdleConnection)
	app.Harsher = NewPasswordHasher()
	app.Logger = NewEchoLogger()

	// JWT
	reader := commons.NewFileReader()
	privateKeyBytes, err := reader.ReadFile("private_key.pem")
	if err != nil {
		panic(errors.ErrInvalidLoadPrivateKey.Error())
	}

	publicKeyBytes, err := ioutil.ReadFile("public_key.pem")
	if err != nil {
		panic(errors.ErrInvalidLoadPublicKey.Error())
	}
	app.Jwt = NewJWTRSATokenRepository(privateKeyBytes, publicKeyBytes)

	app.Validator = NewCustomValidator()
	if err := app.Validator.RegisterValidation(commons.ValidatorPhoneNumber, commons.ValidatePhoneNumber); err != nil {
		panic(errors.ErrRegisterValidatorPhoneNumber.Error())
	}

	if err := app.Validator.RegisterValidation(commons.ValidatorPassword, commons.ValidatePassword); err != nil {
		panic(errors.ErrRegisterValidatorPassword.Error())
	}

	if err := app.Validator.RegisterValidation(commons.ValidatorFullName, commons.ValidationFullName); err != nil {
		panic(errors.ErrRegisterValidatorFullName.Error())
	}
	return *app
}
