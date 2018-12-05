package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/models"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to init database with error %+v", err)
	}
	router := InitRouter(&ctx{tokens: dao.NewTokensDao(db)})
	// fixme: handle error here (e.g. EADDRINUSE)
	_ = http.ListenAndServe(":8000", router)
}

func initDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=postgres dbname=postgres password=postgres sslmode=disable", "localhost", "5432")
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.AutoMigrate(&models.Token{})
	return db, nil
}
