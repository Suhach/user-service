package database

import (
	"github.com/Suhach/User-service/internal/user"
	"gorm.io/driver/postgres"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=12345678 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}
}
