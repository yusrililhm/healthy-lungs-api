package app

import (
	"expert_systems_api/disease/disease_handler"
	"expert_systems_api/disease/disease_repo/disease_pg"
	"expert_systems_api/disease/disease_service"
	"expert_systems_api/infra/config"
	"expert_systems_api/infra/db"
	"expert_systems_api/symtomp/symtomp_handler"
	"expert_systems_api/symtomp/symtomp_repo/symtomp_pg"
	"expert_systems_api/symtomp/symtomp_service"
	"expert_systems_api/user/user_handler"
	"expert_systems_api/user/user_repo/user_pg"
	"expert_systems_api/user/user_service"

	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartApplication() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	config.LoadEnv()

	db.InitializeDatabase()
	db := db.GetDbInstance()

	// dependency injection
	ur := user_pg.NewUserRepo(db)
	us := user_service.NewUserService(ur)
	uh := user_handler.NewUserHandler(us)

	sr := symtomp_pg.NewSymtompPg(db)
	ss := symtomp_service.NewSymtompService(sr)
	sh := symtomp_handler.NewSymtompHandler(ss)

	dr := disease_pg.NewDiseasePg(db)
	ds := disease_service.NewDiseaseService(dr)
	dh := disease_handler.NewDiseaseHandler(ds)

	r.Post("/user/signin", uh.SignIn)
	r.Post("/user/signup", uh.SignUp)

	r.Group(func(r chi.Router) {
		r.Use(us.Authentication)

		r.Get("/user", uh.Profile)
		r.Patch("/user", uh.Modify)
		r.Patch("/user/change-password", uh.ChangePassword)

		r.Get("/symtomp", sh.Fetch)

		r.Get("/disease", dh.Fetch)
	})

	log.Printf("[server is running] on port %s", config.AppConfig().AppPort)
	http.ListenAndServe(":"+config.AppConfig().AppPort, r)
}
