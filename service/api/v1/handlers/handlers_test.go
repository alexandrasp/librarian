package handlers_test

import (
	"librarian/handlers"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooksNotFoundAuthor(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "api/v1/books?authors=2&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := ExecuteRequest(request, handlers.GetBooks)
	CheckResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		log.Fatalf("Expected Empty Response")
	}
}

func TestGetBooks(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "api/v1/books?authors=1&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := ExecuteRequest(request, handlers.GetBooks)
	CheckResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		log.Fatalf("No Registers Found in DB")
	}
}

func TestGetBooksWrongParameter(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "api/v1/books?action=3", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := ExecuteRequest(request, handlers.GetBooks)
	CheckResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestGetAuthors(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "api/v1/authors", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := ExecuteRequest(request, handlers.GetAuthors)
	CheckResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		log.Fatalf("No Registers Found in DB")
	}
}

func TestGetGenres(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "api/v1/genres", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := ExecuteRequest(request, handlers.GetGenres)
	CheckResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		log.Fatalf("No Registers Found in DB")
	}
}

func TestGetSize(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "api/v1/size", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := ExecuteRequest(request, handlers.GetSizes)
	CheckResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		log.Fatalf("No Registers Found in DB")
	}
}

func TestGetEras(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "api/v1/eras", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := ExecuteRequest(request, handlers.GetEras)
	CheckResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		log.Fatalf("No Registers Found in DB")
	}
}

func ExecuteRequest(req *http.Request, h http.HandlerFunc) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h)
	handler.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
