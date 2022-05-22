package router

import (
	"github.com/Lmineor/mto/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	glog.V(1).Info("use middleeare cors")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	// TODO use swagger api
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		InitDwzRouter(PublicGroup)
	}
	glog.V(1).Info("register router success")
	return Router
}
