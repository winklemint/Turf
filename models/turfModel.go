package models

import (
	"gorm.io/gorm"
)

type Time_Slot struct {
	gorm.Model
	Start_time     string
	End_time       string
	Day            string
	Unique_slot_id string
	Status         int
	Branch_id      uint
}

type Turf_Bookings struct {
	gorm.Model
	User_id                  uint
	Slot_id                  int
	Date                     string
	Is_booked                int
	Package_slot_relation_id int
	Package_id               uint
	Price                    float64
	Minimum_amount_to_pay    float64
	Order_id                 string
	Branch_id                int
}

type Package_slot_relationship struct {
	gorm.Model
	Package_id uint
	Slot_id    string
}

type Package struct {
	gorm.Model
	Name       string ` grom:"unique"`
	Price      float64
	Status     int
	Branch_id  int
	Avail_days string
}
type Confirm_Booking_Table struct {
	gorm.Model
	User_id                 uint
	Date                    string
	Booking_order_id        string
	Total_price             float64
	Total_min_amount_to_pay float64
	Paid_amount             float64
	Remaining_amount_to_pay float64
	Booking_status          int
	Branch_id               int
}

type Screenshot struct {
	gorm.Model
	Booking_order_id   string
	Amount             float64
	Payment_screenshot string
}

type Branch_info_management struct {
	gorm.Model
	Turf_name             string
	Branch_name           string
	Branch_address        string
	Branch_email          string
	Branch_contact_number string
	GST_no                string
	Status                int
}
