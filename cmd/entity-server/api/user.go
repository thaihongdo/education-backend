package api

import (
	"net/http"
	"sme-education-backend/cmd/entity-server/msg"
	"sme-education-backend/cmd/entity-server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	appG := Gin{C: c}

	service := services.UserReq{}
	isValid := appG.BindAndValidate(&service)
	if isValid {
		_, err := service.AddUser()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), nil, nil)
	}

}
func UpdateUser(c *gin.Context) {
	appG := Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	var service services.UserReq
	isValid := appG.BindAndValidate(&service)
	if isValid {
		service.ID = uint(id)
		_, err := service.UpdateUser()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), nil, nil)
	}
}
