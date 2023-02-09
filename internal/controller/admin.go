package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xuandax/ginx/internal/dao"
	"net/http"
)

var Admin = &cAdmin{}

type cAdmin struct{}

func (c *cAdmin) GetList(ctx *gin.Context) {
	dao.Admin.Ctx(ctx).Find(1)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "123",
	})
}
