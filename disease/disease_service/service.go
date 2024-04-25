package disease_service

import "expert_systems_api/disease/disease_repo"

type diseaseService struct {
	dr disease_repo.DiseaseRepo
}

type DiseaseService interface {
}

func NewDiseaseService(dr disease_repo.DiseaseRepo) DiseaseService {
	return &diseaseService{
		dr: dr,
	}
}
