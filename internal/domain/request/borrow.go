package request

type BorrowRequest struct {
	UserID string `json:"user_id"`
	BookID string `json:"book_id"`
}
