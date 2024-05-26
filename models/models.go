package models

import "github.com/google/uuid"

// "time"

// "gorm.io/gorm"
// "github.com/google/uuid"
// "github.com/jmoiron/sqlx"

type User struct {
	// gorm.Model
	ID  uuid.UUID  `db:"id"`
	// `gorm:"primaryKey"`
	Name string `db:"name"`
	Age string `db:"age"`
	Phone_number   string `db:"phone_number"`
	Email string `db:"email"`
	Password string `db:"password"`

}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
   }
   
type LoginResponse struct {
	Token string `json:"token"`
   }
