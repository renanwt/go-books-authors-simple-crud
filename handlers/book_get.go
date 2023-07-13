package handlers

import (
	"BookStore/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func BookGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	book, err := models.BookGet(int64(id))
	if err != nil {
		log.Printf("Failed getting book: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	authors, err := models.GetAuthorIDByBookID(int64(id))
	if err != nil {
		log.Printf("Failed getting author IDs: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	authorsIDs := make([]int64, len(authors))
	for i, author := range authors {
		authorsIDs[i] = author.ID
	}

	response := models.BookResponse{
		ID:              book.ID,
		BookName:        book.BookName,
		Edition:         book.Edition,
		PublicationYear: book.PublicationYear,
		Authors:         authorsIDs,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
