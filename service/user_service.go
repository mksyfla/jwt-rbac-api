package service

import (
	"errors"
	"sayembara/entity/model"
	"sayembara/entity/request"
	"sayembara/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(bodyRequest request.UserRegisterRequest) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) Create(bodyRequest request.UserRegisterRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password), 10)

	if err != nil {
		return "", err
	}

	if !s.userRepository.IsEmailAvailable(bodyRequest.Email) {
		return "", errors.New("email is used")
	}

	user := model.UserPassword{
		User: model.User{
			Name:     bodyRequest.Nama,
			Email:    bodyRequest.Email,
			Banner:   "",
			Profile:  "",
			Category: bodyRequest.Category,
		},
		Password: string(password),
	}

	id, err := s.userRepository.Create(user)

	return id, err
}
