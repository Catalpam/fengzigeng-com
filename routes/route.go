package routes

import (
	"fengzigneg/godocker/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/iosmeeting/register", controller.Register)
	r.POST("/iosmeeting/login", controller.Login)

	

	return r
}
