package connectors

import (
	"github.com/jinzhu/gorm"

	"github.com/getupandgo/snooper-wooper/models"
	"github.com/getupandgo/snooper-wooper/utils/gorm_helpers"
)

func InitDB() (*gorm.DB, error) {
	connectionString, err := gorm_helpers.BuildConnectionString(gorm_helpers.Postgres, map[string]string{"host": "postgres-snoop"})
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(gorm_helpers.Postgres, connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Token{})
	return db, nil
}
