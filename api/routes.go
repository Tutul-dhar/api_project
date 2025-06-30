package api

import (
	"book-server/api/handler"
	"book-server/api/middleware"
	// "book-server/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func RegisterRoutes(r *chi.Mux, h *handler.BookHandler, auth *jwtauth.JWTAuth, enableAuth bool) {
	r.Route("/api/v1/books", func(r chi.Router) {
		if enableAuth {
			r.Use(jwtauth.Verifier(auth))
			r.Use(jwtauth.Authenticator(auth))
		}
		r.Post("/", h.Create)
		r.Get("/", h.List)
		r.Get("/{id}", h.Get)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})

	r.Get("/api/v1/get-token", middleware.GetTokenHandler)
}
