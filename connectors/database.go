package connectors

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"github.com/getupandgo/snooper-wooper/models"
	"github.com/getupandgo/snooper-wooper/utils/gorm_helpers"
)

func InitDB(conf *viper.Viper) (*gorm.DB, error) {
	options := map[string]string{
		"host": conf.GetString("db_host"),
		"port": conf.GetString("db_port"),
	}
	connectionString, err := gorm_helpers.BuildConnectionString(gorm_helpers.Postgres, options)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(gorm_helpers.Postgres, connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(conf.GetBool("db_debug"))
	db.AutoMigrate(&models.Token{})
	return db, nil
}
