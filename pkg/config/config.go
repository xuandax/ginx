package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xuandax/ginx/g"
)

func ViperServerConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./configs/server.yaml")
	err := v.ReadInConfig()
	if err != nil {
		g.Log.Fatalf("ViperServerConfig v.ReadInConfig err:%s", err)
	}
	watchChange(v)
	return v
}

func ViperDBConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./configs/db.yaml")
	err := v.ReadInConfig()
	if err != nil {
		g.Log.Fatalf("ViperDBConfig v.ReadInConfig err:%s", err)
	}
	watchChange(v)
	return v
}

// 监听配置文件变化
func watchChange(v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		v.SetConfigFile(in.Name)
		err := v.ReadInConfig()
		if err != nil {
			g.Log.Fatalf("watchChange v.OnConfigChange v.ReadInConfig err:%s", err)
		}
	})
}
