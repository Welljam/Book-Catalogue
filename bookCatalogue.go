package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type newBook struct {
	Title  string
	Author string
	Genre  string
}

func main() {
	connStr := "postgres://postgres:password@localhost:5432/bookCatalogue?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createBook(db)

	var i, j, k string
	fmt.Print("Lägg till en bok, ange titel, författare och genre\n")
	fmt.Scan(&i, &j, &k)

	book := newBook{i, j, k}
	pk := insertBook(db, book)
	fmt.Printf("ID = %d\n", pk)
}

func createBook(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS book(
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL, 
		author VARCHAR(255) NOT NULL,
		genre VARCHAR(255) NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertBook(db *sql.DB, book newBook) int {
	query := `INSERT INTO book (title, author, genre)
		VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, book.Title, book.Author, book.Genre).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func deleteBook(db *sql.DB) {

}
