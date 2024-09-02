package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	errorcode "github.com/kzm/go-bookstore/api/helper"
	"github.com/kzm/go-bookstore/api/models"
	"github.com/kzm/go-bookstore/api/models/response"
	"github.com/kzm/go-bookstore/api/repository"
	"github.com/kzm/go-bookstore/api/utils"
)

type CategoryController struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryController(categoryRepo repository.CategoryRepository) *CategoryController {
	return &CategoryController{CategoryRepo: categoryRepo}
}

// Category godoc
// @Summary Get Category List
// @Description Get Category List
// @Tags Category
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseModel[[]models.SwaggerCategory]
// @Security BearerAuth
// @Router /category [get]
func (c *CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API : GetCategory")
	newCategorys := c.CategoryRepo.GetAllCategorys()
	data := map[string]interface{}{
		"list": newCategorys,
	}

	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, data)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Category godoc
// @Summary Get Category By ID
// @Description Get Category By ID
// @Tags Category
// @Param categoryId path string true "Category ID"
// @Produce  json
// @Success 200 {object} response.ResponseModel[models.SwaggerCategory]
// @Security BearerAuth
// @Router /category/{categoryId} [get]
func (c *CategoryController) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	ID, err := strconv.ParseInt(categoryId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	categoryDetails := c.CategoryRepo.GetCategoryById(ID)
	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, categoryDetails)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Category godoc
// @Summary Create category
// @Description Create category
// @Tags Category
// @Accept  json
// @Produce  json
// @Param category body models.SwaggerCategory true "Category model"
// @Success 200 {object} response.ResponseModel[models.SwaggerCategory]
// @Security BearerAuth
// @Router /category [post]
func (c *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	CreateCategory := &models.Category{}
	utils.ParseBody(r, CreateCategory)
	b := c.CategoryRepo.CreateCategory(*CreateCategory)
	//res, _ := json.Marshal(b)

	res, _ := json.Marshal((response.GetResponseModel("E0", b)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Category godoc
// @Summary Delete Category
// @Description Delete Category
// @Tags Category
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ResponseModel[[]models.SwaggerCategory]
// @Security BearerAuth
// @Router /category/{categoryId} [delete]
func (c *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	ID, err := strconv.ParseInt(categoryId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	Category := c.CategoryRepo.DeleteCategory(ID)
	//res, _ := json.Marshal(Category)
	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, Category)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Category godoc
// @Summary Update category
// @Description Update category
// @Tags Category
// @Param categoryId path string true "Category ID"
// @Param category body models.SwaggerCategory true "Category data"
// @Success 200 {object} response.ResponseModel[models.SwaggerCategory]
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Router /category/{categoryId} [put]
func (c *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var updateCategory = &models.Category{}
	utils.ParseBody(r, updateCategory)
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]
	ID, err := strconv.ParseInt(categoryId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	categorysDetails := c.CategoryRepo.UpdateCategory(ID, updateCategory)
	//res, _ := json.Marshal(categorysDetails)
	res, _ := json.Marshal((response.GetResponseModel(errorcode.E0, categorysDetails)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
