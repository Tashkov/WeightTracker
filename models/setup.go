package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const connStr string = "host=localhost user=postgres password=//EnterPassword// dbname=weight_loss-db port=5432 sslmode=disable"

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&WeightLog{})
	database.Debug().Migrator().CreateConstraint(&User{}, "WeightLogs")
	database.Debug().Migrator().CreateConstraint(&User{}, "fk_users_weight_log")

	DB = database
}
