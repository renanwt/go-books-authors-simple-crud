package models

import "BookStore/db"

func AuthorGet(id int64) (authors Authors, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM authors WHERE id=$1`, id)
	err = row.Scan(&authors.ID, &authors.AuthorName)

	return
}

func AuthorGetByName(authorname string) (authors Authors, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM authors WHERE author_name=$1`, authorname)
	err = row.Scan(&authors.ID, &authors.AuthorName)

	return
}

func GetBooksByAuthorID(authorID int64) (books []Books, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT books.id, books.book_name, books.edition, books.publication_year 
		FROM books
		INNER JOIN books_authors ON books.id = books_authors.book_id
		WHERE books_authors.author_id = $1`, authorID)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book Books
		err = rows.Scan(&book.ID, &book.BookName, &book.Edition, &book.PublicationYear)
		if err != nil {
			return
		}
		books = append(books, book)
	}

	err = rows.Err()
	return
}
