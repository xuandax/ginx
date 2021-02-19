package global

import (
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var (
	GDB       *gorm.DB
	RedisPool *redis.Pool
)
