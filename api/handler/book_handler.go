package handler

import (
	"book-server/domain/entity"
	"book-server/service"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	Service *service.BookService
}

func NewBookHandler(s *service.BookService) *BookHandler {
	return &BookHandler{s}
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := h.Service.Create(book)
	if err != nil {
		http.Error(w, "Failed to create", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := h.Service.Get(id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) List(w http.ResponseWriter, r *http.Request) {
	books, _ := h.Service.List()
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updated, err := h.Service.Update(id, book)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := h.Service.Delete(id)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}
