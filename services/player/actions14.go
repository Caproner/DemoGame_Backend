package player

import (
	"errors"
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/services/pmoney"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/Caproner/DemoGame_Backend/utils/tr"
	"github.com/gin-gonic/gin"
)

// 处理玩家货币新增删减
func handleMoneyOpt(ctx *gin.Context, p *global.Player) {
	r := &proto.C2SMoneyOpt{}
	_  = ctx.BindJSON(r)
	lastMoney := pmoney.GetTypeMoney(tr.InterfaceToInt(r.Type), p)
	switch tr.InterfaceToInt(r.Opt) {
	case 1: // 增加操作
		newMoney := lastMoney + tr.InterfaceToInt64(r.Num)
		_ = pmoney.SetTypeMoney(tr.InterfaceToInt(r.Type), newMoney, p)
		responseresult.ResponseOk(ctx, proto.MSGS2COk, proto.S2COk{})
	case 2: // 扣除操作
		if lastMoney >= tr.InterfaceToInt64(r.Num){
			newMoney := lastMoney - tr.InterfaceToInt64(r.Num)
			_ = pmoney.SetTypeMoney(tr.InterfaceToInt(r.Type), newMoney, p)
			responseresult.ResponseOk(ctx, proto.MSGS2COk, proto.S2COk{})
		}else{
			responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: errors.New("del money error")})
		}
	default:
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: errors.New("not valiable opt error")})
	}
	log.Info(r)
}
