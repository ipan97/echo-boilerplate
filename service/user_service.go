package service

import (
	"github.com/ipan97/echo-boilerplate/models"
	"github.com/ipan97/echo-boilerplate/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func (s UserService) FindAllUser() []*models.User {
	return nil
}
