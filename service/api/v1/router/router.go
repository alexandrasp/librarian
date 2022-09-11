package router

import (
	"librarian/handlers"
	"net/http"
)

func Initialize() {
	http.HandleFunc("/api/v1/books", handlers.GetBooks)
	http.HandleFunc("/api/v1/authors", handlers.GetAuthors)
	http.HandleFunc("/api/v1/genres", handlers.GetGenres)
	http.HandleFunc("/api/v1/sizes", handlers.GetSizes)
	http.HandleFunc("/api/v1/eras", handlers.GetEras)
}
