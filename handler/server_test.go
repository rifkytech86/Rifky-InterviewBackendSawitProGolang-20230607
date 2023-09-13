package handler

import (
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		opts NewServerOptions
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "initial server",
			args: args{
				opts: NewServerOptions{},
			},
			want: NewServer(NewServerOptions{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
