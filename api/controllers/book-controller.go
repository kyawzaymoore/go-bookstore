package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	errorcode "github.com/kzm/go-bookstore/api/helper"
	"github.com/kzm/go-bookstore/api/models/request"
	"github.com/kzm/go-bookstore/api/models/response"
	"github.com/kzm/go-bookstore/api/repository"
	"github.com/kzm/go-bookstore/api/utils"
)

//var NewBook models.Book
type BookController struct {
	BookRepo repository.BookRepository
}

func NewBookController(bookRepo repository.BookRepository) *BookController {
	return &BookController{BookRepo: bookRepo}
}

// Book godoc
// @Summary Get Book List
// @Description Get Book List
// @Tags Book
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseModel[[]response.BookResponse]
// @Security BearerAuth
// @Router /book [get]
func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API : GetBook")
	newBooks := c.BookRepo.GetAllBooks()
	data := map[string]interface{}{
		"list": newBooks,
	}

	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, data)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Book godoc
// @Summary Get Book By ID
// @Description Get Book By ID
// @Tags Book
// @Param bookId path string true "Book ID"
// @Produce  json
// @Success 200 {object} response.ResponseModel[response.BookResponse]
// @Security BearerAuth
// @Router /book/{bookId} [get]
func (c *BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails := c.BookRepo.GetBookById(ID)
	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, bookDetails)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Book godoc
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept  json
// @Produce  json
// @Param book body request.BookRequest true "Book model"
// @Success 200 {object} response.ResponseModel[response.BookResponse]
// @Security BearerAuth
// @Router /book [post]
func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &request.BookRequest{}
	utils.ParseBody(r, CreateBook)
	b := c.BookRepo.CreateBook(*CreateBook)
	//res, _ := json.Marshal(b)

	res, _ := json.Marshal((response.GetResponseModel("E0", b)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Book godoc
// @Summary Delete Book By ID
// @Description Delete Book By ID
// @Tags Book
// @Param bookId path string true "Book ID"
// @Produce  json
// @Security BearerAuth
// @Router /book/{bookId} [delete]
func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := c.BookRepo.DeleteBook(ID)
	//res, _ := json.Marshal(book)
	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, book)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Book godoc
// @Summary Update book
// @Description Update book
// @Tags Book
// @Param bookId path string true "Book ID"
// @Param book body request.BookRequest true "Book data"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseModel[response.BookResponse]
// @Security BearerAuth
// @Router /book/{bookId} [put]
func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &request.BookRequest{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails := c.BookRepo.UpdateBook(ID, updateBook)
	//res, _ := json.Marshal(booksDetails)
	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, booksDetails)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
