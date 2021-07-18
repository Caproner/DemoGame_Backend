package login

import (
	"errors"
	"fmt"
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
)

// 个人申请数据，后期变更为配置数据
var appid = "wx565c09529f19cbaa"
var secret = "c841730a2f3f3fa58e00def83ecf2999"

// InitRouter 初始化login路由，并建立websocket链接
func InitRouter(router *gin.Engine) *gin.Engine {
	router.GET("/login", loginHandler)
	return router
}

// 检测vx接口返回信息并向Svr和wsHub注册
func loginHandler(ctx *gin.Context) {
	code := ctx.Query(variable.CODE)
	openId, sessionKey, err := checkCode(code)
	if err != nil {
		responseresult.ResponseFalse(ctx, 200, 0, err.Error())
	}else{
		msg := proto.S2CLogin{OpenID: openId,SessionKey: sessionKey,Token: "no token"}
		responseresult.ResponseOk(ctx, proto.MSGS2CLogin, msg)
	}
}
// vx接口返回值,未做校验
func checkCode(code string) (openId string, sessionKey string, err error) {
	res, wxerr := weapp.Login(appid, secret, code)
	if wxerr != nil {
		err = wxerr
	}else if res.ErrCode != 0{
		log.Info(res.ErrMSG)
		err = errors.New("vxapi check error:" + res.ErrMSG)
	} else{
		fmt.Println(res.ErrCode)
		openId = res.OpenID
		sessionKey = res.SessionKey
		if player.PlayerLogin(openId,sessionKey) {
			err = nil
		}else{
			err =  errors.New("load player error")
		}
	}
	return
}
