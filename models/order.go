package models

import "time"

type Order struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	ProductRefer int `josn:"product_id"`
	Product Product `gorm:"foreignKey:ProductRefer"`
	UserRefer int `josn:"user_id"`
	User Product `gorm:"foreignKey:UserRefer"`
}