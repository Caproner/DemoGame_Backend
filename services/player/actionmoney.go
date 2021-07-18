package player

import (
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/localvar/timevar"
	"github.com/Caproner/DemoGame_Backend/services/ptime"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)

func handldMoneyGet(ctx *gin.Context, p *global.Player){
	lastuptime := ptime.GetTime(timevar.LastActionTimeType, p)
	r := proto.S2CPlayerInfo{
		Lv: p.Lv,
		Exp: p.Exp,
		UpdateTime: lastuptime,
	}
	responseresult.ResponseOk(ctx, proto.MSGS2CPlayerInfo, r)
}