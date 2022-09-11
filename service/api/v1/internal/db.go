package internal

import (
	"database/sql"
	"log"
	"net/url"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func QueryBuilder(parameters url.Values) string {
	mainQuery := "SELECT * FROM book INNER JOIN genre ON book.genre_id = genre.id INNER JOIN author ON book.author_id = author.id "
	wherePieces := make([]string, 0)
	authors, contains := parameters["authors"]
	if contains {
		authorsPieces := make([]string, 0)
		for _, value := range strings.Split(authors[0], ",") {
			authorsPieces = append(authorsPieces, " author_id = "+value)
		}
		wherePieces = append(wherePieces, "("+strings.Join(authorsPieces[:], " OR ")+")")
	}
	genre, contains := parameters["genres"]
	if contains {
		genrePieces := make([]string, 0)
		for _, value := range strings.Split(genre[0], ",") {
			genrePieces = append(genrePieces, " genre_id = "+value)
		}
		wherePieces = append(wherePieces, "("+strings.Join(genrePieces[:], " OR ")+")")
	}
	minYear, contains := parameters["min-year"]
	if contains {
		if minYear[0] == "0" {
			minYear[0] = "1800"
		}
		wherePieces = append(wherePieces, " year_published >= "+minYear[0])
	}
	maxYear, contains := parameters["max-year"]
	if contains {
		if maxYear[0] == "0" {
			maxYear[0] = "2100"
		}
		wherePieces = append(wherePieces, " year_published <= "+maxYear[0])
	}
	minPages, contains := parameters["min-pages"]
	if contains {
		if minPages[0] == "0" {
			minPages[0] = "1"
		}
		wherePieces = append(wherePieces, " pages >= "+minPages[0])
	}
	maxPages, contains := parameters["max-pages"]
	if contains {
		if maxPages[0] == "0" {
			maxPages[0] = "10000"
		}
		wherePieces = append(wherePieces, " pages <= "+maxPages[0])
	}
	if len(wherePieces) > 0 {
		mainQuery += "WHERE" + strings.Join(wherePieces[:], " AND ")
	}
	limit, contains := parameters["limit"]
	if contains {
		mainQuery = mainQuery + "LIMIT " + limit[0]
	}
	return mainQuery
}
