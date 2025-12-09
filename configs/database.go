package configs

import (
	"golang_auth/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang_auth?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("Fail to connect database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database connected")
}