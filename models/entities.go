package models

type Books struct {
	ID              int64  `json:"id"`
	BookName        string `json:"book_name"`
	Edition         string `json:"edition"`
	PublicationYear string `json:"publication_year"`
}

type Authors struct {
	ID         int64  `json:"id"`
	AuthorName string `json:"Author_name"`
}

type BooksAuthors struct {
	BookID   int64 `json:"book_id"`
	AuthorID int64 `json:"author_id"`
}

type BookRequest struct {
	BookName        string  `json:"book_name"`
	Edition         string  `json:"edition"`
	PublicationYear string  `json:"publication_year"`
	Authors         []int64 `json:"authors"`
}

type BookResponse struct {
	ID              int64   `json:"id"`
	BookName        string  `json:"book_name"`
	Edition         string  `json:"edition"`
	PublicationYear string  `json:"publication_year"`
	Authors         []int64 `json:"authors"`
}
