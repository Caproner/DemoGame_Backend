package player

import (
	"github.com/Caproner/DemoGame_Backend/include/localvar/timevar"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/services/ptime"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
	"github.com/gin-gonic/gin"
	"time"
)

// InitRouter 简单的http接口测试
func InitRouter(router *gin.Engine) *gin.Engine {
	testRouter := router.Group("/player").Use()
	{
		testRouter.GET("/action", actionHandler)
		testRouter.GET("/sync", syncHandler)
	}
	return router
}

func actionHandler(ctx *gin.Context){
	Proto := ctx.Query(variable.PROTOID)
	player.MsgHandler(ctx, tr.StringToInt(Proto))
}

func syncHandler(ctx *gin.Context){
	openID := ctx.Query(variable.OPENID)
	if p, err := dbapi.FindPlayer(openID);err == nil{
		r := ptime.SetTime(timevar.LastActionTimeType,time.Now().Unix(), p)
		responseresult.ResponseOk(ctx, 0, r)
	}else{
		responseresult.ResponseFalse(ctx, 200, 0,"false")
	}
}
