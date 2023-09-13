package models

import "testing"

func TestUser_TableName(t *testing.T) {
	type fields struct {
		UserID          int
		UserPhoneNumber string
		UserFullName    string
		UserPassword    string
		UserLogged      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "initial table name",
			fields: fields{},
			want:   "users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := User{
				UserID:          tt.fields.UserID,
				UserPhoneNumber: tt.fields.UserPhoneNumber,
				UserFullName:    tt.fields.UserFullName,
				UserPassword:    tt.fields.UserPassword,
				UserLogged:      tt.fields.UserLogged,
			}
			if got := us.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
