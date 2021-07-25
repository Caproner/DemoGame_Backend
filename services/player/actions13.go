package player

import (
	"errors"
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/localvar/timevar"
	"github.com/Caproner/DemoGame_Backend/services/pcul"
	"github.com/Caproner/DemoGame_Backend/services/plv"
	"github.com/Caproner/DemoGame_Backend/services/pmoney"
	"github.com/Caproner/DemoGame_Backend/services/ptime"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
	"github.com/gin-gonic/gin"
	"time"
)

func handleLoopSync(ctx *gin.Context, p *global.Player){
	r := ptime.SetTime(timevar.LastActionTimeType,time.Now().Unix(), p)
	responseresult.ResponseOk(ctx, proto.MSGS2COk, proto.S2COk{})
	log.Info(r)
}

// 处理玩家货币同步
func handleMoneySync(ctx *gin.Context, p *global.Player){
	r := &proto.C2SMoneySync{}
	_  = ctx.BindJSON(r)
	if pmoney.SyncMoney(tr.InterfaceToMIntInt64(r.Money), p){
		responseresult.ResponseOk(ctx, proto.MSGS2COk, proto.S2COk{})
	}else{
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: errors.New("sync money error")})
	}
	log.Info(r)
}

//
func handleLvExpSync(ctx *gin.Context, p *global.Player){
	r := &proto.C2SLvSync{}
	_  = ctx.BindJSON(r)
	if plv.SyncLv(tr.InterfaceToInt(r.Lv), tr.InterfaceToInt64(r.Exp),p){
		responseresult.ResponseOk(ctx, proto.MSGS2COk, proto.S2COk{})
	}else{
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: errors.New("sync lv error")})
	}
	log.Info(r)
}

func handleBuildsSync(ctx *gin.Context, p *global.Player){
	r := &proto.C2SBuildsSync{}
	_  = ctx.BindJSON(r)
	if pcul.SyncBuildS(tr.InterfaceToIntList(r.Builds), p){
		responseresult.ResponseOk(ctx, proto.MSGS2COk, proto.S2COk{})
	}else{
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: errors.New("sync builds error")})
	}
	log.Info(r)
}

func handleNpcsSync(ctx *gin.Context, p *global.Player){
	r := &proto.C2SNpcsSync{}
	_  = ctx.BindJSON(r)
	if pcul.SyncNpcS(tr.InterfaceToIntList(r.Npcs), p){
		responseresult.ResponseOk(ctx, proto.MSGS2COk, proto.S2COk{})
	}else{
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: errors.New("sync builds error")})
	}
	log.Info(r)
}