package symtomp_handler

import (
	"expert_systems_api/pkg/helper"
	"expert_systems_api/symtomp/symtomp_service"
	"net/http"
)

type symtompHandler struct {
	ss symtomp_service.SymtompService
}

type SymtompHandler interface {
	Fetch(w http.ResponseWriter, r *http.Request)
}

func NewSymtompHandler(ss symtomp_service.SymtompService) SymtompHandler {
	return &symtompHandler{
		ss: ss,
	}
}

// Fetch implements SymtompHandler.
func (s *symtompHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	res, err := s.ss.Fetch()

	if err != nil {
		w.WriteHeader(err.Status())
		w.Write(helper.ResponseJSON(err))
		return
	}

	w.WriteHeader(res.Status)
	w.Write(helper.ResponseJSON(res))
}
