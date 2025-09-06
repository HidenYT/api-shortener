package shortener_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var testDataRaw = []byte(`
	{
		"store": {
			"book": [
				{
					"title": "The Great Gatsby",
					"author": "F. Scott Fitzgerald",
					"price": 12.99,
					"inStock": true
				},
				{
					"title": "1984",
					"author": "George Orwell",
					"price": 10.50,
					"inStock": false
				},
				{
					"title": "To Kill a Mockingbird",
					"author": "Harper Lee",
					"price": 9.99,
					"inStock": true
				}
			],
			"location": "New York"
		}
	}
`)

var testRules = map[string]string{
	"allBookNames":     "$.store.book[*].title",
	"allBooksInStock":  "$.store.book[?(@.inStock == true)]",
	"allCheapBooks":    "$.store.book[?(@.price < 11)]",
	"secondBookAuthor": "$.store.book[1].author",
	"storeLocation":    "$.store.location",
}

var testData any

var expectedResponse = map[string]any{
	"allBookNames": []any{"The Great Gatsby", "1984", "To Kill a Mockingbird"},
	"allBooksInStock": []any{
		map[string]any{
			"title":   "The Great Gatsby",
			"author":  "F. Scott Fitzgerald",
			"price":   12.99,
			"inStock": true,
		},
		map[string]any{
			"title":   "To Kill a Mockingbird",
			"author":  "Harper Lee",
			"price":   9.99,
			"inStock": true,
		},
	},
	"allCheapBooks": []any{
		map[string]any{
			"title":   "1984",
			"author":  "George Orwell",
			"price":   10.50,
			"inStock": false,
		},
		map[string]any{
			"title":   "To Kill a Mockingbird",
			"author":  "Harper Lee",
			"price":   9.99,
			"inStock": true,
		},
	},
	"secondBookAuthor": []any{"George Orwell"},
	"storeLocation":    []any{"New York"},
}

var testHttpServer *httptest.Server

func TestMain(m *testing.M) {
	json.Unmarshal(testDataRaw, &testData)

	testHttpServer = httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "application/json")
				w.Header().Add("X-Some-Random-Header", "abcd123")
				w.Write(testDataRaw)
			},
		),
	)

	code := m.Run()
	testHttpServer.Close()
	os.Exit(code)
}
