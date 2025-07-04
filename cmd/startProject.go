package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/tutul/book-server/api/handler"
	"github.com/tutul/book-server/api/middleware"
	"github.com/tutul/book-server/infrastructure/persistence/inmemory"
	"github.com/tutul/book-server/service"
	"net/http"
)

func main() {
	var port string
	var enableAuth bool

	flag.StringVar(&port, "port", "8080", "Port for server")
	flag.BoolVar(&enableAuth, "auth", true, "Enable JWT authentication")
	flag.Parse()

	bookRepo := inmemory.NewInMemoryRepo()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	userRepo := inmemory.NewInMemoryUserRepo()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)

	handler.RegisterRoutes(r, bookHandler, userHandler, middleware.TokenAuth, enableAuth)

	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, r)
}
