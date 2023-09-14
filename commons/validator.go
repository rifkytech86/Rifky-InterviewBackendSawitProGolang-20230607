package commons

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const (
	ValidatorPassword    = "validationPassword"
	ValidatorPhoneNumber = "validationPhoneNumber"
	ValidatorFullName    = "validationFullName"
	CountryCode          = "+62"
)

func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()
	if phoneNumber == "" {
		return false
	}

	if len(phoneNumber) < 3 || phoneNumber[0:3] != CountryCode {
		return false
	}

	if len(phoneNumber[3:len(phoneNumber)]) < 9 {
		return false
	}

	totalDigit := len(phoneNumber[3:len(phoneNumber)]) + 1
	if totalDigit > 13 {
		return false
	}

	return true
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if password == "" {
		return false
	}

	if len(password) < 6 {
		return false
	}
	if len(password) > 63 {
		return false
	}
	if !checkHasSpecialCharNumberCapital(password) {
		return false
	}

	return true
}

func ValidationFullName(fl validator.FieldLevel) bool {
	fullName := fl.Field().String()
	if fullName == "" {
		return false
	}

	if len(fullName) < 3 || len(fullName) > 60 {
		return false
	}

	return true
}

func checkHasSpecialCharNumberCapital(password string) bool {
	hasSpecialChar := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(password)
	hasNumber := regexp.MustCompile(`\d`).MatchString(password)
	hasCapital := regexp.MustCompile(`[A-Z]`).MatchString(password)

	return hasSpecialChar && hasNumber && hasCapital
}

func GetCustomMessage(msgError string, field string) string {
	if msgError == ValidatorPassword {
		customErrorMessage := fmt.Sprintf(ErrorValidatorPassword, field)
		return customErrorMessage
	}

	if msgError == ValidatorPhoneNumber {
		customErrorMessage := fmt.Sprintf(ErrorValidatorPhoneNumber, field)
		return customErrorMessage
	}

	if msgError == ValidatorFullName {
		customErrorMessage := fmt.Sprintf(ErrorValidatorFullName, field)
		return customErrorMessage
	}

	return fmt.Sprintf(ErrorDefaultValidator, field, msgError)

}
