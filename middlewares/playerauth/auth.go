package playerauth

import (
	"github.com/Caproner/DemoGame_Backend/data/proto"
	"github.com/Caproner/DemoGame_Backend/include/variable"
	"github.com/Caproner/DemoGame_Backend/utils/database/dbapi"
	"github.com/Caproner/DemoGame_Backend/utils/responseresult"
	"github.com/gin-gonic/gin"
)

/*
在这个模块会每次都校对前端发过来的token
*/

func UserToken() gin.HandlerFunc{
	return func(context *gin.Context) {
		openID := context.Query(variable.OPENID)
		if openID == ""{responseresult.ResponseOk(context, proto.MSGS2CErr, "not this code")}
		if b := dbapi.CheckUserToken(openID);b{
			context.Next()
		}else{
			context.Abort()
			responseresult.ResponseOk(context, proto.MSGS2CErr, "token time out")
		}
	}
}