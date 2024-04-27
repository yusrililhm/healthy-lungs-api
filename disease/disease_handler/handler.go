package disease_handler

import (
	"expert_systems_api/disease/disease_service"
	"expert_systems_api/pkg/helper"
	"net/http"
)

type diseaseHandler struct {
	ds disease_service.DiseaseService
}

type DiseaseHandler interface {
	Fetch(w http.ResponseWriter, r *http.Request)
}

func NewDiseaseHandler(ds disease_service.DiseaseService) DiseaseHandler {
	return &diseaseHandler{
		ds: ds,
	}
}

// Fetch implements DiseaseHandler.
func (d *diseaseHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	res, err := d.ds.Fetch()

	if err != nil {
		w.WriteHeader(err.Status())
		w.Write(helper.ResponseJSON(err))
		return
	}

	w.WriteHeader(res.Status)
	w.Write(helper.ResponseJSON(res))
}
