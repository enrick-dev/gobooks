package web

import (
	"encoding/json"
	"net/http"

	"github.com/enrick-dev/gobooks.git/intermal/service"
)

type BookHandlers struct {
	service *service.BookService
}

func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, "failed to get Books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book service.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.service.CreateBook(&book)
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder((w)).Encode(book)
}
