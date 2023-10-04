package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique"`
	Password  string
	Is_active int
}
type Booking struct {
	gorm.Model
	Date      string
	Day       string
	Slot      int
	StartSlot string
	EndSlot   string
}
