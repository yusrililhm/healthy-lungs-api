package user_handler

import (
	"encoding/json"
	"net/http"

	"expert_systems_api/dto"
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/pkg/helper"
	"expert_systems_api/user/user_service"
)

type userHandler struct {
	us user_service.UserService
}

type UserHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request)
	Modify(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(us user_service.UserService) UserHandler {
	return &userHandler{
		us: us,
	}
}

// ChangePassword implements UserHandler.
func (uh *userHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*entity.User)

	payload := &dto.UserChangePassword{}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		unprocessableEntityError := exception.NewUnprocessableEntityError("invalid JSON body request")
		w.WriteHeader(unprocessableEntityError.Status())
		w.Write(helper.ResponseJSON(unprocessableEntityError))
		return
	}

	if e := helper.ValidationStruct(payload); e != nil {
		w.WriteHeader(e.Status())
		w.Write(helper.ResponseJSON(e))
		return
	}

	ur, err := uh.us.ChangePassword(user.Id, payload)

	if err != nil {
		w.WriteHeader(err.Status())
		w.Write(helper.ResponseJSON(err))
		return
	}

	w.WriteHeader(ur.Status)
	w.Write(helper.ResponseJSON(ur))
}

// Modify implements UserHandler.
func (uh *userHandler) Modify(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*entity.User)

	payload := &dto.UserModifyPayload{}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		unprocessableEntityError := exception.NewUnprocessableEntityError("invalid JSON body request")
		w.WriteHeader(unprocessableEntityError.Status())
		w.Write(helper.ResponseJSON(unprocessableEntityError))
		return
	}

	if e := helper.ValidationStruct(payload); e != nil {
		w.WriteHeader(e.Status())
		w.Write(helper.ResponseJSON(e))
		return
	}

	ur, err := uh.us.Modify(user.Id, payload)

	if err != nil {
		w.WriteHeader(err.Status())
		w.Write(helper.ResponseJSON(err))
		return
	}

	w.WriteHeader(ur.Status)
	w.Write(helper.ResponseJSON(ur))
}

// Profile implements UserHandler.
func (uh *userHandler) Profile(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(*entity.User)

	ur, err := uh.us.Profile(user.Id)

	if err != nil {
		w.WriteHeader(err.Status())
		w.Write(helper.ResponseJSON(err))
		return
	}

	w.WriteHeader(ur.Status)
	w.Write(helper.ResponseJSON(ur))
}

// SignIn implements UserHandler.
func (uh *userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	payload := &dto.UserSignInPayload{}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		unprocessableEntityError := exception.NewUnprocessableEntityError("invalid JSON body request")
		w.WriteHeader(unprocessableEntityError.Status())
		w.Write(helper.ResponseJSON(unprocessableEntityError))
		return
	}

	if e := helper.ValidationStruct(payload); e != nil {
		w.WriteHeader(e.Status())
		w.Write(helper.ResponseJSON(e))
		return
	}

	ur, err := uh.us.SignIn(payload)

	if err != nil {
		w.WriteHeader(err.Status())
		w.Write(helper.ResponseJSON(err))
		return
	}

	w.WriteHeader(ur.Status)
	w.Write(helper.ResponseJSON(ur))
}

// SignUp implements UserHandler.
func (uh *userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	payload := &dto.UserSignUpPayload{}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		unprocessableEntityError := exception.NewUnprocessableEntityError("invalid JSON body request")
		w.WriteHeader(unprocessableEntityError.Status())
		w.Write(helper.ResponseJSON(unprocessableEntityError))
		return
	}

	if e := helper.ValidationStruct(payload); e != nil {
		w.WriteHeader(e.Status())
		w.Write(helper.ResponseJSON(e))
		return
	}

	ur, err := uh.us.SignUp(payload)

	if err != nil {
		w.WriteHeader(err.Status())
		w.Write(helper.ResponseJSON(err))
		return
	}

	w.WriteHeader(ur.Status)
	w.Write(helper.ResponseJSON(ur))
}
