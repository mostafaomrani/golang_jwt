package initializers

import "jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
