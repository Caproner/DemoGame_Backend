package testecho

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandlerRsp struct {
	Msg interface{} `json:"msg"`
}

func InitRouter(router *gin.Engine) *gin.Engine {
	testRouter := router.Group("/api/test").Use()
	{
		testRouter.GET("/echo", TestHandler)
	}
	return router
}

// Test Echo Handler
func TestHandler(c *gin.Context) {
	rsp := TestHandlerRsp{
		Msg: "rsp ok, this is demo backend",
	}
	c.JSON(http.StatusOK, gin.H{
		"ret":  0,
		"msg":  "ok",
		"data": rsp,
	})
}
