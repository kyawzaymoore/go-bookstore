package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kzm/go-bookstore/api/config"
)

var db *gorm.DB

func init()  {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&UserSession{})
	db.AutoMigrate(&User{})
}