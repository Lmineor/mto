package router

import (
	v1 "github.com/Lmineor/mto/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func InitDwzRouter(Router *gin.RouterGroup) {
	glog.V(3).Info("init dwz router")
	DwzRouter := Router.Group("dwz")
	{
		DwzRouter.POST("generate", v1.Generate)
		DwzRouter.GET("recovery", v1.Recovery)
	}
}
