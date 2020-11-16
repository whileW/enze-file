package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-file/model"
	"github.com/whileW/enze-file/router"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/initialize"
	"github.com/whileW/enze-global/log"
	"github.com/whileW/enze-global/utils/resp"
	_"github.com/whileW/enze-file/service/qiniu"
	_"github.com/whileW/enze-file/service/local"
)

func main(){
	global.GVA_LOG.Info(global.GVA_CONFIG.Setting)
	init_db_tables()
	init_server()
}

// 注册数据库表
func init_db_tables() {
	initialize.Db()
	db := global.GVA_DB.Get(model.FileDb)
	db.AutoMigrate(model.File{})
	global.GVA_LOG.Info("register table success")
}

//加载http监听
func init_server() {
	//配置gin
	gin.DefaultWriter = &log.GinLog{}
	gin.DefaultErrorWriter = &log.GinErrLog{}
	if global.GVA_CONFIG.SysSetting.Env != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}


	r := gin.Default()
	// 跨域
	r.Use(resp.Cors())
	//捕获异常
	r.Use(gin.Recovery())
	//挂载路由
	ApiGroup := r.Group("")
	router.InitFileRouter(ApiGroup)				//file

	global.GVA_LOG.Error(r.Run(":"+global.GVA_CONFIG.SysSetting.HttpAddr))
}