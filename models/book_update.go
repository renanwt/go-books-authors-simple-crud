package models

import "BookStore/db"

func BookUpdate(id int64, books Books) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE books SET book_name=$2, edition=$3, publication_year=$4 WHERE id=$1`, id,
		books.BookName, books.Edition, books.PublicationYear)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
