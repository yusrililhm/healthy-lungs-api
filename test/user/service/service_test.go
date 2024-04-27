package service_test

import (
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
