package user_service

import (
	"context"
	"expert_systems_api/dto"
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/pkg/helper"
	"expert_systems_api/user/user_repo"

	"net/http"
)

type userService struct {
	ur user_repo.UserRepo
}

type UserService interface {
	SignIn(payload *dto.UserSignInPayload) (*helper.HTTPResponse, exception.Exception)
	SignUp(payload *dto.UserSignUpPayload) (*helper.HTTPResponse, exception.Exception)
	Modify(userId int, payload *dto.UserModifyPayload) (*helper.HTTPResponse, exception.Exception)
	Profile(userId int) (*helper.HTTPResponse, exception.Exception)
	ChangePassword(userId int, payload *dto.UserChangePassword) (*helper.HTTPResponse, exception.Exception)
	Authentication(next http.Handler) http.Handler
}

func NewUserService(ur user_repo.UserRepo) UserService {
	return &userService{
		ur: ur,
	}
}

// Authentication implements UserService.
func (us *userService) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearerToken := r.Header.Get("Authorization")

		user := &entity.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			w.WriteHeader(err.Status())
			w.Write(helper.ResponseJSON(err))
			return
		}

		u, err := us.ur.FetchByEmail(user.Email)

		if err != nil {
			w.WriteHeader(err.Status())
			w.Write(helper.ResponseJSON(err))
			return
		}

		ctx := context.WithValue(context.Background(), "user", u)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// ChangePassword implements UserService.
func (us *userService) ChangePassword(userId int, payload *dto.UserChangePassword) (*helper.HTTPResponse, exception.Exception) {

	user, err := us.ur.FetchById(userId)

	if err != nil {
		return nil, err
	}

	isValidPassword := user.CompareHashPassword(payload.OldPassword)

	if isValidPassword {
		return nil, exception.NewBadRequestError("invalid user")
	}

	if payload.NewPassword != payload.ConfirmNewPassword {
		return nil, exception.NewBadRequestError("password didn't match")
	}

	if err := us.ur.ChangePassword(userId, payload.NewPassword); err != nil {
		return nil, err
	}

	return &helper.HTTPResponse{
		Status:  http.StatusOK,
		Message: "password successfully changed",
		Data:    nil,
	}, nil
}

// Modify implements UserService.
func (us *userService) Modify(userId int, payload *dto.UserModifyPayload) (*helper.HTTPResponse, exception.Exception) {

	_, err := us.ur.FetchById(userId)

	if err != nil {
		return nil, err
	}

	if err := us.ur.Modify(userId, &entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
	}); err != nil {
		return nil, err
	}

	return &helper.HTTPResponse{
		Status:  http.StatusOK,
		Message: "user successfully modified",
		Data:    nil,
	}, nil
}

// Profile implements UserService.
func (us *userService) Profile(userId int) (*helper.HTTPResponse, exception.Exception) {

	u, err := us.ur.FetchById(userId)

	if err != nil {
		return nil, err
	}

	return &helper.HTTPResponse{
		Status:  http.StatusOK,
		Message: "user successfully fetched",
		Data: &dto.UserData{
			Id:        u.Id,
			FullName:  u.FullName,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
	}, nil
}

// SignIn implements UserService.
func (us *userService) SignIn(payload *dto.UserSignInPayload) (*helper.HTTPResponse, exception.Exception) {

	user, err := us.ur.FetchByEmail(payload.Email)

	if err != nil {
		return nil, err
	}

	isValidPassword := user.CompareHashPassword(payload.Password)

	if !isValidPassword {
		return nil, exception.NewBadRequestError("invalid user")
	}

	return &helper.HTTPResponse{
		Status:  http.StatusOK,
		Message: "user successfully fetched",
		Data: &dto.TokenString{
			Token: user.GenerateTokenString(),
		},
	}, nil
}

// SignUp implements UserService.
func (us *userService) SignUp(payload *dto.UserSignUpPayload) (*helper.HTTPResponse, exception.Exception) {

	u := &entity.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     "members",
	}

	u.GenerateHashPassword()

	if err := us.ur.Add(u); err != nil {
		return nil, err
	}

	return &helper.HTTPResponse{
		Status:  http.StatusCreated,
		Message: "user successfully created",
		Data:    nil,
	}, nil
}
