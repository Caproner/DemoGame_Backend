package player

import (
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/global"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MsgHandler(ctx *gin.Context, Proto int){
	openID := ctx.Query("openID")
	if p, err := dbapi.FindPlayer(openID);err == nil{
		doHandler(ctx,Proto,p)
	}else{
		responseresult.ResponseFalse(ctx,http.StatusAccepted,0, "not find player")
	}
}

func doHandler(ctx *gin.Context, Proto int, p *global.Player){
	switch Proto {
	case proto.MSGC2SPlayerInfo:
		handldMoneyGet(ctx,p)
	default:
		responseresult.ResponseFalse(ctx, 200, 0,"not find right proto")
	}
}