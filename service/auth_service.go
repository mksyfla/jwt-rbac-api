package service

import (
	"errors"
	"os"
	"sayembara/entity/model"
	"sayembara/entity/request"
	"sayembara/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Create(bodyRequest request.UserRegisterRequest) (string, error)
	Login(bodyRequest request.UserLoginRequest) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *authService {
	return &authService{userRepository}
}

func (s *authService) Create(bodyRequest request.UserRegisterRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password), 10)
	if err != nil {
		return "", err
	}

	if !s.userRepository.IsEmailAvailable(bodyRequest.Email) {
		return "", errors.New("email is used")
	}

	imageUrl := "public/image/default/default-profile.jpg"

	user := model.UserPassword{
		User: model.User{
			Name:     bodyRequest.Name,
			Email:    bodyRequest.Email,
			Banner:   imageUrl,
			Profile:  imageUrl,
			Category: bodyRequest.Category,
		},
		Password: string(password),
	}

	id, err := s.userRepository.Create(user)
	if err != nil {
		return "", err
	}

	return id, err
}

func (s *authService) Login(bodyRequest request.UserLoginRequest) (string, error) {
	result, err := s.userRepository.GetUserByEmail(bodyRequest.Email)
	if err != nil {
		return "", errors.New("email or password wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(bodyRequest.Password))
	if err != nil {
		return "", errors.New("email or password wrong")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": result.Id,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
