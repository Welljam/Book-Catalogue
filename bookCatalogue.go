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

	fmt.Print("För att lägga till en bok skriv 1\nFör att ta bort en bok skriv 2\nFör att redigera en bok skriv 3\nFör att visa alla böcker skriv 4\n")

	var answer string
	fmt.Scan(&answer)
	if answer == "1" {
		addBook(db)
	} else if answer == "2" {
		deleteBook(db)
	} else if answer == "3" {
		changeBook(db)
	} else if answer == "4" {
		showBook(db)
	}
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

func addBook(db *sql.DB) {
	var userTitle, userAuthor, userGenre string
	fmt.Print("Lägg till en bok, ange titel, författare och genre\n")
	fmt.Scan(&userTitle, &userAuthor, &userGenre)

	query := `SELECT title FROM book WHERE title = $1`
	var title string
	err := db.QueryRow(query, userTitle).Scan(&title)
	if err != nil {
		if err == sql.ErrNoRows {
			book := newBook{userTitle, userAuthor, userGenre}
			pk := insertBook(db, book)
			fmt.Printf("ID = %d\n", pk)
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Boken finns redan")
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
	var userDeleteBook string
	fmt.Print("Ange titeln på boken du vill radera: ")
	fmt.Scan(&userDeleteBook)

	query := `DELETE FROM book WHERE title = $1`

	result, err := db.Exec(query, userDeleteBook)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected > 0 {
		fmt.Println("Boken raderades")
	} else {
		fmt.Println("Det finns ingen bok med denna titeln")
	}
}

func changeBook(db *sql.DB) {
	var userChangeBook string
	fmt.Print("Ange titeln på boken du vill redigera: ")
	fmt.Scan(&userChangeBook)

	query := `SELECT title, author, genre FROM book WHERE title = $1`
	row := db.QueryRow(query, userChangeBook)

	var (
		title  string
		author string
		genre  string
	)

	err := row.Scan(&title, &author, &genre)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Det finns ingen bok med denna titeln")
			return
		}
		log.Fatal(err)
	}

	fmt.Print("Ange en ny titel, författare och genre för boken: ")
	var newTitle, newAuthor, newGenre string
	fmt.Scan(&newTitle, &newAuthor, &newGenre)

	updateQuery := `UPDATE book SET title = $1, author = $2, genre = $3 WHERE title = $4`
	_, err = db.Exec(updateQuery, newTitle, newAuthor, newGenre, userChangeBook)
	if err != nil {
		log.Fatal(err)
	}
}

func showBook(db *sql.DB) {
	query := `SELECT * FROM book`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id     int
			title  string
			author string
			genre  string
		)

		err := rows.Scan(&id, &title, &author, &genre)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Book: %d, Title: %s, Author: %s, Genre: %s\n", id, title, author, genre)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
