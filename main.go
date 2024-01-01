package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"jacc/config"
	"jacc/pdf"
)

func main() {
	config, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/html", templ.Handler(Index(config)).ServeHTTP)
	r.Get("/", servePdf())

	// Serve static files from the public directory
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	http.ListenAndServe(":3000", r)
}

func servePdf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pdfFile, err := pdf.GeneratePdf(
			"http://localhost:3000/html",
			"public/cv.pdf",
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.ServeFile(w, r, pdfFile)
	}
}
