package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuandax/ginx/g"
	"github.com/xuandax/ginx/internal/routers"
	"github.com/xuandax/ginx/pkg/config"
	"github.com/xuandax/ginx/pkg/db"
	"github.com/xuandax/ginx/pkg/log"
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

// 1 设置配置文件
func setConfig() {
	g.ServerConfig = config.ViperServerConfig()
	g.DBConfig = config.ViperDBConfig()
}

// 设置zap日志
func setLog() {
	g.Log = log.NewZapSugarLogger()
	g.Logger = log.NewZapLogger()
}

// 设置db
func setGDB() {
	g.GDB = db.NewDb()
	g.RedisPool = db.NewRedisPool()
}
