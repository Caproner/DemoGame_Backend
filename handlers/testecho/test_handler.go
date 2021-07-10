package testecho

import (
	"github.com/Caproner/DemoGame_Backend/utils/database/ormredis"
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
		testRouter.GET("/ts", testKvHandler)
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

func testKvHandler(context *gin.Context) {
	key := context.Query("key")
	value := context.Query("value")
	_ = ormredis.RDB().KVSet(key,value)
	vv,_ := ormredis.RDB().KVGet(key)
	context.JSON(http.StatusOK, gin.H{
		"kv" : vv,
	})
}