package main

import (
	"net/http"

	"github.com/ItaloG/go-rate-limiter-challenge/configs"
	"github.com/ItaloG/go-rate-limiter-challenge/pkg/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middlewares.RateLimitMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		w.Write([]byte("<b>HELLO</b>"))
	})

	http.ListenAndServe(":8080", r)
}
