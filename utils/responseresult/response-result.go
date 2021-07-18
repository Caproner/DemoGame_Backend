package responseresult

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
非socket接口的信息返回的接口
*/

// ResponseOk 返回便准ok信息 200
func ResponseOk(ctx *gin.Context, protoID int ,result interface{}) {
	// fmt.Println(result)
	ctx.JSON(http.StatusOK, gin.H{
		"protoId": protoID,
		"data":    result,
	})
}

// ResponseFalse 返回错误信息
func ResponseFalse(ctx *gin.Context, state int, protoID int, result interface{}) {
	ctx.JSON(state, gin.H{
		"protoId": protoID,
		"data":    result,
	})
}
