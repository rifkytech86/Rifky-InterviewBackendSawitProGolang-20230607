package bootstrap

import "github.com/go-playground/validator/v10"

// IValidator defines the interface for a validator
//
//go:generate mockery --name IValidator
type IValidator interface {
	Struct(s interface{}) []ValidationError
	RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error
}

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// MyValidator is an implementation of the Validator interface that wraps the validator.Validate
type customValidator struct {
	validate *validator.Validate
}

func (mv *customValidator) Struct(s interface{}) []ValidationError {
	var validationErrors []ValidationError

	err := mv.validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Field: err.Field(),
				Error: err.Tag(),
			})
		}
	}

	return validationErrors
}

func (mv *customValidator) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return mv.validate.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func NewCustomValidator() IValidator {
	return &customValidator{
		validate: validator.New(),
	}
}
