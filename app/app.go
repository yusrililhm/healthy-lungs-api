package app

import (
	"expert_systems_api/infra/config"
	"expert_systems_api/infra/db"
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

	r.Post("/user/signin", uh.SignIn)
	r.Post("/user/signup", uh.SignUp)
	r.Get("/user", uh.Profile)
	r.Patch("/user", uh.Modify)

	log.Printf("[server is running] on port %s", config.AppConfig().AppPort)
	http.ListenAndServe(":"+config.AppConfig().AppPort, r)
}
