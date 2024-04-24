package user_pg

import (
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/repository/user_repo"
)

type userRepoMock struct {
}

var (
	Add            func(user *entity.User) exception.Exception
	ChangePassword func(id int, newPassword string) exception.Exception
	FetchByEmail   func(email string) (*entity.User, exception.Exception)
	FetchById      func(id int) (*entity.User, exception.Exception)
	Modify         func(id int, user *entity.User) exception.Exception
)

func NewUserRepoMock() user_repo.UserRepo {
	return &userRepoMock{}
}

// Add implements user_repo.UserRepo.
func (u *userRepoMock) Add(user *entity.User) exception.Exception {
	return Add(user)
}

// ChangePassword implements user_repo.UserRepo.
func (u *userRepoMock) ChangePassword(id int, newPassword string) exception.Exception {
	return ChangePassword(id, newPassword)
}

// FetchByEmail implements user_repo.UserRepo.
func (u *userRepoMock) FetchByEmail(email string) (*entity.User, exception.Exception) {
	return FetchByEmail(email)
}

// FetchById implements user_repo.UserRepo.
func (u *userRepoMock) FetchById(id int) (*entity.User, exception.Exception) {
	return FetchById(id)
}

// Modify implements user_repo.UserRepo.
func (u *userRepoMock) Modify(id int, user *entity.User) exception.Exception {
	return Modify(id, user)
}
