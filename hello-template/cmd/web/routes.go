package main

import (
	"net/http"

	"github.com/calvarado2004/hello-template/pkg/config"
	"github.com/calvarado2004/hello-template/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	//mux := pat.New()
	//mux.Get("/hello-world", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/hello-world", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
