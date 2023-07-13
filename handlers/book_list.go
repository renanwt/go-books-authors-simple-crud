package handlers

import (
	"BookStore/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func BookList(w http.ResponseWriter, r *http.Request) {
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

	books, totalCount, err := models.BookGetAll(page, limit)
	if err != nil {
		log.Printf("Error getting books: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := make(map[string]interface{})
	response["page"] = page
	response["pageSize"] = limit
	response["totalCount"] = totalCount

	booksList := make([]models.BookResponse, 0)
	for _, book := range books {
		authors, err := models.GetAuthorIDByBookID(book.ID)
		if err != nil {
			log.Printf("Failed getting authors IDs for book %s: %v", book.BookName, err)
			continue
		}

		authorsIDs := make([]int64, len(authors))
		for i, author := range authors {
			authorsIDs[i] = author.ID
		}

		bookData := models.BookResponse{
			ID:              book.ID,
			BookName:        book.BookName,
			Edition:         book.Edition,
			PublicationYear: book.PublicationYear,
			Authors:         authorsIDs,
		}

		booksList = append(booksList, bookData)
	}

	response["books"] = booksList

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
