package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Full_Name string
	Email     string `gorm:"unique"`
	Contact   string `gorm:"unique"`
	Password  string
	Is_active int
}
