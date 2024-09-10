package configs

import (
	"fmt"
	"os"

	"example.com/gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DATABASE_DNS")
	db, err := (gorm.Open(postgres.Open(dsn), &gorm.Config{}))

	if err != nil {
		fmt.Println("can't connect server")
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("connect database succes")

	// Migration
	// db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	DB = db
}
