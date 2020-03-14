package repository

import (
	"github.com/ipan97/echo-boilerplate/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) FindAll() (*[]models.User, error) {
	var users []models.User
	err := r.Find(&users).Error
	return &users, err
}
