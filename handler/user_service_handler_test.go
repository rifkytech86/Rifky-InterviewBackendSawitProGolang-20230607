package handler

import (
	originalError "errors"
	"github.com/SawitProRecruitment/UserService/bootstrap"
	"github.com/SawitProRecruitment/UserService/bootstrap/mocks"
	"github.com/SawitProRecruitment/UserService/commons"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/repository"
	mockRepository "github.com/SawitProRecruitment/UserService/repository/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func failedBind() echo.Context {
	e := echo.New()
	userErrJSON := `}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userErrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func validatorFailed() echo.Context {
	e := echo.New()
	userErrJSON := `{"password":"asdqwe1A@","phone_number":"+62"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userErrJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func TestServer_Login(t *testing.T) {

	type fields struct {
		UserServiceRepository repository.IUserServicePointRepository
		Validator             bootstrap.IValidator
		Harsher               bootstrap.IBcryptHasher
		Env                   *bootstrap.ENV
		Logger                bootstrap.ILogger
		JWTRepository         bootstrap.IJWTRSAToken
	}
	type respMockGetUserBYPhone struct {
		user *models.User
		err  error
	}
	type resMockJWT struct {
		token string
		err   error
	}

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name                   string
		fields                 fields
		args                   args
		respMockValidator      []bootstrap.ValidationError
		respMockGetUserBYPhone respMockGetUserBYPhone
		resMockVerifyPassword  error
		resMockJWT             resMockJWT
		resMocUpdateUsers      error
		wantErr                bool
	}{
		{
			name: "test failed bind",
			args: args{
				c: failedBind(),
			},
			wantErr: false,
		},
		{
			name: "test failed validator",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{
				{
					Field: "password",
					Error: "invalid password",
				},
			},
			wantErr: false,
		},
		{
			name: "test failed  get user by phone",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			respMockGetUserBYPhone: respMockGetUserBYPhone{
				err: commons.ErrorInternalServer,
			},
			wantErr: false,
		},
		{
			name: "test failed  verify",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			respMockGetUserBYPhone: respMockGetUserBYPhone{
				err: nil,
				user: &models.User{
					UserID: 1,
				},
			},
			resMockVerifyPassword: commons.ErrorInternalServer,
			wantErr:               false,
		},
		{
			name: "test failed generate token",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			respMockGetUserBYPhone: respMockGetUserBYPhone{
				err: nil,
				user: &models.User{
					UserID: 1,
				},
			},
			resMockVerifyPassword: nil,
			resMockJWT: resMockJWT{
				err: commons.ErrorInternalServer,
			},
			wantErr: false,
		},
		{
			name: "test failed update user",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			respMockGetUserBYPhone: respMockGetUserBYPhone{
				err: nil,
				user: &models.User{
					UserID: 1,
				},
			},
			resMockVerifyPassword: nil,
			resMockJWT: resMockJWT{
				err: nil,
			},
			resMocUpdateUsers: commons.ErrorInternalServer,
			wantErr:           false,
		},
		{
			name: "test success login handler",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			respMockGetUserBYPhone: respMockGetUserBYPhone{
				err: nil,
				user: &models.User{
					UserID: 1,
				},
			},
			resMockVerifyPassword: nil,
			resMockJWT: resMockJWT{
				err:   nil,
				token: "token",
			},
			resMocUpdateUsers: nil,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			s.Env = bootstrap.NewENV()
			s.Logger = bootstrap.NewEchoLogger()
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.respMockValidator)
			s.Validator = mockValidator

			mocksGetUserByPhone := new(mockRepository.IUserServicePointRepository)
			mocksGetUserByPhone.On("GetUserByPhone", mock.Anything, mock.Anything).Return(tt.respMockGetUserBYPhone.user, tt.respMockGetUserBYPhone.err)
			mocksGetUserByPhone.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(tt.resMocUpdateUsers)
			s.UserServiceRepository = mocksGetUserByPhone

			mockHasser2 := new(mocks.IBcryptHasher)
			mockHasser2.On("VerifyPassword", mock.Anything, mock.Anything).Return(tt.resMockVerifyPassword)
			s.Harsher = mockHasser2

			mockJWTRepository := new(mocks.IJWTRepository)
			mockJWTRepository.On("GenerateToken", mock.Anything, mock.Anything).Return(tt.resMockJWT.token, tt.resMockJWT.err)
			s.JWTRepository = mockJWTRepository

			if err := s.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Registration(t *testing.T) {
	type fields struct {
		UserServiceRepository repository.IUserServicePointRepository
		Validator             bootstrap.IValidator
		Harsher               bootstrap.IBcryptHasher
		Env                   *bootstrap.ENV
		Logger                bootstrap.ILogger
		JWTRepository         bootstrap.IJWTRSAToken
	}
	type resMockHasPassword struct {
		hasPassword string
		err         error
	}
	type respMockInsert struct {
		lastID int64
		err    error
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		respMockValidator  []bootstrap.ValidationError
		resMockHasPassword resMockHasPassword
		respMockInsert     respMockInsert
		wantErr            bool
	}{
		{
			name: "test registration error bind",
			args: args{
				c: failedBind(),
			},
			wantErr: false,
		},
		{
			name: "test registration error validate",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{
				{
					Field: "password",
					Error: "invalid password",
				},
			},
			wantErr: false,
		},
		{
			name: "test registration error hasser",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			resMockHasPassword: resMockHasPassword{
				err: commons.ErrorInternalServer,
			},
			wantErr: false,
		},
		{
			name: "test registration error insert user",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			resMockHasPassword: resMockHasPassword{
				err: nil,
			},
			respMockInsert: respMockInsert{
				err: commons.ErrorInternalServer,
			},
			wantErr: false,
		},
		{
			name: "test registration error insert no row effected",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			resMockHasPassword: resMockHasPassword{
				err: nil,
			},
			respMockInsert: respMockInsert{
				err:    nil,
				lastID: 0,
			},
			wantErr: false,
		},
		{
			name: "test registration success",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			resMockHasPassword: resMockHasPassword{
				err: nil,
			},
			respMockInsert: respMockInsert{
				err:    nil,
				lastID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			s.Env = bootstrap.NewENV()
			s.Logger = bootstrap.NewEchoLogger()
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.respMockValidator)
			s.Validator = mockValidator
			mockHasser2 := new(mocks.IBcryptHasher)
			mockHasser2.On("HashPassword", mock.Anything, mock.Anything).Return(tt.resMockHasPassword.hasPassword, tt.resMockHasPassword.err)
			s.Harsher = mockHasser2
			mocksGetUserByPhone := new(mockRepository.IUserServicePointRepository)
			mocksGetUserByPhone.On("InsetUser", mock.Anything, mock.Anything).Return(tt.respMockInsert.lastID, tt.respMockInsert.err)
			s.UserServiceRepository = mocksGetUserByPhone
			if err := s.Registration(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Registration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_GetMyProfile(t *testing.T) {
	setContext := validatorFailed()
	setContext.Set("userID", 1)
	type fields struct {
		UserServiceRepository repository.IUserServicePointRepository
		Validator             bootstrap.IValidator
		Harsher               bootstrap.IBcryptHasher
		Env                   *bootstrap.ENV
		Logger                bootstrap.ILogger
		JWTRepository         bootstrap.IJWTRSAToken
	}
	type respMockGetUserByID struct {
		user *models.User
		err  error
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		respMockGetUserByID respMockGetUserByID
		wantErr             bool
	}{
		{
			name: "test error get my profile bind",
			args: args{
				c: failedBind(),
			},
			wantErr: true,
		},
		{
			name: "test error get User by id",
			args: args{
				c: setContext,
			},
			respMockGetUserByID: respMockGetUserByID{
				err: commons.ErrorInternalServer,
			},
			wantErr: false,
		},
		{
			name: "test get profile success",
			args: args{
				c: setContext,
			},
			respMockGetUserByID: respMockGetUserByID{
				err: nil,
				user: &models.User{
					UserFullName:    "tester",
					UserPhoneNumber: "+6285722811111",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			s.Env = bootstrap.NewENV()
			s.Logger = bootstrap.NewEchoLogger()

			mocksGetUserByPhone := new(mockRepository.IUserServicePointRepository)
			mocksGetUserByPhone.On("GetUserByUserID", mock.Anything, mock.Anything).Return(tt.respMockGetUserByID.user, tt.respMockGetUserByID.err)
			s.UserServiceRepository = mocksGetUserByPhone

			if err := s.GetMyProfile(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetMyProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_UpdateProfile(t *testing.T) {
	setContext := validatorFailed()
	setContext.Set("userID", 2)

	type fields struct {
		UserServiceRepository repository.IUserServicePointRepository
		Validator             bootstrap.IValidator
		Harsher               bootstrap.IBcryptHasher
		Env                   *bootstrap.ENV
		Logger                bootstrap.ILogger
		JWTRepository         bootstrap.IJWTRSAToken
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		respMockValidator []bootstrap.ValidationError
		resMocUpdateUsers error
		wantErr           bool
	}{
		{
			name: "error Update Profile Bind",
			args: args{
				c: failedBind(),
			},
			wantErr: false,
		},
		{
			name: "error Update Profile validator",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{
				{
					Field: "password",
					Error: "invalid password",
				},
			},
			wantErr: false,
		},
		{
			name: "error Update profile cant get user id",
			args: args{
				c: validatorFailed(),
			},
			respMockValidator: []bootstrap.ValidationError{},
			wantErr:           false,
		},
		{
			name: "error Update profile update user with duplicate",
			args: args{
				c: setContext,
			},
			resMocUpdateUsers: originalError.New(commons.DuplicateKey),
			respMockValidator: []bootstrap.ValidationError{},
			wantErr:           false,
		},
		{
			name: "error Update profile update non duplicate",
			args: args{
				c: setContext,
			},
			resMocUpdateUsers: commons.ErrorInternalServer,
			respMockValidator: []bootstrap.ValidationError{},
			wantErr:           false,
		},
		{
			name: "success update profile",
			args: args{
				c: setContext,
			},
			resMocUpdateUsers: nil,
			respMockValidator: []bootstrap.ValidationError{},
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{}
			s.Env = bootstrap.NewENV()
			s.Logger = bootstrap.NewEchoLogger()
			mockValidator := new(mocks.IValidator)
			mockValidator.On("Struct", mock.Anything).Return(tt.respMockValidator)
			s.Validator = mockValidator
			mocksGetUserByPhone := new(mockRepository.IUserServicePointRepository)
			mocksGetUserByPhone.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(tt.resMocUpdateUsers)
			s.UserServiceRepository = mocksGetUserByPhone

			if err := s.UpdateProfile(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
