package main

import (
	"librarian/router"
	"log"
	"net/http"
)

func main() {
	router.Initialize()
	log.Fatal(http.ListenAndServe(":5001", nil))
}
