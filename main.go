package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/getupandgo/snooper-wooper/connectors"
	"github.com/getupandgo/snooper-wooper/controllers"
	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/utils/config"
)

func main() {
	conf, err := config.InitConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	db, err := connectors.InitDB(conf)
	if err != nil {
		log.Fatalf("Failed to init database with error %+v", err)
	}

	r := controllers.InitRouter(dao.NewTokensDao(db))

	if !conf.GetBool("http_debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	httpPort := conf.GetInt("http_port")
	fmt.Printf("starting API server on port %d", httpPort)

	//fixme: handle error here (e.g. EADDRINUSE)
	_ = r.Run(fmt.Sprintf(":%d", httpPort))
}
