package disease_service

import (
	"expert_systems_api/disease/disease_repo"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/pkg/helper"
	"net/http"
)

type diseaseService struct {
	dr disease_repo.DiseaseRepo
}

type DiseaseService interface {
	Fetch() (*helper.HTTPResponse, exception.Exception)
}

func NewDiseaseService(dr disease_repo.DiseaseRepo) DiseaseService {
	return &diseaseService{
		dr: dr,
	}
}

// Fetch implements DiseaseService.
func (d *diseaseService) Fetch() (*helper.HTTPResponse, exception.Exception) {

	diseases, err := d.dr.Fetch()

	if err != nil {
		return nil, err
	}

	return &helper.HTTPResponse{
		Status:  http.StatusOK,
		Message: "disease successfully fetched",
		Data:    diseases,
	}, nil
}
