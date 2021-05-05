package testecho

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandlerRsp struct {
	Msg interface{} `json:"msg"`
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
