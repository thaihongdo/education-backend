package api

import (
	"net/http"

	"sme-education-backend/cmd/entity-server/msg"
	auth_service "sme-education-backend/cmd/entity-server/services"
	"sme-education-backend/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `json:"email" valid:"email~Email is not valid"`
	Password string `json:"password" valid:"stringlength(6|50)~Password is at least 6 characters"`
}

type RegisterReq struct {
	Email     string `json:"email" valid:"email~Email is not valid"`
	Password  string `json:"password" valid:"stringlength(6|50)~Password is at least 6 characters"`
	FullName  string `json:"full_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func Login(c *gin.Context) {
	appG := Gin{C: c}

	var loginReq LoginReq
	isValid := appG.BindAndValidate(&loginReq)

	if isValid {
		service := auth_service.UserReq{Email: loginReq.Email, Password: loginReq.Password}
		user, err := service.Login()
		if err != nil {
			appG.Response(http.StatusUnauthorized, false, msg.GetMsg(msg.ERROR_AUTH_FAIL), nil, nil)
			return
		}

		j := utils.NewJWT()
		tokenInfo, err := j.GenerateToken(user.ID, user.Email, user.FullName)
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		user.Token = tokenInfo.Token
		user.ExpiredAt = tokenInfo.ExpiredAt
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), user, nil)
	}
}
func Register(c *gin.Context) {
	appG := Gin{C: c}

	var registerReq RegisterReq
	isValid := appG.BindAndValidate(&registerReq)

	if isValid {
		service := auth_service.UserReq{Email: registerReq.Email, Password: registerReq.Password, FirstName: registerReq.FirstName, LastName: registerReq.LastName, Phone: registerReq.Phone}
		isAdded, err := service.Register()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		if !isAdded {
			appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.ERROR_ADD_FAIL), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), nil, nil)
	}
}
