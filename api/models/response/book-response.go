package response

type BookResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Author      string          `json:"author"`
	Publication string          `json:"publication"`
	Category    CategoryResponse `json:"category"`
}