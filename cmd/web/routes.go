package main

import (
	"net/http"

	"github.com/vinay1668/gobookings/pkg/config"
	"github.com/vinay1668/gobookings/pkg/handlers"

	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
    //    mux:= pat.New()

    //    mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//    mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()

	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}