package main

import (
	"fmt"
	"net/http"

	"github.com/Caproner/DemoGame_Backend/routers"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 全局中间件加这里

	// 定时任务放这里

	// 找不到路由则返回404，并给出错误信息
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  "LogicError: API route not found",
		})
	})

	// TODO: 配置化
	routers.InitRoute(router)
	srv := &http.Server{
		Addr:    "0.0.0.0:17263",
		Handler: router,
	}
	if err := gracehttp.Serve(srv); err != nil {
		fmt.Println("Start Server Failed: %s", err)
		return
	}

	defer func(srv *http.Server) {
		err := recover()
		if err != nil {
			fmt.Println("Server Panic: %s", err)
			return
		}
		fmt.Println("Server Shutdown Success")
	}(srv)

	fmt.Println("Start Server Success")
}
