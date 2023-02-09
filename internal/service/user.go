package service

import (
	"github.com/xuandax/ginx/internal/model"
)

type UserServicer interface {
	GetById(id int) (user model.User, err error)
	List() (users []*model.User, err error)
}

type UserService struct {
	UserModeler model.UserModeler
}

func (s *UserService) GetById(id int) (user model.User, err error) {
	s.UserModeler = &model.User{
		Id: id,
	}
	return s.UserModeler.GetById()
}

func (s *UserService) List() (users []*model.User, err error) {
	return s.UserModeler.List()
}
