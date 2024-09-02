package models

import (
	"time"
)

type User struct {
	ID           string    `gorm:"type:char(36);primaryKey""json:"id"`
	Username     string    `gorm:"type:varchar(500);not null""json:"username"`
	Password     string    `gorm:"type:varchar(500);not null""json:"password"`
	Name         string    `gorm:"type:varchar(500);not null""json:"name"`
	RegisterDate time.Time `gorm:"type:datetime;not null""json:"registerDate"`
}


