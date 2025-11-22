package handler

import (
	"github.com/go-chi/chi"
	"rafael.br/simple-crud/internal/middleware"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/balance", GetBalance)
	})
}