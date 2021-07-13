package global

import (
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/services/playersvr"
	"github.com/Caproner/DemoGame_Backend/services/websocket/ws"
)

func init() {

}

func StartGlobalTask() {
	// 开启玩家进程管理
	variable.PlayerSvr = playersvr.New()
	go variable.PlayerSvr.(*playersvr.PlayerSvr).Loop()
	// 开启信息管理
	variable.WSHub = ws.New()
	go variable.WSHub.(*ws.WsHub).Loop()
}
