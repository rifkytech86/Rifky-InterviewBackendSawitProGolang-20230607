package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewCustomValidator(t *testing.T) {
	initNew := NewCustomValidator()
	tests := []struct {
		name string
		want IValidator
	}{
		{
			name: "new custom validator",
			want: initNew,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := initNew
			assert.NotNil(t, validator)

		})
	}
}

func Test_customValidator_RegisterValidation(t *testing.T) {
	type fields struct {
		validate *validator.Validate
	}
	type args struct {
		tag                      string
		fn                       validator.Func
		callValidationEvenIfNull []bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "register validator",
			fields: fields{
				validate: validator.New(),
			},
			args: args{
				tag: "test",
				fn: func(fl validator.FieldLevel) bool {
					return true
				},
				callValidationEvenIfNull: []bool{true},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mv := &customValidator{
				validate: tt.fields.validate,
			}
			if err := mv.RegisterValidation(tt.args.tag, tt.args.fn, tt.args.callValidationEvenIfNull...); (err != nil) != tt.wantErr {
				t.Errorf("RegisterValidation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type ExampleStruct struct {
	Username string `validate:"required"`
	Email    string `validate:"email"`
}

func Test_customValidator_Struct(t *testing.T) {
	type fields struct {
		validate *validator.Validate
	}
	type args struct {
		s interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []ValidationError
	}{
		{
			name: "test struct invalid",
			fields: fields{
				validate: validator.New(),
			},
			args: args{
				s: ExampleStruct{
					Username: "",
					Email:    "john@x.com",
				},
			},
			want: []ValidationError{{
				Field: "Username",
				Error: "required",
			}},
		},
		{
			name: "test struct valid",
			fields: fields{
				validate: validator.New(),
			},
			args: args{
				s: ExampleStruct{
					Username: "testere",
					Email:    "john@x.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mv := &customValidator{
				validate: tt.fields.validate,
			}
			if got := mv.Struct(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Struct() = %v, want %v", got, tt.want)
			}
		})
	}
}
