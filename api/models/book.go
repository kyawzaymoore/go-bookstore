package models

import (
	"github.com/jinzhu/gorm"
)


type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
	CategoryID int64 `gorm:""json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
}