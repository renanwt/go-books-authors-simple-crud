package models

import "BookStore/db"

func BookInsert(books Books) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO books (book_name, edition, publication_year) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, books.BookName, books.Edition, books.PublicationYear).Scan(&id)
	if err != nil {
		return 0, err
	}

	return
}

func BookAuthorsInsert(books_authors BooksAuthors) (author_id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO books_authors (book_id, author_id) VALUES ($1, $2)`

	err = conn.QueryRow(sql, books_authors.BookID, books_authors.AuthorID).Scan(&author_id)

	return
}
