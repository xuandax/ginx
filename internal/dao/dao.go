package dao

import (
	"context"
	"github.com/xuandax/ginx/g"
	"gorm.io/gorm"
)

type Dao struct {
}

func (dao Dao) DB() *gorm.DB {
	return g.GDB
}

func (dao Dao) Ctx(ctx context.Context) *gorm.DB {
	return g.GDB.WithContext(ctx)
}
