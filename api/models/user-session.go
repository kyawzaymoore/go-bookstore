package models

import (
	"time"
)

type UserSession struct {
	ID          string    `gorm:"type:char(36);primaryKey""json:"id"`
	UserID      string    `gorm:"type:char(36);not null""json:"userId"`
	Token       string    `gorm:"type:varchar(500);not null""json:"token"`
	CreatedDate time.Time `gorm:"type:datetime;not null""json:"registerDate"`
}
