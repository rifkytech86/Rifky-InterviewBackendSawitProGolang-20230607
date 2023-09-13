package models

type User struct {
	UserID          int    `json:"user_id"`
	UserPhoneNumber string `json:"user_phone_number"`
	UserFullName    string `json:"user_full_name"`
	UserPassword    string `json:"user_password"`
	UserLogged      int    `json:"user_logged"`
}

func (User) TableName() string {
	return "users"
}
