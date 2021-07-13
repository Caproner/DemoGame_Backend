package player

import (
	r "github.com/Caproner/DemoGame_Backend/include/global/r/player"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/utils/libsend"
)

func HandleMsg(msg variable.Message, p *r.Player) {
	// fmt.Println(msg)
	libsend.Send(p.UUID, msg)
}
