package routers

import (
	"jobs-crawler/models"
	v1 "jobs-crawler/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// gin.SetMode(setting.RunMode)
	models.Init()
	apiv1 := r.Group("app/v1")
	{
		//获取标签列表
		apiv1.GET("/task", v1.GetTasks)
		//新建标签
	}
	return r
}
