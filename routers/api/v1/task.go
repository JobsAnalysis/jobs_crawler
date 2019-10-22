package v1

import (
	"jobs-crawler/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	data := make(map[string]interface{})
	data["list"] = models.GetTaks()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
