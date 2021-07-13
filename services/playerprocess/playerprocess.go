package playerprocess

import (
	"encoding/json"

	r "github.com/Caproner/DemoGame_Backend/include/global/r/player"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	logicp "github.com/Caproner/DemoGame_Backend/logic/player"
	"github.com/Caproner/DemoGame_Backend/utils/log"
)

type PPC struct {
	Player *r.Player
	RecMsg chan []byte
}

func New() *PPC {
	return &PPC{
		RecMsg: make(chan []byte),
	}
}

// 接受socket的信息并传递到处理函数
func (ppc *PPC) Loop() {
	log.Infof("%d has enter loop\n", ppc.Player.UUID)
	for {
		select {
		case m := <-ppc.RecMsg:
			Msg := &variable.Message{}
			_ = json.Unmarshal(m, Msg)
			logicp.HandleMsg(*Msg, ppc.Player)
		}
	}
}
