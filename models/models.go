package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	id  uint  `gorm:"primaryKey"`
	name string 
	age string 
	phone_number   string 
	created_at   time.Time 
	modified_at   time.Time 
}
