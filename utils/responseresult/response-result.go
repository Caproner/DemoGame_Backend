package responseresult

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOk(ctx *gin.Context, result interface{}) {
	fmt.Println(result)
	ctx.JSON(http.StatusOK, gin.H{
		"protoId": 0,
		"data":    result,
	})
}

func ResponseFalse(ctx *gin.Context, state int, result interface{}) {
	ctx.JSON(state, gin.H{
		"protoId": 0,
		"data":    result,
	})
}
