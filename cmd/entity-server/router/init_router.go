package router

import (
	"fmt"
	"sme-education-backend/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(env string) *gin.Engine {
	var Router = gin.Default()
	Router.Use(utils.Cors())
	urlPrefix := fmt.Sprintf("%s/api/v1", env)
	ApiGroup := Router.Group(urlPrefix)

	initAuthRouter(ApiGroup)
	initUserRouter(ApiGroup)
	return Router
}
