package player

import (
	"github.com/Caproner/DemoGame_Backend/services/playersvr"
	"github.com/gin-gonic/gin"
)

func PlayerLogin(openId, sessionKey string) int64 {
	p := playersvr.LoadPlayer(openId)
	p.SessionKey = sessionKey
	playersvr.LoginPlayer(p)
	return p.UUID
}

//这里处理不同协议号的逻辑，并返回数据
func PlayerAcrion(ctx *gin.Context){

}

// 应该在校验那里获得了玩家openId，考虑数据唯一性，此处采用管道数据处理
func doSwitchProto(ctx *gin.Context){

}