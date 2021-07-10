package login

import (
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
)

// 个人申请数据，后期变更为配置数据
var appid  = "wx565c09529f19cbaa"
var secret = "c841730a2f3f3fa58e00def83ecf2999"

type LoginHandlerRps struct {
	UUID int64
}

func InitRouter(router *gin.Engine) *gin.Engine {
	router.GET("/login", loginHandler)
	return router
}

// 
func loginHandler(ctx *gin.Context) {
	code := ctx.DefaultQuery("code", "no code login")
	openId, sessionKey, err := checkCode(code)
	if err != nil {
		responseresult.ResponseFalse(ctx, 200, "")
	}

	uuid := player.PlayerLogin(openId, sessionKey)

	rsp := &LoginHandlerRps{UUID: uuid + 1000}
	responseresult.ResponseOk(ctx, rsp)
}

func checkCode(code string) (openId string, sessionKey string, err error) {
	res, wxerr := weapp.Login(appid, secret, code)
	if wxerr != nil{
		err = wxerr
	}
	openId = res.OpenID
	sessionKey = res.SessionKey
	err = nil
	return
}
