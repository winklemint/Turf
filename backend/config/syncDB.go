package config

import "turf/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Time_Slot{})
	DB.AutoMigrate(&models.Turf_Bookings{})
	DB.AutoMigrate(&models.Package{})
	DB.AutoMigrate(&models.Package_slot_relationship{})
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.Confirm_Booking_Table{})
	DB.AutoMigrate(models.Screenshot{})
	DB.AutoMigrate(models.Branch_info_management{})
	DB.AutoMigrate(models.Testi_Monial{})
	DB.AutoMigrate(&models.Content{})
	DB.AutoMigrate(&models.Carousel{})
	DB.AutoMigrate(&models.Navbar{})
	DB.AutoMigrate(&models.Heading{})
	DB.AutoMigrate(&models.Icon{})
	DB.AutoMigrate(&models.TermsAndConditions{})
}
