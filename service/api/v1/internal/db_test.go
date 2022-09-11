package internal_test

import (
	"librarian/internal"
	"net/url"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	queryValues := url.Values{}
	queryValues.Add("limit", "10")

	bQuery := internal.QueryBuilder(queryValues)
	if bQuery != "SELECT * FROM book INNER JOIN genre ON book.genre_id = genre.id INNER JOIN author ON book.author_id = author.id LIMIT 10" {
		t.Errorf("Query mismatch")
	}

}
