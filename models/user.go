package models

import "time"

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `josn:"first_name"`
	LastName string `josn:"last_name"`
 	
}