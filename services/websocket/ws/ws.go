package ws

import (
	"encoding/json"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/utils/log"
)

type WsHub struct {
	Register   chan *WsClient
	UnRegister chan *WsClient
	ClientEvt chan variable.Message
	Broadcast  chan []byte
	Clients    map[string]*WsClient
}


func (hub *WsHub) Loop() {
	log.Info("ws enter loop")
	for {
		select {
		case reg := <-hub.Register:
			connID := reg.UUID
			log.Info("wshub", connID)
			hub.Clients[connID] = reg
			jsonMsg, _ := json.Marshal(&variable.Message{Content: "connect ok"})
			reg.Sender <- jsonMsg
		case reg := <-hub.UnRegister:
			if _, ok := hub.Clients[reg.UUID]; ok {
				jsonMessage, _ := json.Marshal(&variable.Message{Content: "A socket has disconnected"})
				reg.Sender <- jsonMessage
				close(reg.Sender)
				delete(hub.Clients, reg.UUID)
			}
		case message := <-hub.Broadcast:
			// 到时候这里会的处理会去到play那边处理
			MS := &variable.Message{}
			_ = json.Unmarshal(message, MS)
			for _, oneClient := range hub.Clients {
				MS.Recver = oneClient.UUID
				mss,_ := json.Marshal(MS)
				select {
				case oneClient.Sender <- mss:
				default:
					close(oneClient.Sender)
					delete(hub.Clients, oneClient.UUID)
				}
			}
		case evtmessage := <- hub.ClientEvt:

			if oneClient,ok := hub.Clients[evtmessage.Sender];ok{
				b,_ := json.Marshal(evtmessage)
				oneClient.WsPSend <- b
			}else{
				log.Error("not find")
			}

		}
	}
}

func New() *WsHub {
	return &WsHub{
		Register:   make(chan *WsClient),
		UnRegister: make(chan *WsClient),
		Broadcast:  make(chan []byte),
		ClientEvt:  make(chan variable.Message),
		Clients:    make(map[string]*WsClient),
	}
}
