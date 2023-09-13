package bootstrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewENV(t *testing.T) {
	tests := []struct {
		name string
		want *ENV
	}{
		{
			name: "intial env",
			want: NewENV(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewENV(), "NewENV()")
		})
	}
}
