package bootstrap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var p = `LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlCUEFJQkFBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VXNjbGFFKzlaUUg5Q2VpOGIxcUVmCnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUUpCQUw4ZjRBMUlDSWEvQ2ZmdWR3TGMKNzRCdCtwOXg0TEZaZXMwdHdtV3Vha3hub3NaV0w4eVpSTUJpRmI4a25VL0hwb3piTnNxMmN1ZU9wKzVWdGRXNApiTlVDSVFENm9JdWxqcHdrZTFGY1VPaldnaXRQSjNnbFBma3NHVFBhdFYwYnJJVVI5d0loQVBOanJ1enB4ckhsCkUxRmJxeGtUNFZ5bWhCOU1HazU0Wk1jWnVjSmZOcjBUQWlFQWhML3UxOVZPdlVBWVd6Wjc3Y3JxMTdWSFBTcXoKUlhsZjd2TnJpdEg1ZGdjQ0lRRHR5QmFPdUxuNDlIOFIvZ2ZEZ1V1cjg3YWl5UHZ1YStxeEpXMzQrb0tFNXdJZwpQbG1KYXZsbW9jUG4rTkVRdGhLcTZuZFVYRGpXTTlTbktQQTVlUDZSUEs0PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==`
var pub = `LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZ3d0RRWUpLb1pJaHZjTkFRRUJCUUFEU3dBd1NBSkJBTzVIKytVM0xrWC91SlRvRHhWN01CUURXSTdGU0l0VQpzY2xhRSs5WlFIOUNlaThiMXFFZnJxR0hSVDVWUis4c3UxVWtCUVpZTER3MnN3RTVWbjg5c0ZVQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==`

func Test_jwtRSATokenRepository_GenerateToken(t *testing.T) {

	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	type args struct {
		userID      int
		expiredTime int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test generate token error decode string",
			fields: fields{
				publicKey:  []byte("c2RmYXNkZmFz"),
				privateKey: []byte("--234234"),
			},
			args: args{
				userID:      1,
				expiredTime: 1,
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "test generate token Parse RSA PEM",
			fields: fields{
				publicKey:  []byte("c2RmYXNkZmFz"),
				privateKey: []byte("c2RmYXNkZmFz"),
			},
			args: args{
				userID:      1,
				expiredTime: 1,
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "test generate token Error JWT Claim",
			fields: fields{
				publicKey:  []byte("c2RmYXNkZmFz"),
				privateKey: []byte(p),
			},
			args: args{
				userID:      1,
				expiredTime: 1,
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			name: "test generate token success",
			fields: fields{
				publicKey:  []byte("c2RmYXNkZmFz"),
				privateKey: []byte(p),
			},
			args: args{
				userID:      1,
				expiredTime: 1,
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtRSATokenRepository{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := j.GenerateToken(tt.args.userID, tt.args.expiredTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GenerateToken(%v, %v)", tt.args.userID, tt.args.expiredTime)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GenerateToken(%v, %v)", tt.args.userID, tt.args.expiredTime)
		})
	}
}

func Test_jwtRSATokenRepository_ParserToken(t *testing.T) {
	type fields struct {
		publicKey  []byte
		privateKey []byte
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "tester parser token failed decode",
			fields: fields{
				publicKey: []byte("testsersfsd"),
			},
			args: args{
				tokenString: "tester",
			},
			want: 0,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "tester parser token failed parser key PEM ",
			fields: fields{
				publicKey: []byte("c2RmYXNkZmFz"),
			},
			args: args{
				tokenString: "tester",
			},
			want: 0,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "tester parser token jwt parser ",
			fields: fields{
				publicKey: []byte(pub),
			},
			args: args{
				tokenString: "",
			},
			want: 0,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
		{
			name: "tester parser token success jwt parser",
			fields: fields{
				publicKey: []byte(pub),
			},
			args: args{
				tokenString: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjYxNzI5NzEsInN1YiI6MTEsInRva2VuX3V1aWQiOiI2ZjY0NTAzMS1mNjM0LTRhYmQtYTZkMS05OTU4OTZjOTcwNzkifQ.6AB504gcagHXnMDRvutJQfGvVuQCcg1vaTQk-wi0DcnvFeClXNJS4MVZr5Nw3X3DFRNcxN7kAt8zNJEwIRD8bA",
			},
			want: 11,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtRSATokenRepository{
				publicKey:  tt.fields.publicKey,
				privateKey: tt.fields.privateKey,
			}
			got, err := j.ParserToken(tt.args.tokenString)
			if !tt.wantErr(t, err, fmt.Sprintf("ParserToken(%v)", tt.args.tokenString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParserToken(%v)", tt.args.tokenString)
		})
	}
}
