package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kzm/go-bookstore/api/config"
	"github.com/kzm/go-bookstore/api/models"
)

type CategoryRepository interface {
	GetAllCategorys() []models.Category
	GetCategoryById(id int64) *models.Category
	CreateCategory(data models.Category) *models.Category
	DeleteCategory(id int64) models.Category
	UpdateCategory(id int64, updateCategory *models.Category) *models.Category
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{db: config.GetDB()}
}

func (r *categoryRepository) GetAllCategorys() []models.Category {
	var Categorys []models.Category
	r.db.Find(&Categorys)
	return Categorys
}

func (r *categoryRepository) GetCategoryById(id int64) *models.Category {
	var getCategory models.Category
	r.db.Where("ID=?", id).Find(&getCategory)
	return &getCategory
}

func (r *categoryRepository) CreateCategory(data models.Category) *models.Category {
	r.db.NewRecord(data)
	r.db.Create(&data)
	return &data
}

func (r *categoryRepository) DeleteCategory(id int64) models.Category {
	var Category models.Category
	r.db.Where("ID=?", id).Delete(Category)
	return Category
}

func (r *categoryRepository) UpdateCategory(id int64, updateCategory *models.Category) *models.Category {
	CategorysDetails := r.GetCategoryById(id)
	if updateCategory.Name != "" {
		CategorysDetails.Name = updateCategory.Name
	}
	r.db.Save(&CategorysDetails)
	return CategorysDetails
}
