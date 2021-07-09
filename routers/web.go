package routers

import (
	"github.com/Caproner/DemoGame_Backend/handlers/gm"
	"github.com/Caproner/DemoGame_Backend/handlers/login"
	"github.com/Caproner/DemoGame_Backend/handlers/testecho"
	"github.com/gin-gonic/gin"
)

func InitRoute(router *gin.Engine) *gin.Engine {
	// login
	router = login.InitRouter(router)
	///api/test
	router = testecho.InitRouter(router)
	//gm
	router = gm.InitRouter(router)

	// player
	// router = player.InitRouter(router)

	return router
}
