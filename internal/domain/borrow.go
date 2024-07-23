package domain

import "time"

type Status string

var (
	StatusBorrowed Status = "borrowed"
	StatusReturned Status = "returned"
	StatusOverdue  Status = "overdue"
)

type Borrow struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	BookID    string    `json:"book_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    Status    `json:"status"`
}
