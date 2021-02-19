package model

import (
	"github.com/xuanxiaox/ginx/global"
	"gorm.io/gorm"
	"time"
)

type UserModeler interface {
	GetById() (user User, err error)
	List() (users []*User, err error)
}

type User struct {
	Id        int            `json:"id" gorm:"type:int(10) unsigned;primarykey"`
	Username  string         `json:"username" gorm:"type:varchar(30);not null;default:'';comment:姓名"`
	Password  string         `json:"password" gorm:"type:varchar(32);not null;default:'';comment:密码"`
	Status    int            `json:"status" gorm:"type:tinyint unsigned;not null;default:1;comment:状态 0-禁用 1-启用"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m *User) GetById() (user User, err error) {
	err = global.GDB.Model(m).Where("id = ?", m.Id).First(&user).Error
	return
}

func (m *User) List() (users []*User, err error) {
	err = global.GDB.Model(m).Find(&users).Error
	return
}
