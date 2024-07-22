package main

import (
	"github.com/nadiannis/libry-api/internal/domain/request"
	"github.com/nadiannis/libry-api/internal/usecase"
)

var bookInputs = []*request.BookRequest{
	{
		Title:  "Grit",
		Author: "Angela Duckworth",
	},
	{
		Title:  "Mindset",
		Author: "Carol Dweck",
	},
	{
		Title:  "Ikigai",
		Author: "Hector Garcia",
	},
	{
		Title:  "Drive",
		Author: "Daniel H. Pink",
	},
	{
		Title:  "Range",
		Author: "David Epstein",
	},
	{
		Title:  "Sapiens",
		Author: "Yuval N. Harari",
	},
	{
		Title:  "Essentialism",
		Author: "Greg McKeown",
	},
	{
		Title:  "Deep Work",
		Author: "Cal Newport",
	},
	{
		Title:  "The Lean Startup",
		Author: "Eric Ries",
	},
	{
		Title:  "Atomic Habits",
		Author: "James Clear",
	},
}

func prepopulateBooks(usecase usecase.IBookUsecase) {
	for _, bookInput := range bookInputs {
		book := &request.BookRequest{
			Title:  bookInput.Title,
			Author: bookInput.Author,
		}
		usecase.Add(book)
	}
}
