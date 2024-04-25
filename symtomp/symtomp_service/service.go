package symtomp_service

import (
	"expert_systems_api/pkg/exception"
	"expert_systems_api/pkg/helper"
	"expert_systems_api/symtomp/symtomp_repo"
	"net/http"
)

type symtompService struct {
	sr symtomp_repo.SymtompRepo
}

type SymtompService interface {
	Fetch() (*helper.HTTPResponse, exception.Exception)
}

func NewSymtompService(sr symtomp_repo.SymtompRepo) SymtompService {
	return &symtompService{
		sr: sr,
	}
}

// Fetch implements SymtompService.
func (s *symtompService) Fetch() (*helper.HTTPResponse, exception.Exception) {

	symtomps, err := s.sr.Fetch()

	if err != nil {
		return nil, err
	}

	return &helper.HTTPResponse{
		Status:  http.StatusOK,
		Message: "symtomps successfully fetched",
		Data:    symtomps,
	}, nil
}
