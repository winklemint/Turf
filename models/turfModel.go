package models

import (
	"time"

	"gorm.io/gorm"
)

type Time_Slot struct {
	gorm.Model
	Start_time time.Time
	End_time   time.Time
	Status     int
}

type Turf_Bookings struct {
	gorm.Model
	User_id                  int
	Slot_id                  int
	Date                     time.Time
	Is_booked                bool
	package_slot_relation_id int
}

type Package_slot_relationship struct {
	gorm.Model
	Package_id int
	Slot_id    int
}

type Package struct {
	gorm.Model
	Name   string
	Price  float64
	Status bool
}
