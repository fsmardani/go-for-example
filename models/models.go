package models

import (
	// "time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID  uint  `gorm:"primaryKey"`
	Name string `json:name`
	Age string `json:age`
	Phone_number   string `json:phone_number`
	Email string `json:"email"`
	Password string `json:"password"`

}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
   }
   
type LoginResponse struct {
	Token string `json:"token"`
   }
