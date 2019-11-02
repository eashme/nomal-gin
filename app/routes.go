package app

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"nomal-gin/views"
)

// 初始化路由
func (a *app) initRouter() {
	g := gin.New()
	// 绑定全局handler
	g.Use(gin.Recovery(), gin.Logger(), handleCors())
	// 启用swagger api文档
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 心跳检测
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 创建版本路由组
	route := g.Group(config.Version)
	// 初始化服务,并为服务绑定路由
	a.initService(route)

	a.g = g
}

func (a *app) initService(route gin.IRouter) {

	// 根据service层构建 view层
	uv := views.NewUserView(a.db, a.cache)

	// 为view层绑定路由
	uv.BindRoute(route)
}
