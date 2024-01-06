package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/bootstrap"
)

func main() {

	// gin服务
	engine := gin.New()

	bootstrap.InitService1(engine)

	// 启动http服务

	engine.Run(":8192")

}
