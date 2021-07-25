package player

import (
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
	"github.com/Caproner/DemoGame_Backend/utils/log"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)

// MsgHandler获得openid以获得玩家数据进行处理
func MsgHandler(ctx *gin.Context, Proto int){
	openID := ctx.Query(variable.OPENID)
	if p, err := dbapi.FindPlayer(openID);err == nil{
		doHandler(ctx,Proto,p)
	}else{
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error: "not find player"})
	}
}

func doHandler(ctx *gin.Context, Proto int, p *global.Player){
	//log.Info(Proto)
	switch Proto {
	// 12+
	case proto.MSGC2SPlayerInfo: // 获取玩家所有数据
		handleInfoGet(ctx, p)
	// 13+
	case proto.MSGC2SLoopSync: // 处理玩家数据同步
		handleLoopSync(ctx, p)
	case proto.MSGC2SMoneySync: // 处理玩家货币数据同步
		handleMoneySync(ctx, p)
	case proto.MSGC2SLvSync: // 处理玩家等级和经验
		handleLvExpSync(ctx, p)
	case proto.MSGC2SBuildsSync: // 处理玩家已解锁建筑
		handleBuildsSync(ctx, p)
	case proto.MSGC2SNpcsSync: // 处理玩家已解锁npc
		handleNpcsSync(ctx, p)
	// 14+
	case proto.MSGC2SMoneyOpt: // 处理玩家货币变更
		handleMoneyOpt(ctx, p)

	default:
		responseresult.ResponseOk(ctx, proto.MSGS2CErr, proto.S2CErr{Error:"not find right proto"})
	}
	if savef := dbapi.SavePlayer(p); savef != nil{
		log.Info(savef)
	}
}