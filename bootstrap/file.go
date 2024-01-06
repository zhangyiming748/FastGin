package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/controller"
)

func InitFile(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.FileController)
		routeGroup.GET("/v1/file/upload", c.Upload)
	}
}
