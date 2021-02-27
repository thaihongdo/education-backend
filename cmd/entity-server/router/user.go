package router

import (
	"sme-education-backend/cmd/entity-server/api"
	ginutil "sme-education-backend/internal/pkg/gin-util"

	"github.com/gin-gonic/gin"
)

func initUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(ginutil.JWTAuth())
	{
		userRouter.POST("/add", api.AddUser)
		userRouter.GET("/all", api.GetAllUser)
		userRouter.GET("/detail", api.Get)
		userRouter.PUT("/update/:id", api.UpdateUser)
		userRouter.DELETE("/remove/:id", api.DeleteUser)
	}

}
