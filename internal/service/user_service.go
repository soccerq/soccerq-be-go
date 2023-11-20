package service

import "soccerq/internal/repository"

type UserService interface {
	GetAllUser()
	GetUser(id int64)
}

type service struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &service{repo: userRepo}
}

func (service *service) GetAllUser() {}

func (service *service) GetUser(id int64) {}
