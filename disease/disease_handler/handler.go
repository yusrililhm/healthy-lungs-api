package disease_handler

import "expert_systems_api/disease/disease_service"

type diseaseHandler struct {
	ds disease_service.DiseaseService
}

type DiseaseHandler interface {
}

func NewDiseaseHandler(ds disease_service.DiseaseService) DiseaseHandler {
	return &diseaseHandler{
		ds: ds,
	}
}
