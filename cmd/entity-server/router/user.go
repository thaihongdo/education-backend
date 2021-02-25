package router

import (
	"sme-education-backend/cmd/entity-server/api"

	"github.com/gin-gonic/gin"
)

func initUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.POST("/add", api.AddUser)
		userRouter.PUT("/update/:id", api.UpdateUser)
	}

}
