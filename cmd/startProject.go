package main

import (
	"book-server/api"
	"book-server/api/handler"
	myMiddleware "book-server/api/middleware"
	"book-server/infrastructure/persistence/inmemory"
	"book-server/service"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)

	api.RegisterRoutes(r, bookHandler, myMiddleware.TokenAuth, enableAuth)


	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, r)
}
