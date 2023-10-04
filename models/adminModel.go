package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name      string
	Contact   string ` grom:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Role      int
	LastLogin time.Time
}
type Slot struct {
	gorm.Model
	StartSlot string ` grom:"unique"`
	EndSlot   string
}
