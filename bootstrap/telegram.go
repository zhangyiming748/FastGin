package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/controller"
)

func InitTelegram(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.TelegramController)
		//routeGroup.GET("/v1/s1/gethello", c.GetHello)
		routeGroup.POST("/v1/telegram/download", c.DownloadAll)
	}
}
