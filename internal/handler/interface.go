package handler

import "net/http"

type IUserHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}

type IBookHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
}

type IBorrowHandler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Borrow(w http.ResponseWriter, r *http.Request)
	Return(w http.ResponseWriter, r *http.Request)
	UpdateDates(w http.ResponseWriter, r *http.Request)
}
