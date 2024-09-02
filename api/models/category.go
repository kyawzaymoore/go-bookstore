package models

import (
	"time"

	"github.com/jinzhu/gorm"
)


type Category struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Books []Book `gorm:"foreignKey:CategoryID""json:"books"`
}

type SwaggerCategory struct {
    ID        uint       `json:"id"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at,omitempty"`
    Name      string     `json:"name"`
}
