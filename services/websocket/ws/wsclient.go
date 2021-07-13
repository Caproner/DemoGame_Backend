package ws

import (
	"encoding/json"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/gorilla/websocket"
)

type WsClient struct {
	Hub    *WsHub
	UUID   string // 方便解析，所以改成string类型了
	Socket *websocket.Conn
	Sender chan []byte
	WsPSend chan []byte
}

// Read 接收前端的信息
func (wsc *WsClient) Read() {
	defer func() {
		wsc.Hub.UnRegister <- wsc
		_ = wsc.Socket.Close()
	}()
	for {
		wsc.Socket.PongHandler()
		_, message, err := wsc.Socket.ReadMessage()
		if err != nil {
			wsc.Hub.UnRegister <- wsc
			_ = wsc.Socket.Close()
			break
		}

		wsc.Hub.ClientEvt <- wsc.addPArm(message)
	}
}

func (wsc *WsClient)addPArm(message []byte)variable.Message{
	MS := &variable.Message{}
	_ = json.Unmarshal(message, MS)
	MS.Sender = wsc.UUID
	return *MS
 }

// Send 发送信息给前端
func (wsc *WsClient) Send() {
	defer func() {
		_ = wsc.Socket.Close()
	}()
	for {
		select {
		case msg, ok := <-wsc.Sender:
			if !ok {
				_ = wsc.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = wsc.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
