package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/books", app.handlers.Books.ListBooksHandler)
	mux.HandleFunc("POST /api/books", app.handlers.Books.AddBookHandler)

	return requestLogger(mux)
}
