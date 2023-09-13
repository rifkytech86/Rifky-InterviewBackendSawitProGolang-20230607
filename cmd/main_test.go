package main

import (
	"github.com/SawitProRecruitment/UserService/bootstrap"
	mockBootstrap "github.com/SawitProRecruitment/UserService/bootstrap/mocks"
	"github.com/SawitProRecruitment/UserService/commons"
	"github.com/SawitProRecruitment/UserService/errors"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/SawitProRecruitment/UserService/repository/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func registerValidator(h handler.Server) {
	if err := h.Validator.RegisterValidation(commons.ValidatorPhoneNumber, commons.ValidatePhoneNumber); err != nil {
		panic(errors.ErrRegisterValidatorPhoneNumber.Error())
	}

	if err := h.Validator.RegisterValidation(commons.ValidatorPassword, commons.ValidatePassword); err != nil {
		panic(errors.ErrRegisterValidatorPassword.Error())
	}

	if err := h.Validator.RegisterValidation(commons.ValidatorFullName, commons.ValidationFullName); err != nil {
		panic(errors.ErrRegisterValidatorFullName.Error())
	}

}
func Test_Endoint(t *testing.T) {
	e := echo.New()
	// Wrong parameters
	userJSON := `{"password":"Jon Snow","phone_number":"jon@labstack.com"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handler.Server{}
	h.Env = bootstrap.NewENV()
	h.Validator = bootstrap.NewCustomValidator()
	h.Logger = bootstrap.NewEchoLogger()
	registerValidator(h)
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	userJSON = `{"password":"Jon Snow","phone_number":"jon@labstack.com"}`
	reqBind := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	//reqBind.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recBind := httptest.NewRecorder()
	cBind := e.NewContext(reqBind, recBind)
	hBind := handler.Server{}
	hBind.Env = bootstrap.NewENV()
	hBind.Validator = bootstrap.NewCustomValidator()
	hBind.Logger = bootstrap.NewEchoLogger()
	registerValidator(h)
	if assert.NoError(t, hBind.Login(cBind)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}

	// invalid get user
	userJSON = `{"password":"asdqwe1A@","phone_number":"+6285722811111"}`
	req2 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)
	h2 := handler.Server{}
	h2.Env = bootstrap.NewENV()
	h2.Validator = bootstrap.NewCustomValidator()
	h2.Logger = bootstrap.NewEchoLogger()
	registerValidator(h2)
	mocksUserService := new(mocks.IUserServicePointRepository)
	mocksUserService.On("GetUserByPhone", mock.Anything, mock.Anything).Return(nil, errors.ErrorInvalidRequest)
	h2.UserServiceRepository = mocksUserService
	if assert.NoError(t, h2.Login(c2)) {
		assert.Equal(t, http.StatusBadRequest, rec2.Code)
	}

	// invalid get user
	userJSON = `{"password":"asdqwe1A@","phone_number":"+6285722811111"}`
	req3 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)
	h3 := handler.Server{}
	h3.Env = bootstrap.NewENV()
	h3.Validator = bootstrap.NewCustomValidator()
	h3.Logger = bootstrap.NewEchoLogger()
	registerValidator(h3)
	mocksUserService = new(mocks.IUserServicePointRepository)
	mocksUserService.On("GetUserByPhone", mock.Anything, mock.Anything).Return(&models.User{}, nil)
	h3.UserServiceRepository = mocksUserService
	mockHasser := new(mockBootstrap.IBcryptHasher)
	mockHasser.On("VerifyPassword", mock.Anything, mock.Anything).Return(errors.ErrorInvalidRequest)
	h3.Harsher = mockHasser
	if assert.NoError(t, h3.Login(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}

	userJSON = `{"password":"asdqwe1A@","phone_number":"+6285722811111"}`
	req4 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e.NewContext(req4, rec4)
	h4 := handler.Server{}
	h4.Env = bootstrap.NewENV()
	h4.Validator = bootstrap.NewCustomValidator()
	h4.Logger = bootstrap.NewEchoLogger()
	registerValidator(h4)
	mocksUserService = new(mocks.IUserServicePointRepository)
	mocksUserService.On("GetUserByPhone", mock.Anything, mock.Anything).Return(&models.User{}, nil)
	h4.UserServiceRepository = mocksUserService
	mockHasser = new(mockBootstrap.IBcryptHasher)
	mockHasser.On("VerifyPassword", mock.Anything, mock.Anything).Return(nil)
	h4.Harsher = mockHasser
	mockJWTRepository := new(mockBootstrap.IJWTRepository)
	mockJWTRepository.On("GenerateToken", mock.Anything, mock.Anything).Return("", errors.ErrorInvalidRequest)
	h4.JWTRepository = mockJWTRepository
	if assert.NoError(t, h4.Login(c4)) {
		assert.Equal(t, http.StatusInternalServerError, rec4.Code)
	}

	userJSON = `{"password":"asdqwe1A@","phone_number":"+6285722811111"}`
	req5 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req5.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec5 := httptest.NewRecorder()
	c5 := e.NewContext(req5, rec5)
	h5 := handler.Server{}
	h5.Env = bootstrap.NewENV()
	h5.Validator = bootstrap.NewCustomValidator()
	h5.Logger = bootstrap.NewEchoLogger()
	registerValidator(h5)
	mocksUserService = new(mocks.IUserServicePointRepository)
	mocksUserService.On("GetUserByPhone", mock.Anything, mock.Anything).Return(&models.User{}, nil)
	mocksUserService.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(errors.ErrorInvalidRequest)
	h5.UserServiceRepository = mocksUserService
	mockHasser = new(mockBootstrap.IBcryptHasher)
	mockHasser.On("VerifyPassword", mock.Anything, mock.Anything).Return(nil)
	h5.Harsher = mockHasser
	mockJWTRepository = new(mockBootstrap.IJWTRepository)
	mockJWTRepository.On("GenerateToken", mock.Anything, mock.Anything).Return("token", nil)
	h5.JWTRepository = mockJWTRepository
	if assert.NoError(t, h5.Login(c5)) {
		assert.Equal(t, http.StatusInternalServerError, rec5.Code)
	}

	userJSON = `{"password":"asdqwe1A@","phone_number":"+6285722811111"}`
	req6 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(userJSON))
	req6.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec6 := httptest.NewRecorder()
	c6 := e.NewContext(req6, rec6)
	h6 := handler.Server{}
	h6.Env = bootstrap.NewENV()
	h6.Validator = bootstrap.NewCustomValidator()
	h6.Logger = bootstrap.NewEchoLogger()
	registerValidator(h6)
	mocksUserService = new(mocks.IUserServicePointRepository)
	mocksUserService.On("GetUserByPhone", mock.Anything, mock.Anything).Return(&models.User{
		UserID: 1,
	}, nil)
	mocksUserService.On("UpdateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	h6.UserServiceRepository = mocksUserService
	mockHasser = new(mockBootstrap.IBcryptHasher)
	mockHasser.On("VerifyPassword", mock.Anything, mock.Anything).Return(nil)
	h6.Harsher = mockHasser
	mockJWTRepository = new(mockBootstrap.IJWTRepository)
	mockJWTRepository.On("GenerateToken", mock.Anything, mock.Anything).Return("token", nil)
	h6.JWTRepository = mockJWTRepository
	if assert.NoError(t, h6.Login(c6)) {
		assert.Equal(t, http.StatusOK, rec6.Code)
	}
}
