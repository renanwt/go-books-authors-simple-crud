package models

import (
	"BookStore/db"
	"fmt"
)

func AuthorGetAll(page int, limit int) (authorslist []Authors, totalCount int, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	offset := (page - 1) * limit
	rows, err := conn.Query(fmt.Sprintf("SELECT * FROM authors LIMIT %v OFFSET %v", limit, offset))
	if err != nil {
		return
	}

	for rows.Next() {
		var authors Authors

		err = rows.Scan(&authors.ID, &authors.AuthorName)
		if err != nil {
			continue
		}

		authorslist = append(authorslist, authors)
	}

	totalCountQuery := "SELECT COUNT(*) FROM authors"
	err = conn.QueryRow(totalCountQuery).Scan(&totalCount)
	if err != nil {
		return
	}

	return
}
