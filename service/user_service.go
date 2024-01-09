package service

import (
	"sayembara/entity/mapping"
	"sayembara/entity/response"
	"sayembara/repository"
)

type UserService interface {
	GetUsers() ([]response.GetUsers, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) GetUsers() ([]response.GetUsers, error) {
	users, err := s.userRepository.GetUsers()

	var usersMap []response.GetUsers

	for _, user := range users {
		userMap := mapping.UsersMap(user)
		usersMap = append(usersMap, userMap)
	}
	return usersMap, err
}
