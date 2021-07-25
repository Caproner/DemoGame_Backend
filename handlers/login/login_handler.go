package login

import (
	"errors"
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
)

// 个人申请数据，后期变更为配置数据
var appid = "wx8534674533e0c0d0"
var secret = "3397873007ab85a4a9a8372a6faeb05e"

// InitRouter 初始化login路由，并建立websocket链接
func InitRouter(router *gin.Engine) *gin.Engine {
	router.GET("/login", loginHandler)
	return router
}

// 检测vx接口返回信息并向Svr和wsHub注册
func loginHandler(ctx *gin.Context) {
	code := ctx.Query(variable.CODE)
	openId, sessionKey, token, err := checkCode(code)
	if err != nil {
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: err.Error()})
	}else{
		msg := proto.S2CLogin{OpenID: openId,SessionKey: sessionKey,Token: token}
		responseresult.ResponseOk(ctx, proto.MSGS2CLogin, msg)
	}
}
// vx接口返回值,未做校验
func checkCode(code string) (openId, sessionKey, token string, err error) {
	if code == "gmcode"{ // 兼容测试数据，到时候补充环境设定
		openId = "o6d_Z5KqkxYXjQ04tPTLnK0L9CrI"
		sessionKey = "sessionkey_" + openId
		t:= player.PlayerLogin(openId,sessionKey)
		token = t
		err = nil
	}else{
		res, wxerr := weapp.Login(appid, secret, code)
		if wxerr != nil {
			err = wxerr
		}else if res.ErrCode != 0{
			log.Info(res.ErrMSG)
			err = errors.New("vxapi check error:" + res.ErrMSG)
		} else{
			openId = res.OpenID
			sessionKey = res.SessionKey
			t:= player.PlayerLogin(openId,sessionKey)
			token = t
			err = nil
		}
	}
	return
}
