package utils

import (
	"github.com/kzm/go-bookstore/api/models"
	"github.com/kzm/go-bookstore/api/models/response"
)

// MapToResponse takes a slice of models and a mapper function, and returns a slice of response structs
func MapToResponseList[T any, U any](models []T, mapper func(T) U) []U {
	var responses []U
	for _, model := range models {
		responses = append(responses, mapper(model))
	}
	return responses
}

func MapToSingleResponse[T any, U any](model T, mapper func(T) U) U {
	return mapper(model)
}

func MapBookToResponse(book models.Book) response.BookResponse {
	return response.BookResponse{
		ID:          book.ID,
		Name:        book.Name,
		Author:      book.Author,
		Publication: book.Publication,
		Category: response.CategoryResponse{
			ID:   book.Category.ID,
			Name: book.Category.Name,
		},
	}
}