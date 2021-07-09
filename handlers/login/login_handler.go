package login

import (
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)

type LoginHandlerRps struct {
	UUID int64
}

func InitRouter(router *gin.Engine) *gin.Engine {
	router.GET("/login", loginHandler)
	return router
}

// 
func loginHandler(ctx *gin.Context) {
	code := ctx.Query("code")
	openId, sessionKey, err := checkCode(code)
	if err != nil {
		responseresult.ResponseFalse(ctx, 200, "")
	}

	uuid := player.PlayerLogin(openId, sessionKey)

	rsp := &LoginHandlerRps{UUID: uuid + 1000}
	responseresult.ResponseOk(ctx, rsp)
}

func checkCode(code string) (openId string, sessionKey string, err error) {
	openId = "vx10000"
	sessionKey = openId + "_sessionKey"
	err = nil
	return
}
