package models

import "time"

type Product struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name string `josn:"name"`
	SerialNumber string `josn:"serial_number"`
}