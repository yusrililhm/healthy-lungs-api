package user_service

import (
	"expert_systems_api/dto"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/pkg/helper"
	"net/http"
)

type userServiceMock struct {
}

var (
	Authentication func(next http.Handler) http.Handler
	ChangePassword func(userId int, payload *dto.UserChangePassword) (*helper.HTTPResponse, exception.Exception)
	Modify         func(userId int, payload *dto.UserModifyPayload) (*helper.HTTPResponse, exception.Exception)
	Profile        func(userId int) (*helper.HTTPResponse, exception.Exception)
	SignIn         func(payload *dto.UserSignInPayload) (*helper.HTTPResponse, exception.Exception)
	SignUp         func(payload *dto.UserSignUpPayload) (*helper.HTTPResponse, exception.Exception)
)

func NewUserServiceMock() UserService {
	return &userServiceMock{}
}

// Authentication implements UserService.
func (u *userServiceMock) Authentication(next http.Handler) http.Handler {
	return Authentication(next)
}

// ChangePassword implements UserService.
func (u *userServiceMock) ChangePassword(userId int, payload *dto.UserChangePassword) (*helper.HTTPResponse, exception.Exception) {
	return ChangePassword(userId, payload)
}

// Modify implements UserService.
func (u *userServiceMock) Modify(userId int, payload *dto.UserModifyPayload) (*helper.HTTPResponse, exception.Exception) {
	return Modify(userId, payload)
}

// Profile implements UserService.
func (u *userServiceMock) Profile(userId int) (*helper.HTTPResponse, exception.Exception) {
	return Profile(userId)
}

// SignIn implements UserService.
func (u *userServiceMock) SignIn(payload *dto.UserSignInPayload) (*helper.HTTPResponse, exception.Exception) {
	return SignIn(payload)
}

// SignUp implements UserService.
func (u *userServiceMock) SignUp(payload *dto.UserSignUpPayload) (*helper.HTTPResponse, exception.Exception) {
	return SignUp(payload)
}
