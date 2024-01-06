package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/controller"
)

func InitService1(engine *gin.Engine) {
	routeGroup := engine.Group("/api")

	{
		c := new(controller.S1Controller)
		routeGroup.GET("/v1/s1/gethello", c.GetHello)
		routeGroup.POST("/v1/s1/posthello", c.PostHello)
	}
}
