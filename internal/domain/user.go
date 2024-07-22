package domain

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Books    []*Book `json:"books"`
}
