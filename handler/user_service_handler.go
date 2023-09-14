package handler

import (
	"context"
	"github.com/SawitProRecruitment/UserService/commons"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

// Login is used for user authentication.
func (s *Server) Login(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(s.Env.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(generated.LoginJSONBody)
	if err := c.Bind(req); err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, commons.ErrorBindRequest.Error())
	}

	validatorErrors := s.Validator.Struct(req)
	if len(validatorErrors) > 0 {
		var errorMessages []string
		for _, err := range validatorErrors {
			errorMessages = append(errorMessages, commons.GetCustomMessage(err.Error, err.Field))
		}
		s.Logger.Error(errorMessages)
		return commons.ErrorResponses(c, http.StatusBadRequest, commons.ErrorInvalidRequest.Error(), errorMessages)
	}

	// get user from databases
	user, err := s.UserServiceRepository.GetUserByPhone(ctx, req.PhoneNumber)
	if err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, commons.ErrorInvalidRequest.Error())
	}

	err = s.Harsher.VerifyPassword(user.UserPassword, req.Password)
	if err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, commons.ErrorInvalidRequest.Error())
	}

	// generate token jwt rs256
	token, err := s.JWTRepository.GenerateToken(user.UserID, s.Env.ExpiredAuthTime)
	if err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusInternalServerError, commons.ErrorInternalServer.Error())
	}

	fieldUpdate := map[string]interface{}{
		"user_logged": user.UserLogged + 1,
	}

	err = s.UserServiceRepository.UpdateUser(ctx, user.UserID, fieldUpdate)
	if err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusInternalServerError, commons.ErrorInternalServer.Error())
	}

	resp := generated.LoginResponse{
		Data: &struct {
			AuthJwt *string `json:"auth_jwt,omitempty"`
			UserId  *int    `json:"user_id,omitempty"`
		}{},
	}
	resp.Data.UserId = &user.UserID
	resp.Data.AuthJwt = &token
	return commons.SuccessResponse(c, http.StatusOK, resp.Data)
}

// Registration is used for registration users
func (s *Server) Registration(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(s.Env.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(generated.RegistrationJSONBody)
	if err := c.Bind(req); err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// validation request login
	validatorErrors := s.Validator.Struct(req)
	if len(validatorErrors) > 0 {
		var errorMessages []string
		for _, err := range validatorErrors {
			errorMessages = append(errorMessages, commons.GetCustomMessage(err.Error, err.Field))
		}
		s.Logger.Error(errorMessages)
		return commons.ErrorResponses(c, http.StatusBadRequest, commons.ErrorInvalidRequest.Error(), errorMessages)
	}
	hasPassword, err := s.Harsher.HashPassword(req.Password)
	if err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, commons.ErrorInvalidRequest.Error())
	}
	payloadDataInsert := models.User{
		UserFullName:    req.FullName,
		UserPhoneNumber: req.PhoneNumber,
		UserPassword:    hasPassword,
		UserLogged:      0,
	}
	lastInsertID, err := s.UserServiceRepository.InsetUser(ctx, payloadDataInsert)
	if err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if lastInsertID == 0 {
		return commons.ErrorResponse(c, http.StatusBadRequest, commons.ErrorInternalServer.Error())
	}
	userID := int(lastInsertID)
	resp := generated.RegistrationResponse{
		Data: &struct {
			UserId *int `json:"user_id,omitempty"`
		}{},
	}
	resp.Data.UserId = &userID

	return commons.SuccessResponse(c, http.StatusCreated, resp.Data)
}

// GetMyProfile is handler for get data user by authorization
func (s *Server) GetMyProfile(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(s.Env.ContextTimeOut)*time.Second)
	defer cancel()
	iUserID := c.Get("userID")
	userID, ok := iUserID.(int)
	if !ok {
		s.Logger.Error(commons.ErrorGetUserID)
		return commons.ErrorGetUserID
	}
	user, err := s.UserServiceRepository.GetUserByUserID(ctx, userID)
	if err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, commons.ErrorInvalidRequest.Error())
	}
	resp := generated.GetMyProfileResponse{
		Data: &struct {
			FullName    *string `json:"full_name,omitempty"`
			PhoneNumber *string `json:"phone_number,omitempty"`
		}{},
	}
	resp.Data.FullName = &user.UserFullName
	resp.Data.PhoneNumber = &user.UserPhoneNumber
	return commons.SuccessResponse(c, http.StatusOK, resp.Data)
}

// UpdateProfile is handler for update profile by authorization
func (s *Server) UpdateProfile(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(s.Env.ContextTimeOut)*time.Second)
	defer cancel()

	req := new(generated.UpdateProfileJSONRequestBody)
	if err := c.Bind(req); err != nil {
		s.Logger.Error(err)
		return commons.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	validatorErrors := s.Validator.Struct(req)
	if len(validatorErrors) > 0 {
		var errorMessages []string
		for _, err := range validatorErrors {
			errorMessages = append(errorMessages, commons.GetCustomMessage(err.Error, err.Field))
		}
		return commons.ErrorResponses(c, http.StatusBadRequest, commons.ErrorInvalidRequest.Error(), errorMessages)
	}

	iUserID := c.Get("userID")
	userID, ok := iUserID.(int)
	if !ok {
		return commons.ErrorResponse(c, http.StatusBadRequest, commons.ErrorGetUserID.Error())
	}

	fieldUpdate := map[string]interface{}{
		"user_phone_number": req.PhoneNumber,
		"user_full_name":    req.FullName,
	}
	err := s.UserServiceRepository.UpdateUser(ctx, userID, fieldUpdate)
	if err != nil {
		s.Logger.Error(err)
		if strings.Contains(err.Error(), commons.DuplicateKey) {
			return commons.ErrorResponse(c, http.StatusConflict, commons.ErrorPhoneNumberAlreadyExist.Error())
		}
		return commons.ErrorResponse(c, http.StatusInternalServerError, commons.ErrorInternalServer.Error())
	}

	resp := generated.ResponseUpdateProfile{
		Data: &struct {
			UserId *int `json:"user_id,omitempty"`
		}{},
	}
	resp.Data.UserId = &userID
	return commons.SuccessResponse(c, http.StatusOK, resp.Data)

}
