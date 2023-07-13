package models

import "BookStore/db"

func BookGet(id int64) (books Books, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM books WHERE id=$1`, id)

	err = row.Scan(&books.ID, &books.BookName, &books.Edition, &books.PublicationYear)

	return
}

func GetAuthorIDByBookID(bookID int64) (authors []Authors, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`
		SELECT authors.id
		FROM books_authors
		JOIN authors ON books_authors.author_id = authors.id
		WHERE books_authors.book_id = $1`, bookID)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var author Authors
		err = rows.Scan(&author.ID)
		if err != nil {
			return
		}
		authors = append(authors, author)
	}

	err = rows.Err()
	return
}
