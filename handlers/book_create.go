package handlers

import (
	"BookStore/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BookCreate(w http.ResponseWriter, r *http.Request) {
	var bookReq models.BookRequest

	err := json.NewDecoder(r.Body).Decode(&bookReq)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	book := models.Books{
		BookName:        bookReq.BookName,
		Edition:         bookReq.Edition,
		PublicationYear: bookReq.PublicationYear,
	}

	// Insert book and get the generated book ID
	bookID, err := models.BookInsert(book)
	if err != nil {
		log.Printf("Failed to insert book: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	for _, authorID := range bookReq.Authors {
		booksAuthors := models.BooksAuthors{
			AuthorID: authorID,
			BookID:   bookID,
		}

		_, err := models.BookAuthorsInsert(booksAuthors)
		if err != nil {
		}
	}

	resp := map[string]interface{}{
		"Message": fmt.Sprintf("Book successfully inserted! ID: %d", bookID),
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
