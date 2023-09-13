package bootstrap

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"runtime"
)

//go:generate mockery --name ILogger
type ILogger interface {
	Log(message interface{})
	Warning(message interface{})
	Info(message interface{})
	Danger(message interface{})
	Error(message interface{})
}

func NewEchoLogger() ILogger {
	return &echoLogger{}
}

type echoLogger struct {
	*log.Logger
}

func (l *echoLogger) Log(message interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Infof("%s:%d - %s", file, line, l.convertMessage(message))
}

func (l *echoLogger) Warning(message interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Warnf("%s:%d - %s", file, line, l.convertMessage(message))

}

func (l *echoLogger) Info(message interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Infof("%s:%d - %s", file, line, l.convertMessage(message))

}

func (l *echoLogger) Error(message interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Errorf("%s:%d - %s", file, line, l.convertMessage(message))

}

func (l *echoLogger) Danger(message interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Debugf("%s:%d - %s", file, line, l.convertMessage(message))
}

func (l *echoLogger) convertMessage(message interface{}) string {
	switch msg := message.(type) {
	case string:
		return fmt.Sprintf("%s:", msg)
	case error:
		return fmt.Sprintf("%s:", msg.Error())
	case int:
		return fmt.Sprintf("%d:", msg)
	default:
		res, err := json.Marshal(msg)
		if err != nil {
			return "-"
		}

		return fmt.Sprintf("%s:", string(res))
	}
}
