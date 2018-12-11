package main

import (
	"github.com/getupandgo/snooper-wooper/controllers"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/getupandgo/snooper-wooper/connectors"
	"github.com/getupandgo/snooper-wooper/dao"
)

func main() {
	db, err := connectors.InitDB()
	if err != nil {
		log.Fatalf("Failed to init database with error %+v", err)
	}

	r := controllers.InitRouter(dao.NewTokensDao(db))

	//fixme: handle error here (e.g. EADDRINUSE)
	_ = r.Run(":8000")
}
