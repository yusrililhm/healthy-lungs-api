package dto

import "time"

type UserSignInPayload struct {
	Email    string `valid:"required~Email can't empty, email" json:"email"`
	Password string `valid:"required~Password can't empty" json:"password"`
}

type UserSignUpPayload struct {
	FullName string `valid:"required~Full name can't empty" json:"full_name"`
	Email    string `valid:"required~Email can't empty, email" json:"email"`
	Password string `valid:"required~Password can't empty" json:"password"`
}

type UserModifyPayload struct {
	FullName string `valid:"required~Full name can't empty" json:"full_name"`
	Email    string `valid:"required~Email can't empty, email" json:"email"`
}

type UserData struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TokenString struct {
	Token string `valid:"required~ can't empty" json:"token"`
}

type UserChangePassword struct {
	OldPassword        string `valid:"required~Old password can't empty" json:"old_password"`
	NewPassword        string `valid:"required~New password can't empty" json:"new_password"`
	ConfirmNewPassword string `valid:"required~Confirm new password can't empty" json:"confirm_new_password"`
}
