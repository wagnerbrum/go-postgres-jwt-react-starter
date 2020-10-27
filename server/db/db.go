package db

import (
	"fmt"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/config"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB instance
var DB *gorm.DB

//Connect to DB
func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", config.DbUser, config.DbPassword, config.DbName, config.HostName)
	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	executeMigrations()
}

func executeMigrations() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.CustomerLocation{})
	DB.AutoMigrate(&models.PaymentMethod{})
	DB.AutoMigrate(&models.PaymentBillingAddress{})

	// DB.Create(&models.Customer{CustomerID: "aab-123-123", Name: "teste 1", CustomerLocation: models.CustomerLocation{Country: "Brazil"}})
}
