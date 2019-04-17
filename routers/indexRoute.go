package routers

import (
	"github.com/gin-gonic/gin"
	controller "github.com/lethanhsonthd/EOS_API_create_account/Controller"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.POST("/account/create", controller.CreateAccount)
	}
	return router
}
