package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	dbHost     = ""
	dbPort     = 5432
	dbUser     = ""
	dbPassword = ""
	dbName     = ""
)

func main() {
	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Open the CSV file
	csvFile, err := os.Open("populate/['your db file name to be populated'].csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read the first row from the CSV file
	row, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Extract the authors from the first row
	authors := strings.Split(row[0], ";")

	// Prepare the INSERT statement
	stmt, err := db.Prepare("INSERT INTO authors (author_name) VALUES ($1)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	fmt.Println("Loading... please wait.")

	// Insert each author into the table
	totalAuthors := len(authors)
	progressIndicator := NewProgressIndicator(totalAuthors)
	batchCount := 0
	for i, author := range authors {
		_, err := stmt.Exec(author)
		if err != nil {
			log.Fatal(err)
		}

		progressIndicator.Update(i + 1)

		// Check if the batch count has reached 50,000
		if (i+1)%50000 == 0 {
			progressIndicator.Complete()
			batchCount++
			fmt.Printf("Batch: %d, Loading next batch... please wait.\n", batchCount)

			// Reset the progress indicator for the next batch
			progressIndicator = NewProgressIndicator(totalAuthors)
		}
	}
	progressIndicator.Complete()

	fmt.Println("Data inserted successfully!")
}

// ProgressIndicator is used to display the progress of a task
type ProgressIndicator struct {
	total    int
	current  int
	previous int
}

// NewProgressIndicator creates a new instance of ProgressIndicator
func NewProgressIndicator(total int) *ProgressIndicator {
	return &ProgressIndicator{
		total:    total,
		current:  0,
		previous: 0,
	}
}

// Update updates the progress indicator with the current progress
func (pi *ProgressIndicator) Update(current int) {
	pi.current = current

	if pi.current != pi.previous {
		if pi.current%50000 == 0 {
			fmt.Printf("\rBatch: %d, Progress: %d/%d", pi.current/50000, pi.current, pi.total)
		} else {
			fmt.Printf("\rBatch: %d, Progress: %d", pi.current/50000, pi.current)
		}
		pi.previous = pi.current
	}
}

// Complete marks the progress indicator as complete
func (pi *ProgressIndicator) Complete() {
	fmt.Println()
}
