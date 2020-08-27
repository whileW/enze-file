package router

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-file/api/v1"
)

func InitFileRouter(Router *gin.RouterGroup) {
	File := Router
	{
		File.POST("v1/upload", v1.Upload)   //上传
		File.GET("v1/get/:code",v1.Get)		//获取文件
	}
}