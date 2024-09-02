package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kzm/go-bookstore/api/config"
	"github.com/kzm/go-bookstore/api/models"
	"github.com/kzm/go-bookstore/api/models/request"
	"github.com/kzm/go-bookstore/api/models/response"
	"github.com/kzm/go-bookstore/api/utils"
)

type BookRepository interface {
	GetAllBooks() []response.BookResponse 
	GetBookById(id int64) response.BookResponse
	CreateBook(data request.BookRequest) response.BookResponse
	DeleteBook(id int64) models.Book
	UpdateBook(id int64, updateBook *request.BookRequest) response.BookResponse
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository() BookRepository {
	return &bookRepository{db: config.GetDB()}
}

func (r *bookRepository) GetAllBooks() []response.BookResponse {
	var Books []models.Book
	r.db.Preload("Category").Find(&Books)
	return utils.MapToResponseList(Books, utils.MapBookToResponse)
}

func (r *bookRepository) GetBookById(id int64) response.BookResponse {
	var getBook models.Book
	r.db.Preload("Category").Where("ID=?", id).Find(&getBook)
	return utils.MapToSingleResponse(getBook, utils.MapBookToResponse)
}

func (r *bookRepository) CreateBook(data request.BookRequest) response.BookResponse{

	book := models.Book{
		Name: data.Name,
		Author: data.Author,
		Publication: data.Publication,
		CategoryID: data.CategoryID,
	}

	r.db.Create(&book)
	return r.GetBookById(int64(book.ID))
}

func (r *bookRepository) DeleteBook(id int64) models.Book {
	var book models.Book
	r.db.Where("ID=?", id).Delete(book)
	return book
}

func (r *bookRepository) UpdateBook(id int64, updateBook *request.BookRequest) response.BookResponse {
	var booksDetails models.Book
	r.db.Where("ID=?", id).Find(&booksDetails)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	r.db.Save(&booksDetails)
	return r.GetBookById(int64(booksDetails.ID))
}
