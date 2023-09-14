package bootstrap

import (
	"github.com/SawitProRecruitment/UserService/commons"
	"io/ioutil"
)

type Application struct {
	Env            *ENV
	PostgresClient *PostgresClient
	Validator      IValidator
	Harsher        IBcryptHasher
	Jwt            IJWTRSAToken
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
		panic(commons.ErrInvalidLoadPrivateKey.Error())
	}

	publicKeyBytes, err := ioutil.ReadFile("public_key.pem")
	if err != nil {
		panic(commons.ErrInvalidLoadPublicKey.Error())
	}
	app.Jwt = NewJWTRSAToken(privateKeyBytes, publicKeyBytes)

	app.Validator = NewCustomValidator()
	if err := app.Validator.RegisterValidation(commons.ValidatorPhoneNumber, commons.ValidatePhoneNumber); err != nil {
		panic(commons.ErrRegisterValidatorPhoneNumber.Error())
	}

	if err := app.Validator.RegisterValidation(commons.ValidatorPassword, commons.ValidatePassword); err != nil {
		panic(commons.ErrRegisterValidatorPassword.Error())
	}

	if err := app.Validator.RegisterValidation(commons.ValidatorFullName, commons.ValidationFullName); err != nil {
		panic(commons.ErrRegisterValidatorFullName.Error())
	}
	return *app
}
