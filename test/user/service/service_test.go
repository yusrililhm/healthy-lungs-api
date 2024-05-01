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
)

var repoMock = user_pg.NewUserRepoMock()
var userService = user_service.NewUserService(repoMock)

var signUpPayload = &dto.UserSignUpPayload{}
var signInPayload = &dto.UserSignInPayload{}

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
