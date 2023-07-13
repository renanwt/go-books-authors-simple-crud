package handlers

import (
	"BookStore/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func AuthorList(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Default values if page or limit are not provided or invalid
	defaultPage := 1
	defaultLimit := 10

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	authors, totalCount, err := models.AuthorGetAll(page, limit)
	if err != nil {
		log.Printf("Error obtaining author registers: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := make(map[string]interface{})
	response["page"] = page
	response["pageSize"] = limit
	response["totalCount"] = totalCount

	authorsList := make([]map[string]interface{}, 0)
	for _, author := range authors {
		books, err := models.GetBooksByAuthorID(author.ID)
		if err != nil {
			log.Printf("Failed getting books for author %s: %v", author.AuthorName, err)
			continue
		}

		authorData := map[string]interface{}{
			"author": author,
			"books":  books,
		}

		authorsList = append(authorsList, authorData)
	}

	response["authors"] = authorsList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
