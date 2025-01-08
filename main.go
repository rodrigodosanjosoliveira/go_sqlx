package main

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var tables = `
CREATE TABLE authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);

CREATE TABLE books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    author_id INTEGER,
    published_year INTEGER,
    genre TEXT,
    FOREIGN KEY (author_id) REFERENCES authors(id)
);

CREATE TABLE members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    join_date TEXT NOT NULL DEFAULT CURRENT_DATE
);
`

type Author struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type Book struct {
	ID            int            `db:"id"`
	Title         string         `db:"title"`
	AuthorID      int            `db:"author_id"`
	PublishedYear int            `db:"published_year"`
	Genre         sql.NullString `db:"genre"`
}

type Member struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	JoinDate string `db:"join_date"`
}

func main() {
	db, err := sqlx.Connect("sqlite3", "sqlx_demo.db")
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(tables)

	db.MustExec("INSERT INTO authors (name,email) VALUES ($1,$2)", "J.K. Rowling", "jkrowling@gmail.com")

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO authors (name,email) VALUES ($1,$2)", "George R.R. Martin", "georgerrmartin@gmail.com")
	tx.MustExec("INSERT INTO books (title,author_id,published_year,genre) VALUES ($1,$2,$3,$4)", "A Game of Thrones", 2, 1996, "Fantasy")
	tx.MustExec("INSERT INTO books (title,author_id,published_year,genre) VALUES ($1,$2,$3,$4)", "A Clash of Kings", 2, 1998, "Fantasy")
	tx.MustExec("INSERT INTO members (name,email) VALUES ($1,$2)", "John Doe", "johndoe@gmail.com")
	tx.Commit()
}
