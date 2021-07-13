package libsend

import (
	"encoding/json"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/services/websocket/ws"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
)

// Send 快速调用信息返回给前端
func Send(uuID int64, rsp interface{}){
	uuIDStr := tr.Int64ToString(uuID)
	oneWsc,_ := variable.WSHub.(*ws.WsHub).Clients[uuIDStr]
	log.Info(oneWsc,rsp)
	b,_ := json.Marshal(&variable.Message{Sender: uuIDStr, Rsp: rsp})
	oneWsc.Sender <- b
}
