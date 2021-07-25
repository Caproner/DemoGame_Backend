package player

import (
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/middlewares/playerauth"
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
	"github.com/gin-gonic/gin"
)

// InitRouter 简单的http接口测试
func InitRouter(router *gin.Engine) *gin.Engine {
	playerRouter := router.Group("/player").Use(playerauth.UserToken())
	{
		playerRouter.GET("/action", actionHandler)
		playerRouter.POST("/action", actionHandler)
	}
	return router
}

// 提取协议号
func actionHandler(ctx *gin.Context){
	Proto := ctx.Query(variable.PROTOID)
	player.MsgHandler(ctx, tr.StringToInt(Proto))
}
