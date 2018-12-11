package connectors

import (
	"fmt"
	"github.com/getupandgo/snooper-wooper/models"
	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=postgres dbname=postgres password=postgres sslmode=disable", "localhost", "5432")
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Token{})
	return db, nil
}
