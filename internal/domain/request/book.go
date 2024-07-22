package request

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
