package player

import (
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)


func InitRouter(router *gin.Engine) *gin.Engine {
	playerRouter := router.Group("/player").Use()
	{
		playerRouter.GET("/action", actionHandler)
	}
	return router
}

func actionHandler(ctx *gin.Context){
	protoId := ctx.DefaultQuery("protoId", "none")
	if protoId == "none"{
		responseresult.ResponseFalse(ctx, 201, "no request protoId")
	}else{
		copyctx := ctx.Copy()
		player.PlayerAcrion(copyctx) // 采用异步返回
		responseresult.ResponseOk(ctx, "ok")
	}
}