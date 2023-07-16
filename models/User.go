package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	
	name     string `gorm:"not null" json:"name"`
	lastname string `gorm:"not null" json:"lastname"`
	email    string `gorm:"not null;unique_index" json:"email"`
	password string `gorm:"not null" json:"password"`
}