package disease_repo

import (
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
)

type DiseaseRepo interface {
	Fetch() ([]*entity.Disease, exception.Exception)
}
