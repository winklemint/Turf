package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name           string
	Contact        string ` grom:"unique"`
	Email          string `gorm:"unique"`
	Password       string
	Role           int
	LastLogin      time.Time
	Turf_branch_id uint
}
type Content struct {
	gorm.Model
	Heading    string
	SubHeading string
	Button     string
	Status     int
}
type Carousel struct {
	gorm.Model
	Image  string
	Status string
}
