package db

import (
	"github.com/gomodule/redigo/redis"
	"github.com/xuanxiaox/ginx/global"
	"time"
)

func NewRedisPool() *redis.Pool {
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				global.DBConfig.GetString("redis.host"),
				redis.DialPassword(global.DBConfig.GetString("redis.password")),
				redis.DialDatabase(global.DBConfig.GetInt("redis.db")),
			)
			if err != nil {
				global.Log.Errorf("redis.Dial err:%v", err)
			}
			if _, err := c.Do("AUTH", global.DBConfig.GetString("redis.password")); err != nil {
				c.Close()
				global.Log.Errorf("c.Do AUTH err:%v", err)
				return nil, err
			}
			if _, err := c.Do("SELECT", global.DBConfig.GetInt("redis.db")); err != nil {
				c.Close()
				global.Log.Errorf("c.Do SELECT err:%v", err)
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         5,
		MaxActive:       20,
		IdleTimeout:     30,
		Wait:            false,
		MaxConnLifetime: 30,
	}
	return pool
}
