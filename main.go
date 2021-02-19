package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanxiaox/ginx/global"
	"github.com/xuanxiaox/ginx/internal/routers"
	"github.com/xuanxiaox/ginx/pkg/config"
	"github.com/xuanxiaox/ginx/pkg/db"
	"github.com/xuanxiaox/ginx/pkg/log"
)

func main() {
	r := gin.Default()
	routers.Router(r)
	setGlobal()
	r.Run(":8080")
}

func setGlobal() {
	//1 设置配置文件
	setConfig()
	//设置zap日志
	setLog()
	//设置db
	setGDB()
}

//1 设置配置文件
func setConfig() {
	global.ServerConfig = config.ViperServerConfig()
	global.DBConfig = config.ViperDBConfig()
}

//设置zap日志
func setLog() {
	global.Log = log.NewZapSugarLogger()
	global.Logger = log.NewZapLogger()
}

//设置db
func setGDB() {
	global.GDB = db.NewDb()
	global.RedisPool = db.NewRedisPool()
}
