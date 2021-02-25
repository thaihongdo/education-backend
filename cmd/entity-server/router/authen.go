package router

import (
	"sme-education-backend/cmd/entity-server/api"

	"github.com/gin-gonic/gin"
)

func initAuthRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("auth")
	{
		AuthRouter.POST("login", api.Login)       //login
		AuthRouter.POST("register", api.Register) //register
	}

}
