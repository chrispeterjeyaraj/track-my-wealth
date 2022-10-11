package db

import (
	"log"

	"github.com/chrispeterjeyaraj/track-my-wealth/server/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://tmwapp:tmwv10@trackmywealthdb:5432/trackmywealth"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.UserPref{}, &models.Expense{}, &models.Income{}, &models.Saving{})

	return db
}
