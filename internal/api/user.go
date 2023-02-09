package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuandax/ginx/internal/model"
	"github.com/xuandax/ginx/internal/service"
	"strconv"
)

type User struct {
	UserServicer service.UserServicer
}

func NewUser() *User {
	return &User{
		UserServicer: &service.UserService{UserModeler: new(model.User)},
	}
}

func (a *User) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithError(400, errors.New("id类型错误"))
	}
	hongbaoCrontab, err := a.UserServicer.GetById(id)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, gin.H{
		"data": hongbaoCrontab,
	})
}

func (a *User) List(ctx *gin.Context) {
	list, err := a.UserServicer.List()
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200, gin.H{
		"list": list,
	})
}
