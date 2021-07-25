package player

import (
	rp "github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/localvar/timevar"
	"github.com/Caproner/DemoGame_Backend/services/ptime"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"time"
)

func PlayerLogin(openId, sessionKey string) (key string) {
	p := loadPlayer(openId,sessionKey)
	var has bool
	if has = dbapi.UpdateToken(openId); !has{
		ptime.SetTime(timevar.LoginTimeType, time.Now().Unix(), p)
		ltt := ptime.GetTime(timevar.LastActionTimeType, p)
		ptime.SetTime(timevar.OfflineTimeType, ltt, p)
	}
	key = dbapi.UserToken(openId)
	if savef := dbapi.SavePlayer(p); savef != nil{
		log.Info(savef)
	}
	return
}

func loadPlayer(openId, sessionKey string) *rp.Player {
	p, err := dbapi.FindPlayer(openId)
	if err != nil {
		return newPlayer(openId)
	}
	return p
}

func newPlayer(openID string) *rp.Player{
	p := born(openID)
	_ = dbapi.SavePlayer(p)
	return p
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