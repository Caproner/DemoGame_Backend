package gm

/*
通过访问后台接口可以看到后台数据
*/

import (
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)

type GMHandlerRps struct {
	ShowSvr interface{}
}

// 暂时有bug未能实现
// InitRouter初始化gm路由
func InitRouter(router *gin.Engine) *gin.Engine {
	playerRouter := router.Group("/gm").Use()
	{
		playerRouter.GET("/svr", svrHandler)
	}
	return router
}

func svrHandler(ctx *gin.Context) {
	svr := variable.PlayerSvr
	responseresult.ResponseOk(ctx, svr)

}
