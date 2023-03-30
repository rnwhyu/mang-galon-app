package controller

import "github.com/gin-gonic/gin"

func ApiResponse(c *gin.Context, code int, message string, data ...interface{}) {
	json := map[string]interface{}{
		"message": message,
	}

	if len(data) > 0 {
		json["data"] = data[0]
	}

	c.JSON(code, json)
}

func ApiErrorResponse(c *gin.Context, code int, message string, data ...interface{}) {
	json := map[string]interface{}{
		"message": message,
	}

	if len(data) > 0 {
		json["error"] = data[0]
	}

	c.AbortWithStatusJSON(code, json)
}
