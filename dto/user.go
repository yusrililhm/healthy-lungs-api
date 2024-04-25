package dto

import "time"

type UserSignInPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpPayload struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModifyPayload struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserData struct {
	Id        int
	FullName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserChangePassword struct {
	OldPassword        string `json:"old_password"`
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}
