package service_test

import (
	"expert_systems_api/dto"
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/user/user_repo/user_pg"
	"expert_systems_api/user/user_service"

	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var repoMock = user_pg.NewUserRepoMock()
var userService = user_service.NewUserService(repoMock)

var signUpPayload = &dto.UserSignUpPayload{}

var signInPayload = &dto.UserSignInPayload{
	Password: "secret",
}

var modifyPayload = &dto.UserModifyPayload{}

var changePasswordPayload = &dto.UserChangePassword{
	OldPassword:        "secret",
	NewPassword:        "newpassword",
	ConfirmNewPassword: "newpassword",
}

var hashPassword, _ = bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)

func TestUserProfileSuccess(t *testing.T) {
	user_pg.FetchById = func(id int) (*entity.User, exception.Exception) {
		return &entity.User{}, nil
	}

	res, err := userService.Profile(1)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.Status)
}

func TestUserProfileFailedNotFound(t *testing.T) {
	user_pg.FetchById = func(id int) (*entity.User, exception.Exception) {
		return nil, exception.NewNotFoundError("user not found")
	}

	res, err := userService.Profile(1)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestUserProfileFailedServerError(t *testing.T) {
	user_pg.FetchById = func(id int) (*entity.User, exception.Exception) {
		return nil, exception.NewInternalServerError("something went wrong")
	}

	res, err := userService.Profile(1)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSignUpSuccess(t *testing.T) {
	user_pg.Add = func(user *entity.User) exception.Exception {
		return nil
	}

	res, err := userService.SignUp(signUpPayload)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusCreated, res.Status)
}

func TestSignUpFailedServerError(t *testing.T) {
	user_pg.Add = func(user *entity.User) exception.Exception {
		return exception.NewInternalServerError("something went wrong")
	}

	res, err := userService.SignUp(signUpPayload)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, err.Status())
}

func TestSignInFailedNotFound(t *testing.T) {
	user_pg.FetchByEmail = func(email string) (*entity.User, exception.Exception) {
		return nil, exception.NewNotFoundError("user not found")
	}

	res, err := userService.SignIn(signInPayload)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestSignInSuccess(t *testing.T) {
	user_pg.FetchByEmail = func(email string) (*entity.User, exception.Exception) {
		return &entity.User{}, nil
	}

	res, err := userService.SignIn(signInPayload)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.Status)
}

func TestModifyFailedNotFound(t *testing.T) {
	user_pg.FetchById = func(id int) (*entity.User, exception.Exception) {
		return nil, exception.NewNotFoundError("user not found")
	}

	res, err := userService.Modify(1, modifyPayload)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestModifyFailedServerError(t *testing.T) {
	user_pg.FetchById = func(id int) (*entity.User, exception.Exception) {
		return nil, exception.NewNotFoundError("user not found")
	}

	user_pg.Modify = func(id int, user *entity.User) exception.Exception {
		return exception.NewInternalServerError("something went wrong")
	}

	res, err := userService.Modify(1, modifyPayload)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestChangePasswordFailedUserNotFound(t *testing.T) {

	user_pg.FetchById = func(id int) (*entity.User, exception.Exception) {
		return nil, exception.NewNotFoundError("user not found")
	}

	res, err := userService.ChangePassword(1, changePasswordPayload)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
}

func TestChangePasswordFailedPasswordNotFalid(t *testing.T) {

	user_pg.FetchById = func(id int) (*entity.User, exception.Exception) {
		return &entity.User{
			Password: string(hashPassword),
		}, nil
	}

	res, err := userService.ChangePassword(1, &dto.UserChangePassword{
		OldPassword: "mkamsask",
		NewPassword:        "momoo",
		ConfirmNewPassword: "mimi",
	})

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status())
}
