package models

import (
	"BookStore/db"
	"fmt"
)

func BookGetAll(page int, limit int) (bookslist []Books, totalCount int, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	offset := (page - 1) * limit
	rows, err := conn.Query(fmt.Sprintf("SELECT * FROM books LIMIT %v OFFSET %v", limit, offset))
	if err != nil {
		return
	}

	for rows.Next() {
		var books Books

		err = rows.Scan(&books.ID, &books.BookName, &books.Edition, &books.PublicationYear)
		if err != nil {
			continue
		}

		bookslist = append(bookslist, books)
	}

	totalCountQuery := "SELECT COUNT(*) FROM books"
	err = conn.QueryRow(totalCountQuery).Scan(&totalCount)
	if err != nil {
		return
	}

	return
}
