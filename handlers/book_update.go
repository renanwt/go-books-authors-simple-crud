package handlers

import (
	"BookStore/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func BookUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var books models.Books

	err = json.NewDecoder(r.Body).Decode(&books)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rows, err := models.BookUpdate(int64(id), books)
	if err != nil {
		log.Printf("Failed updating book: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if rows > 1 {
		log.Printf("Error: %d books updated", rows)
	}

	resp := map[string]any{
		"Message": "Book successfully updated!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
