package database

import (
	"log"

	"github.com/JulesAD96/go-jwt-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to database")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database migration completed")
}
