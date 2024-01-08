package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/controller"
)

func InitClipboard(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.ClipboardController)
		routeGroup.POST("/v1/clipboard/upload", c.Upload)
	}
}
