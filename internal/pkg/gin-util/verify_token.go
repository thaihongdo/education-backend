package ginutil

import (
	"net/http"
	jwt "sme-education-backend/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

const (
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 2002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 2003
	ERROR_AUTH_TOKEN               = 2004
)

var MsgFlags = map[int]string{
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token is invalid",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token is time out, please login and try again",
	ERROR_AUTH_TOKEN:               "Token is error, please try again",
}

func GetMsg(code int) string {
	msg, _ := MsgFlags[code]
	return msg
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			Response(c, http.StatusUnauthorized, false, GetMsg(ERROR_AUTH_TOKEN), gin.H{
				"reload": true,
			}, nil)
			c.Abort()
			return
		}
		j := jwt.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				Response(c, http.StatusUnauthorized, false, GetMsg(ERROR_AUTH_CHECK_TOKEN_TIMEOUT), gin.H{
					"reload": true,
				}, nil)

				c.Abort()
				return
			}
			Response(c, http.StatusUnauthorized, false, GetMsg(ERROR_AUTH_CHECK_TOKEN_FAIL), gin.H{
				"reload": true,
			}, nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
