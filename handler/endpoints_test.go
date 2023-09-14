package handler

import (
	"github.com/SawitProRecruitment/UserService/bootstrap"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	"testing"
)

func TestServer_Hello(t *testing.T) {
	s := validatorFailed()
	type fields struct {
		UserServiceRepository repository.IUserServicePointRepository
		Validator             bootstrap.IValidator
		Harsher               bootstrap.IBcryptHasher
		Env                   *bootstrap.ENV
		Logger                bootstrap.ILogger
		JWTRepository         bootstrap.IJWTRSAToken
	}
	type args struct {
		ctx    echo.Context
		params generated.HelloParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "tester hello",
			args: args{
				ctx: s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UserServiceRepository: tt.fields.UserServiceRepository,
				Validator:             tt.fields.Validator,
				Harsher:               tt.fields.Harsher,
				Env:                   tt.fields.Env,
				Logger:                tt.fields.Logger,
				JWTRepository:         tt.fields.JWTRepository,
			}
			if err := s.Hello(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
