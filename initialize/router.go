package initialize

import (
	"github.com/gin-gonic/gin"
	"go-back/global"
	"go-back/router"
	"net/http"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.GroupApp.System

	PublicGroup := Router.Group(global.GB_CONFIG.System.RouterPrefix)
	PublicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	systemRouter.InitBaseRouter(PublicGroup)

	global.GB_LOG.Info("router register success")
	return Router
}
