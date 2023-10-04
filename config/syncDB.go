package config

import "turf/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
