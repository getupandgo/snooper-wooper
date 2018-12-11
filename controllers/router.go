package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/getupandgo/snooper-wooper/controllers/token"
	"github.com/getupandgo/snooper-wooper/dao"
)

func InitRouter(d dao.TokensDao) *gin.Engine {
	tokensCtrl := token.New(d)

	r := gin.New()
	tokensRouter := r.Group("/tokens")
	tokensRouter.GET("", tokensCtrl.GetTopTokens)
	tokensRouter.POST("", tokensCtrl.UpsertToken)

	return r
}
