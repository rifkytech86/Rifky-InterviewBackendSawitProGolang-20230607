package bootstrap

import (
	"github.com/SawitProRecruitment/UserService/commons"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_echoLogger_Log(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		message interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test Log",
			fields: fields{Logger: nil},
			args: args{
				message: "tester",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &echoLogger{
				Logger: tt.fields.Logger,
			}
			l.Log(tt.args.message)
		})
	}
}

func Test_echoLogger_Warning(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		message interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test Log warning",
			fields: fields{Logger: nil},
			args: args{
				message: "tester",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &echoLogger{
				Logger: tt.fields.Logger,
			}
			l.Warning(tt.args.message)
		})
	}
}

func Test_echoLogger_Info(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		message interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test Log Info",
			fields: fields{Logger: nil},
			args: args{
				message: "tester",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &echoLogger{
				Logger: tt.fields.Logger,
			}
			l.Info(tt.args.message)
		})
	}
}

func Test_echoLogger_Error(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		message interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test Log Error",
			fields: fields{Logger: nil},
			args: args{
				message: "tester",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &echoLogger{
				Logger: tt.fields.Logger,
			}
			l.Error(tt.args.message)
		})
	}
}

func Test_echoLogger_Danger(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		message interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test Log Danger",
			fields: fields{Logger: nil},
			args: args{
				message: "tester",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &echoLogger{
				Logger: tt.fields.Logger,
			}
			l.Danger(tt.args.message)
		})
	}
}

func Test_echoLogger_convertMessage(t *testing.T) {
	type fields struct {
		Logger *log.Logger
	}
	type args struct {
		message interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test convert message string",
			fields: fields{Logger: nil},
			args: args{
				message: "tester",
			},
			want: "tester:",
		},
		{
			name:   "test convert message type error",
			fields: fields{Logger: nil},
			args: args{
				message: commons.ErrorInvalidRequest,
			},
			want: "invalid request:",
		},
		{
			name:   "test convert message type int",
			fields: fields{Logger: nil},
			args: args{
				message: 1,
			},
			want: "1:",
		},
		{
			name:   "test convert message default",
			fields: fields{Logger: nil},
			args: args{
				message: float64(6.5),
			},
			want: "6.5:",
		},
		{
			name:   "test convert message default with error encode",
			fields: fields{Logger: nil},
			args: args{
				message: map[string]interface{}{
					"tester": make(chan int),
				},
			},
			want: "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &echoLogger{
				Logger: tt.fields.Logger,
			}
			assert.Equalf(t, tt.want, l.convertMessage(tt.args.message), "convertMessage(%v)", tt.args.message)
		})
	}
}

func TestNewEchoLogger(t *testing.T) {
	tests := []struct {
		name string
		want ILogger
	}{
		{
			name: "initial newEchoLogger",
			want: NewEchoLogger(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewEchoLogger(), "NewEchoLogger()")
		})
	}
}
