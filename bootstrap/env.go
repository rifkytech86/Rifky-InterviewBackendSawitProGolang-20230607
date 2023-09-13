package bootstrap

import (
	"os"
	"strconv"
)

type ENV struct {
	DatabaseURL       string `json:"DATABASE_URL"`
	ContextTimeOut    int    `json:"CONTEXT_TIMEOUT"`
	MaxOpenConnection int    `json:"MAX_OPEN_CONNECTION"`
	MaxIdleConnection int    `json:"MAX_IDLE_CONNECTION"`
	ExpiredAuthTime   int    `json:"EXPIRED_AUTH_TIME"`
}

func NewENV() *ENV {
	env := ENV{}
	env.DatabaseURL = os.Getenv("DATABASE_URL")

	contextTimeOutStr := os.Getenv("CONTEXT_TIMEOUT")
	contextTime, err := strconv.Atoi(contextTimeOutStr)
	if err != nil {
		contextTime = 20
	}
	env.ContextTimeOut = contextTime

	maxOpenStr := os.Getenv("MAX_OPEN_CONNECTION")
	maxOpenConnections, err := strconv.Atoi(maxOpenStr)
	if err != nil {
		maxOpenConnections = 10
	}
	env.MaxOpenConnection = maxOpenConnections

	maxIdleConnectionsStr := os.Getenv("MAX_IDLE_CONNECTION")
	maxIdleConn, err := strconv.Atoi(maxIdleConnectionsStr)
	if err != nil {
		maxIdleConn = 5
	}
	env.MaxIdleConnection = maxIdleConn

	expAuthTimeStr := os.Getenv("EXPIRED_AUTH_TIME")
	expAuthTime, err := strconv.Atoi(expAuthTimeStr)
	if err != nil {
		expAuthTime = 1
	}
	env.ExpiredAuthTime = expAuthTime
	return &env
}
