package handlers

import (
	"BookStore/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func BookDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rows, err := models.BookDelete(int64(id))
	if err != nil {
		log.Printf("Failed deleting book: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if rows > 1 {
		log.Printf("Error: %d books deleted", rows)
	}

	resp := map[string]any{
		"Message": "Data successfully deleted!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
