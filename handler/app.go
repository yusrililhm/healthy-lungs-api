package handler

import (
	"expert_systems_api/handler/user_handler"
	"expert_systems_api/infra/db"
	"expert_systems_api/repository/user_repo/user_pg"
	"expert_systems_api/service/user_service"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func StartApplication()  {
	r := chi.NewRouter()

	db := db.GetDbInstance()

	// dependency injection
	ur := user_pg.NewUserRepo(db)
	us := user_service.NewUserService(ur)
	uh := user_handler.NewUserHandler(us)

	r.Post("/user/signin", uh.SignIn)
	r.Post("/user/signup", uh.SignUp)
	r.Get("/user", uh.Profile)
	r.Patch("/user", uh.Modify)

	http.ListenAndServe(":8080", r)
}
