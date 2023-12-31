package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"jacc/config"
)

func main() {
	config, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", templ.Handler(Body(config.Name)).ServeHTTP)

	http.ListenAndServe(":3000", r)
}
