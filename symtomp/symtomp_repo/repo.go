package symtomp_repo

import (
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
)

type SymtompRepo interface {
	Fetch() ([]*entity.Symtomp, exception.Exception)
}
