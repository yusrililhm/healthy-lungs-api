package user_repo

import (
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
)

type UserRepo interface {
	Add(user *entity.User) exception.Exception
	FetchByEmail(email string) (*entity.User, exception.Exception)
	FetchById(id int) (*entity.User, exception.Exception)
	Modify(id int, user *entity.User) exception.Exception
	ChangePassword(id int, newPassword string) exception.Exception
}
