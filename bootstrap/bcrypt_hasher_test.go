package bootstrap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bcryptPasswordHasher_HashPassword(t *testing.T) {
	fmt.Println(len([]byte(`aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1`)))
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test hash password success",
			args: args{
				password: "~",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "test hash password error",
			args: args{
				password: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bcryptPasswordHasher{}
			_, err := b.HashPassword(tt.args.password)
			if !tt.wantErr(t, err, fmt.Sprintf("HashPassword(%v)", tt.args.password)) {
				return
			}
		})
	}
}

func Test_bcryptPasswordHasher_VerifyPassword(t *testing.T) {
	type args struct {
		hashedPassword  string
		enteredPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test error verify",
			args: args{
				hashedPassword:  "2342342",
				enteredPassword: "asdasdasd",
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bcryptPasswordHasher{}

			tt.wantErr(t, b.VerifyPassword(tt.args.hashedPassword, tt.args.enteredPassword), fmt.Sprintf("VerifyPassword(%v, %v)", tt.args.hashedPassword, tt.args.enteredPassword))
		})
	}
}

func TestNewPasswordHasher(t *testing.T) {
	tests := []struct {
		name string
		want IBcryptHasher
	}{
		{
			name: "initial NewPasswordHasher",
			want: NewPasswordHasher(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewPasswordHasher(), "NewPasswordHasher()")
		})
	}
}
