package playersvr

/*
管理玩家数据的进程
*/

import (
	"fmt"
	"sync"

	r "github.com/Caproner/DemoGame_Backend/app/global/record"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
)

var playerSumNumKey = "playerSumNum"

type playerSvr struct {
	loop      int64
	LineNum   int
	PlayerMap map[string]interface{}
	svrChan   chan r.Player
	mux       sync.Mutex
}

func (svr *playerSvr) Loop() {
	for {
		select {
		case player := <-svr.svrChan:
			fmt.Println(player.UUID)
			svr.PlayerMap[player.OpenId] = player
			svr.LineNum++
		}
	}
}

var defaultPlayerSvr *playerSvr

func init() {
	defaultPlayerSvr = &playerSvr{
		loop:      0,
		LineNum:   0,
		PlayerMap: make(map[string]interface{}),
		svrChan:   make(chan r.Player),
	}
}

func DefaultSvr() *playerSvr {
	return defaultPlayerSvr
}

func LoadPlayer(openId string) *r.Player {
	player, err := dbapi.FindPlayer(openId)
	if err != nil {
		p := NewPlayer(openId)
		return p
	}
	return player
}

var basePlayerNum = 10000

func NewPlayer(openId string) *r.Player {
	uuid := dbapi.ItemLenAdd(playerSumNumKey)
	p := &r.Player{OpenId: openId, UUID: int64(basePlayerNum) + uuid}
	_ = dbapi.SavePlayer(p)
	return p
}

func LoginPlayer(p *r.Player) {
	defaultPlayerSvr.mux.Lock()
	defer defaultPlayerSvr.mux.Unlock()
	defaultPlayerSvr.svrChan <- *p
}
