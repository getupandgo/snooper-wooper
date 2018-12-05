package controllers

import "github.com/gin-gonic/gin"

func InitRouter(c *ctx) *gin.Engine {
	router := gin.Default()

	//todo: implement partial routers?
	router.GET("/tokens", c.GetTokens)
	router.POST("/tokens", c.SaveTokens)

	return router
}
