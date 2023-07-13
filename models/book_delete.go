package models

import "BookStore/db"

func BookDelete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// Delete rows from the books_authors table
	_, err = conn.Exec("DELETE FROM books_authors WHERE book_id = $1", id)
	if err != nil {
		return 0, err
	}

	// Delete row from the books table
	res, err := conn.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
