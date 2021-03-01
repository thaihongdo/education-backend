package ginutil

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, httpCode int, success bool, message string, data interface{}, err error) {
	c.JSON(httpCode, gin.H{
		"success": success,
		"message": message,
		"data":    data,
		"error":   err,
	})
	return
}
