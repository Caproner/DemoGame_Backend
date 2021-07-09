package player

import "github.com/Caproner/DemoGame_Backend/services/playersvr"

func PlayerLogin(openId, sessionKey string) int64 {
	p := playersvr.LoadPlayer(openId)
	playersvr.LoginPlayer(p)
	return p.UUID
}
