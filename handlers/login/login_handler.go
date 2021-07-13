package login

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Caproner/DemoGame_Backend/utils/log"

	r "github.com/Caproner/DemoGame_Backend/include/global/r/player"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/services/player"
	"github.com/Caproner/DemoGame_Backend/services/playerprocess"
	"github.com/Caproner/DemoGame_Backend/services/playersvr"
	"github.com/Caproner/DemoGame_Backend/services/websocket/ws"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	code := ctx.DefaultQuery("code", "no code login")
	openId, sessionKey, err := checkCode(code)
	if err != nil {
		responseresult.ResponseFalse(ctx, 200, "")
	}

	p := player.PlayerLogin(openId, sessionKey)

	register(ctx, p)
}

// 升级成websocket协议
func register(ctx *gin.Context, p *r.Player) {
	// 新建一个player进程
	ppc := playerprocess.New()
	ppc.Player = p
	go ppc.Loop()
	variable.PlayerSvr.(*playersvr.PlayerSvr).Register <- p
	lws := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := lws.Upgrade(ctx.Writer, ctx.Request, nil)
	log.Info("register had create conn")
	if err != nil {
		fmt.Println(err)
		return
	}
	if hub, ok := variable.WSHub.(*ws.WsHub); !ok {
		log.Error("not right hub type adn return")
		return
	} else {
		uuIDStr := strconv.FormatInt(p.UUID, 10)
		wsc := &ws.WsClient{
			Hub:     hub,
			UUID:    uuIDStr,
			Socket:  conn,
			Sender:  make(chan []byte),
			WsPSend: ppc.RecMsg,
		}
		log.Infof("wsc %s had create \n", wsc.UUID)
		hub.Register <- wsc
		go wsc.Read()
		go wsc.Send()
	}
}

// vx接口返回值,未做校验
func checkCode(code string) (openId string, sessionKey string, err error) {
	res, wxerr := weapp.Login(appid, secret, code)
	if wxerr != nil {
		err = wxerr
	}
	openId = res.OpenID
	sessionKey = res.SessionKey
	err = nil
	return
}
