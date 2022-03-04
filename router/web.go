package router

import (
	"carefree/app/controller"
	"github.com/gin-gonic/gin"
)

func web(r *gin.Engine) {
	r.GET("/home", controller.NewDemoController().Home)
	r.POST("/login", controller.NewAuthController().Login)

	r.POST("/refresh_token", controller.NewAuthController().Refresh)
	r.GET("/test", controller.NewDemoController().Test)
}
