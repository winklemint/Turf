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
	User_id                  uint
	Slot_id                  int
	Date                     time.Time
	Is_booked                bool
	Package_slot_relation_id int
	Package_id               int
	Price                    float64
	Minimum_amount_to_pay    float64
	Order_id                 string
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
type Confirm_Booking_Table struct {
	gorm.Model
	User_id                 int
	Date                    string
	Booking_order_id        string
	Total_price             float64
	Total_min_amount_to_pay float64
	Paid_amount             float64
	Remaining_amount_to_pay float64
	Booking_table_id        int
	Booking_status          int
}

type Screenshot struct {
	gorm.Model
	Booking_order_id   string
	Amount             float64
	Payment_screenshot string
}
