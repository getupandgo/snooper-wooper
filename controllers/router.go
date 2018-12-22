package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/getupandgo/snooper-wooper/controllers/token"
	"github.com/getupandgo/snooper-wooper/dao"
	"github.com/getupandgo/snooper-wooper/utils/gin_middlewares"
)

func InitRouter(d dao.TokensDao) *gin.Engine {
	tokensCtrl := token.New(d)

	r := gin.New()
	r.Use(gin_middlewares.AppErrorReporter())
	tokensRouter := r.Group("/tokens")
	tokensRouter.GET("", tokensCtrl.GetTopTokens)
	tokensRouter.POST("", tokensCtrl.UpsertToken)

	return r
}
