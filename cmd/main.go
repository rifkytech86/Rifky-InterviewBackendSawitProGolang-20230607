package main

import (
	"github.com/SawitProRecruitment/UserService/bootstrap"
	"github.com/SawitProRecruitment/UserService/commons"
	"github.com/SawitProRecruitment/UserService/errors"
	"github.com/SawitProRecruitment/UserService/generated"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
)

var unRestricted = map[string]bool{
	"/login":        true,
	"/registration": true,
}

func customMiddleware(jwt bootstrap.IJWTRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			if isRestricted := unRestricted[path]; !isRestricted {
				authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
				if authHeader == "" {
					return commons.ErrorResponse(c, http.StatusForbidden, errors.ErrMissingAuthorizationHeader.Error())
				}

				splitToken := strings.Split(authHeader, "Bearer ")
				if len(splitToken) != 2 {
					return commons.ErrorResponse(c, http.StatusForbidden, errors.ErrInvalidTokenFormat.Error())
				}

				tokenString := splitToken[1]
				userID, err := jwt.ParserToken(tokenString)
				if err != nil {
					return commons.ErrorResponse(c, http.StatusForbidden, errors.ErrInvalidToken.Error())
				}
				if userID == 0 {
					return commons.ErrorResponse(c, http.StatusForbidden, errors.ErrInvalidToken.Error())
				}
				c.Set("userID", userID)
				return next(c)
			}
			return next(c)
		}
	}
}

func main() {

	app := bootstrap.NewApp()

	e := echo.New()
	// add middleware for handling validate authorization
	e.Use(customMiddleware(app.Jwt))

	var server generated.ServerInterface = newServer(app)
	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer(app bootstrap.Application) *handler.Server {
	userServiceRepo := repository.NewUserPointRepo(app.PostgresClient)
	opts := handler.NewServerOptions{
		UserServiceRepository: userServiceRepo,
		Validator:             app.Validator,
		Harsher:               app.Harsher,
		Env:                   app.Env,
		JWTRepository:         app.Jwt,
		Logger:                app.Logger,
	}

	return handler.NewServer(opts)
}
