package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/tutul/book-server/api/middleware"
	"github.com/tutul/book-server/domain/entity"
	"github.com/tutul/book-server/service"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created, err := h.UserService.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.UserService.Login(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func RegisterRoutes(r *chi.Mux, bookHandler *BookHandler, userHandler *UserHandler, auth *jwtauth.JWTAuth, enableAuth bool) {
	r.Route("/api/v1/books", func(r chi.Router) {
		// Use only Basic Authentication for GET /api/v1/books
		r.With(middleware.BasicAuthMiddleware(userHandler.UserService)).Get("/", bookHandler.List)

		// Apply JWT Authentication for other endpoints if enableAuth is true
		if enableAuth {
			r.With(middleware.JWTMiddleware(auth)).Post("/", bookHandler.Create)
			r.With(middleware.JWTMiddleware(auth)).Get("/{id}", bookHandler.Get)
			r.With(middleware.JWTMiddleware(auth)).Put("/{id}", bookHandler.Update)
			r.With(middleware.JWTMiddleware(auth)).Delete("/{id}", bookHandler.Delete)
		} else {
			r.Post("/", bookHandler.Create)
			r.Get("/{id}", bookHandler.Get)
			r.Put("/{id}", bookHandler.Update)
			r.Delete("/{id}", bookHandler.Delete)
		}
	})

	r.Get("/api/v1/get-token", middleware.GetTokenHandler(userHandler.UserService))

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
	})
}
