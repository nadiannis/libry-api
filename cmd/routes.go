package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/books", app.handlers.Books.GetAll)
	mux.HandleFunc("POST /api/books", app.handlers.Books.Add)

	mux.HandleFunc("GET /api/users", app.handlers.Users.GetAll)
	mux.HandleFunc("POST /api/users", app.handlers.Users.Add)

	mux.HandleFunc("GET /api/borrowed-books", app.handlers.Borrows.GetAll)
	mux.HandleFunc("POST /api/borrowed-books", app.handlers.Borrows.Borrow)
	mux.HandleFunc("PATCH /api/borrowed-books", app.handlers.Borrows.Return)

	return requestLogger(mux)
}
