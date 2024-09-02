package request

type BookRequest struct {
	Name        string          `json:"name"`
	Author      string          `json:"author"`
	Publication string          `json:"publication"`
	CategoryID    int64 `json:"categoryId"`
}