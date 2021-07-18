package player

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
)

func PlayerLogin(openId, sessionKey string) bool {
	return loadPlayer(openId,sessionKey)
}

func loadPlayer(openId, sessionKey string) bool {
	_, err := dbapi.FindPlayer(openId)
	if err != nil {
		return newPlayer(openId)
	}
	return true
}

func newPlayer(openID string) bool {
	p := born(openID)
	_ = dbapi.SavePlayer(p)
	return true
}

func born(openID string) *rp.Player {
	return &rp.Player{
		OpenID: openID,

		Lv: 1,
		Exp: int64(0),

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