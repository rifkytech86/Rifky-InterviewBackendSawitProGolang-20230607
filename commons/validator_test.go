package commons

import (
	"encoding/json"
	"fmt"
	"github.com/SawitProRecruitment/UserService/errors"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func registerTestValidation() *validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation(ValidatorPhoneNumber, ValidatePhoneNumber); err != nil {
		fmt.Printf("unable to register custom validator: %s\n", err.Error())
	}
	if err := validate.RegisterValidation(ValidatorPassword, ValidatePassword); err != nil {
		panic(errors.ErrRegisterValidatorPassword.Error())
	}
	if err := validate.RegisterValidation(ValidatorFullName, ValidationFullName); err != nil {
		panic(errors.ErrRegisterValidatorPassword.Error())
	}
	return validate
}

func TestValidatePhoneNumber1(t *testing.T) {
	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		name    string
		args    args
		payload string
		want    string
	}{
		{
			name:    "test validation phone empty",
			payload: `{"phone_number": " ","password": "asdqwe1A@"}`,
			want:    `Key: 'LoginJSONBody.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'validationPhoneNumber' tag`,
		},
		{
			name:    "test validation phone wrong format",
			payload: `{"phone_number": "+62asdfasdas","password": "asdqwe1A@"}`,
			want:    `Key: 'LoginJSONBody.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'validationPhoneNumber' tag`,
		},
		{
			name:    "test validation not using prefix +62 ",
			payload: `{"phone_number": "8888888888888","password": "asdqwe1A@"}`,
			want:    `Key: 'LoginJSONBody.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'validationPhoneNumber' tag`,
		},
		{
			name:    "test validation less then 10 ",
			payload: `{"phone_number": "+6288888888","password": "asdqwe1A@"}`,
			want:    `Key: 'LoginJSONBody.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'validationPhoneNumber' tag`,
		},
		{
			name:    "test validation more then 13",
			payload: `{"phone_number": "+6288888888888888","password": "asdqwe1A@"}`,
			want:    `Key: 'LoginJSONBody.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'validationPhoneNumber' tag`,
		},
		{
			name:    "test validation more then 13",
			payload: `{"phone_number": "+622345678911111","password": "asdqwe1A@"}`,
			want:    `Key: 'LoginJSONBody.PhoneNumber' Error:Field validation for 'PhoneNumber' failed on the 'validationPhoneNumber' tag`,
		},
		{
			name:    "test validation phoneumber pass",
			payload: `{"phone_number": "+6223456789111","password": "asdqwe1A@"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validate := registerTestValidation()
			req := new(generated.LoginJSONBody)
			if err := json.Unmarshal([]byte(tt.payload), &req); err != nil {
				t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
			}
			validatorErrors := validate.Struct(req)
			if validatorErrors != nil {
				assert.Equal(t, tt.want, validatorErrors.Error(), "error got")
			} else {
				assert.Equal(t, nil, validatorErrors, "error got")
			}

		})
	}
}

func TestValidatePassword(t *testing.T) {
	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		name    string
		args    args
		payload string
		want    string
	}{
		{
			name:    "test validation password is empty",
			payload: `{"phone_number": "+6285722811111","password": ""}`,
			want:    `Key: 'LoginJSONBody.Password' Error:Field validation for 'Password' failed on the 'required' tag`,
		},
		{
			name:    "test validation password less then 6",
			payload: `{"phone_number": "+6285722811111","password": "123"}`,
			want:    `Key: 'LoginJSONBody.Password' Error:Field validation for 'Password' failed on the 'validationPassword' tag`,
		},
		{
			name:    "test validation password more then 64",
			payload: `{"phone_number": "+6285722811111","password": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}`,
			want:    `Key: 'LoginJSONBody.Password' Error:Field validation for 'Password' failed on the 'validationPassword' tag`,
		},
		{
			name:    "test validation success",
			payload: `{"phone_number": "+6285722811111","password": "asdqwe1A@"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validate := registerTestValidation()
			req := new(generated.LoginJSONBody)
			if err := json.Unmarshal([]byte(tt.payload), &req); err != nil {
				t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
			}
			validatorErrors := validate.Struct(req)
			if validatorErrors != nil {
				assert.Equal(t, tt.want, validatorErrors.Error(), "error got")
			} else {
				assert.Equal(t, nil, validatorErrors, "error got")
			}

		})
	}
}

func TestValidateFullName(t *testing.T) {
	type args struct {
		fl validator.FieldLevel
	}
	tests := []struct {
		name    string
		args    args
		payload string
		want    string
	}{
		{
			name:    "test validation full_name is empty",
			payload: `{"phone_number": "+6285722811111","password": "asdqwe1A@", "full_name": ""}`,
			want:    `Key: 'RegistrationJSONBody.FullName' Error:Field validation for 'FullName' failed on the 'required' tag`,
		},
		{
			name:    "test validation full_name less then 3",
			payload: `{"phone_number": "+6285722811111","password": "asdqwe1A@", "full_name": "af"}`,
			want:    `Key: 'RegistrationJSONBody.FullName' Error:Field validation for 'FullName' failed on the 'validationFullName' tag`,
		},
		{
			name:    "test validation full_name more then  60",
			payload: `{"phone_number": "+6285722811111","password": "asdqwe1A@", "full_name": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1"}`,
			want:    `Key: 'RegistrationJSONBody.FullName' Error:Field validation for 'FullName' failed on the 'validationFullName' tag`,
		},
		{
			name:    "test validation success",
			payload: `{"phone_number": "+6285722811111","password": "asdqwe1A@", "full_name": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validate := registerTestValidation()
			req := new(generated.RegistrationJSONBody)
			if err := json.Unmarshal([]byte(tt.payload), &req); err != nil {
				t.Errorf("failed to unmarshal lead to JSON: %v", err.Error())
			}
			validatorErrors := validate.Struct(req)
			if validatorErrors != nil {
				assert.Equal(t, tt.want, validatorErrors.Error(), "error got")
			} else {
				assert.Equal(t, nil, validatorErrors, "error got")
			}

		})
	}
}

func TestGetCustomMessage(t *testing.T) {
	type args struct {
		msgError string
		field    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test ValidatorPassword",
			args: args{
				msgError: ValidatorPassword,
				field:    "password",
			},
			want: fmt.Sprintf(errors.ErrorValidatorPassword, "password"),
		},
		{
			name: "test ValidatorPhoneNumber",
			args: args{
				msgError: ValidatorPhoneNumber,
				field:    "phone_number",
			},
			want: fmt.Sprintf(errors.ErrorValidatorPhoneNumber, "phone_number"),
		},
		{
			name: "test ValidatorFullName",
			args: args{
				msgError: ValidatorFullName,
				field:    "full_name",
			},
			want: fmt.Sprintf(errors.ErrorValidatorFullName, "full_name"),
		},
		{
			name: "test default",
			args: args{
				msgError: "default",
				field:    "default",
			},
			want: fmt.Sprintf(errors.ErrorDefaultValidator, "default", "default"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetCustomMessage(tt.args.msgError, tt.args.field), "GetCustomMessage(%v, %v)", tt.args.msgError, tt.args.field)
		})
	}
}
