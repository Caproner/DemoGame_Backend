package player

import (
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/services/pcul"
	"github.com/Caproner/DemoGame_Backend/services/ptime"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)

func handleInfoGet(ctx *gin.Context, p *global.Player){
	r := proto.S2CPlayerInfo{
		Lv: p.Lv,
		Exp: p.Exp,
		Money: p.Money,
		Times: ptime.GetTimes(p),
		Builds: pcul.GetBuildS(p),
		Npcs: pcul.GetNpcS(p),
	}
	responseresult.ResponseOk(ctx, proto.MSGS2CPlayerInfo, r)
}