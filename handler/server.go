package handler

import (
	"github.com/SawitProRecruitment/UserService/bootstrap"
	"github.com/SawitProRecruitment/UserService/repository"
)

type Server struct {
	UserServiceRepository repository.IUserServicePointRepository
	Validator             bootstrap.IValidator
	Harsher               bootstrap.IBcryptHasher
	Env                   *bootstrap.ENV
	Logger                bootstrap.ILogger
	JWTRepository         bootstrap.IJWTRSAToken
}

type NewServerOptions struct {
	UserServiceRepository repository.IUserServicePointRepository
	Validator             bootstrap.IValidator
	Harsher               bootstrap.IBcryptHasher
	Env                   *bootstrap.ENV
	Logger                bootstrap.ILogger
	JWTRepository         bootstrap.IJWTRSAToken
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		UserServiceRepository: opts.UserServiceRepository,
		Validator:             opts.Validator,
		Harsher:               opts.Harsher,
		Env:                   opts.Env,
		Logger:                opts.Logger,
		JWTRepository:         opts.JWTRepository,
	}
}
