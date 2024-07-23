package request

type BorrowRequest struct {
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}

type BorrowDatesUpdateRequest struct {
	BorrowID  string `json:"borrow_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
