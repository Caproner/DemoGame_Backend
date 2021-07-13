package playersvr

/*
管理玩家数据的进程
*/

import (
	r "github.com/Caproner/DemoGame_Backend/include/global/r/player"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
)

type PlayerSvr struct {
	PlayerMap  map[string]*r.Player
	Register   chan *r.Player
	UnRegister chan *r.Player
}

func (svr *PlayerSvr) Loop() {
	for {
		select {
		case p := <-svr.Register:
			puuid := tr.Int64ToString(p.UUID)
			log.Info("player_", p.UUID, "had register")
			svr.PlayerMap[puuid] = p
		}
	}
}

func New() *PlayerSvr {
	return &PlayerSvr{
		PlayerMap:  make(map[string]*r.Player),
		Register:   make(chan *r.Player),
		UnRegister: make(chan *r.Player),
	}
}
