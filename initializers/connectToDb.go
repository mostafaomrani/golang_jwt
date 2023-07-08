package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConncetToDb() {
	var err error
	dsn := os.Getenv("DB_connection")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("ConncetToDb", DB)
		panic("Failed to open database")
	}

}
