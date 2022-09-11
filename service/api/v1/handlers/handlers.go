package handlers

import (
	"database/sql"
	"encoding/json"
	"librarian/entity"
	"librarian/internal"
	"log"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := make([]entity.Book, 0)
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	allowedParameters := []string{"authors", "genres", "min-pages", "max-pages", "min-year", "max-year", "limit"}
	paramters := r.URL.Query()
	for key, _ := range paramters {
		found := false
		for _, par := range allowedParameters {
			if key == par {
				found = true
				break
			}
		}
		if !found {
			w.WriteHeader(400)
			return
		}
	}
	DB := internal.ConnectDB()
	rows, err := DB.Query(internal.QueryBuilder(paramters))

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.YearPublished, &book.Rating, &book.Pages, &book.Genre.ID, &book.Author.ID, &book.Genre.ID, &book.Genre.Title, &book.Author.ID, &book.Author.FirstName, &book.Author.LastName); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	bookJson, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(bookJson)

	defer rows.Close()
	defer DB.Close()
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	DB := internal.ConnectDB()
	authors := make([]entity.Author, 0)
	rows, err := DB.Query("SELECT * FROM author")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var author entity.Author
		if err := rows.Scan(&author.ID, &author.FirstName, &author.LastName); err != nil {
			log.Fatal(err)
		}
		authors = append(authors, author)
	}
	authorJson, err := json.Marshal(authors)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(authorJson)
	defer rows.Close()
	defer DB.Close()
}

func GetGenres(w http.ResponseWriter, r *http.Request) {
	DB := internal.ConnectDB()
	genres := make([]entity.Genre, 0)
	rows, err := DB.Query("SELECT * FROM genre")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var genre entity.Genre
		if err := rows.Scan(&genre.ID, &genre.Title); err != nil {
			log.Fatal(err)
		}
		genres = append(genres, genre)
	}
	genreJson, err := json.Marshal(genres)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(genreJson)

	defer rows.Close()
	defer DB.Close()
}

func GetSizes(w http.ResponseWriter, r *http.Request) {
	DB := internal.ConnectDB()
	sizes := make([]entity.Size, 0)
	var minPages sql.NullInt32
	var maxPages sql.NullInt32
	rows, err := DB.Query("SELECT * FROM size")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var size entity.Size
		if err := rows.Scan(&size.ID, &size.Title, &minPages, &maxPages); err != nil {
			log.Fatal(err)
		}
		size.MinPages = minPages.Int32
		size.MaxPages = maxPages.Int32
		sizes = append(sizes, size)
	}
	sizeJson, err := json.Marshal(sizes)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(sizeJson)

	defer rows.Close()
	defer DB.Close()
}

func GetEras(w http.ResponseWriter, r *http.Request) {
	DB := internal.ConnectDB()
	eras := make([]entity.Era, 0)
	var minYear sql.NullInt32
	var maxYear sql.NullInt32
	rows, err := DB.Query("SELECT * FROM Era")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var era entity.Era
		if err := rows.Scan(&era.ID, &era.Title, &minYear, &maxYear); err != nil {
			log.Fatal(err)
		}
		era.MinYear = minYear.Int32
		era.MaxYear = maxYear.Int32
		eras = append(eras, era)
	}
	eraJson, err := json.Marshal(eras)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(eraJson)

	defer rows.Close()
	defer DB.Close()
}
