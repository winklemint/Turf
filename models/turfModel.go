package models

import (
	"time"

	"gorm.io/gorm"
)

type Time_Slot struct {
	gorm.Model
	Start_time string
	End_time   string
	Status     int
}

type Turf_Bookings struct {
	gorm.Model
	User_id                  int
	Slot_id                  int
	Date                     time.Time
	Is_booked                bool
	Package_slot_relation_id int
	Package_id               int
	Payment_ref_id           int
	Payment_screenshort      int
}

type Package_slot_relationship struct {
	gorm.Model
	Package_id int
	Slot_id    int
}

type Package struct {
	gorm.Model
	Name   string ` grom:"unique"`
	Price  float64
	Status bool
}
type Conform_Booking_Table struct {
	gorm.Model
	User_id          int
	Date             time.Time
	Booking_order_id string
	Slot_id          int
	Booking_table_id int
}
