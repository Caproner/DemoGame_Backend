package player

import (
	r "github.com/Caproner/DemoGame_Backend/include/global/r/player"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
)

func PlayerLogin(openId, sessionKey string) *r.Player {
	p := loadPlayer(openId)
	p.SessionKey = sessionKey
	return p
}

func loadPlayer(openId string) *r.Player {
	player, err := dbapi.FindPlayer(openId)
	if err != nil {
		p := newPlayer(openId)
		return p
	}
	return player
}

func newPlayer(openID string) *r.Player {
	uuID := dbapi.ItemLenAdd(variable.PlayerNumKey)
	uuID = variable.PlayerBaseID + uuID
	p := born(openID, uuID)
	_ = dbapi.SavePlayer(p)
	return p
}

func born(openID string, uuID int64) *r.Player {
	return &r.Player{
		OpenID: openID,
		UUID:   uuID,

		Bag:       make(map[int]interface{}),
		Money:     make(map[int]int64),
		Playing:   make(map[string]interface{}),
		Task:      make(map[int]interface{}),
		Cultivate: make(map[int]interface{}),
		Mail:      make(map[int64]interface{}),
		TimeClock: make(map[string]interface{}),

		Log:  make([]interface{}, 0),
		Goal: make([]interface{}, 0),
	}
}
