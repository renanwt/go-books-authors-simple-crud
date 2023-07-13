package handlers

import (
	"BookStore/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func AuthorGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	authors, err := models.AuthorGet(int64(id))
	if err != nil {
		log.Printf("Failed getting author: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	books, err := models.GetBooksByAuthorID(int64(id))
	if err != nil {
		log.Printf("Failed getting books: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"author": authors,
		"books":  books,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AuthorGetByName(w http.ResponseWriter, r *http.Request) {
	authorname := chi.URLParam(r, "authorname")

	author, err := models.AuthorGetByName(authorname)
	if err != nil {
		log.Printf("Failed getting author: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	books, err := models.GetBooksByAuthorID(author.ID)
	if err != nil {
		log.Printf("Failed getting books by author ID: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"author": author,
		"books":  books,
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
