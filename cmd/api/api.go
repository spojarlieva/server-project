package main

import (
	"log"
	"net/http"
	"server/auth"
	"server/config"
	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"
	"server/utils"
)

type App struct {
	config        *config.Config
	handlers      handlers.Handlers
	authenticator *auth.JWTAuthenticator
}

func (a *App) start() error {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/app/static/assets"))))
	mux.Handle("/", http.FileServer(http.Dir("/app/static/pages")))
	mux.Handle("POST /users/register", utils.EnableCors(a.handlers.UserHandler.Register()))
	mux.Handle("POST /users/login", utils.EnableCors(a.handlers.UserHandler.Login()))
	mux.Handle("GET /events", utils.EnableCors(a.handlers.EventHandle.GetEvents()))
	mux.Handle("POST /events/register/{id}", utils.EnableCors(a.authenticator.Middleware(a.handlers.EventHandle.RegisterForEvents())))
	mux.Handle("DELETE /events/register/{id}", utils.EnableCors(a.authenticator.Middleware(a.handlers.EventHandle.UnregisterForEvents())))
	mux.Handle("GET /events/registered", utils.EnableCors(a.authenticator.Middleware(a.handlers.EventHandle.GetRegisteredEvents())))
	mux.Handle("GET /gallery/images", utils.EnableCors(a.handlers.EventHandle.GetGalleryImages()))

	return http.ListenAndServe(a.config.ServerAddr, mux)
}

func main() {
	conf := config.NewConfig()
	db, err := database.Connect(&conf.DatabaseConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	authenticator := auth.NewJWTAuthenticator(conf.AuthConfig.JwtSecret, conf.AuthConfig.JwtIssuer)

	app := App{
		config:        conf,
		authenticator: authenticator,
		handlers: handlers.Handlers{
			UserHandler: handlers.NewDefaultUserHandler(
				services.NewDefaultUserService(
					repositories.NewDefaultUserRepository(db),
					authenticator,
				),
			),
			EventHandle: handlers.NewDefaultEventHandler(
				services.NewDefaultEventService(
					repositories.NewDefaultEventRepository(db),
				),
			),
		},
	}

	err = app.start()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
