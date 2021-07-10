package gm

/*
通过访问后台接口可以看到后台数据
*/

import (
	"github.com/Caproner/DemoGame_Backend/services/playersvr"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)

type GMHandlerRps struct {
	ShowSvr interface{}
}

func InitRouter(router *gin.Engine) *gin.Engine {
	gmrouter := router.Group("/gm").Use()
	{
		gmrouter.GET("/showSvr", showSvrHandler)
	}
	return router
}

func showSvrHandler(ctx *gin.Context) {
	rsp := GMHandlerRps{
		ShowSvr: playersvr.DefaultSvr(),
	}
	responseresult.ResponseOk(ctx, rsp)
}
