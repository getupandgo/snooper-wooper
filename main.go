package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/zerolog/log"

	"github.com/getupandgo/snooper-wooper/connectors"
	"github.com/getupandgo/snooper-wooper/controllers"
	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/utils/config"
)

func main() {
	conf, err := config.InitConfiguration()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to get configuration file")
	}

	db, err := connectors.InitDB(conf)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to init database")
	}

	if !conf.GetBool("http_debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	r := controllers.InitRouter(dao.NewTokensDao(db))

	httpPort := conf.GetInt("http_port")

	if err = r.Run(fmt.Sprintf(":%d", httpPort)); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start server")
	}

	log.Info().Msgf("starting API server on port %d", httpPort)
}
