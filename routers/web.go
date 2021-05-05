package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Caproner/DemoGame_Backend/handlers/testecho"
)

func InitRoute(router *gin.Engine) *gin.Engine {
	testRouter := router.Group("/api/test").Use()
	{
		testRouter.GET("/echo", testecho.TestHandler)
	}

	return router
}
