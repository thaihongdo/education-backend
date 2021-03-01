package api

import (
	"net/http"
	"sme-education-backend/cmd/entity-server/config"
	"sme-education-backend/cmd/entity-server/msg"
	"sme-education-backend/cmd/entity-server/services"
	"sme-education-backend/internal/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	appG := Gin{C: c}
	offset, limit := utils.GetPage(c, config.GetConfig().DefaultPageNum, config.GetConfig().DefaultPageLimit)
	service := services.UserReq{
		PageNum:  offset,
		PageSize: limit,
	}
	list, err := service.GetAllUser()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_GET_FAIL), nil, nil)
		return
	}
	total, err := service.GetTotal()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_GET_FAIL), nil, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = list
	data["total"] = total

	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), data, nil)

}
func Get(c *gin.Context) {
	appG := Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	service := services.UserReq{}
	service.ID = uint(id)
	objRes, err := service.Get()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
}

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

func DeleteUser(c *gin.Context) {
	appG := Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	var service services.UserReq
	service.ID = uint(id)
	_, err = service.DeleteUser()
	if err != nil {
		appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}
	appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), nil, nil)
}
