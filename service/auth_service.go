package service

import (
	"github.com/ctirouzh/tiny-url/dto"
	"github.com/ctirouzh/tiny-url/model"
)

type AuthService struct {
	userRepo   model.UserRepository
	jwtService model.JwtService
}

func NewAuthService(r model.UserRepository, j model.JwtService) *AuthService {
	return &AuthService{
		userRepo:   r,
		jwtService: j,
	}
}

func (s *AuthService) SignUp(userDto *dto.SignUp) (*model.User, error) {
	return s.userRepo.CreateUser(userDto)
}

func (s *AuthService) Login(credentials *dto.SignIn) (*model.AccessTokenResponse, error) {
	var token *model.AccessTokenResponse
	user, err := s.userRepo.ValidateUser(credentials)
	if err != nil {
		return nil, err
	}
	token, err = s.jwtService.GenerateJwtToken(user)
	if err != nil {
		return nil, err
	}
	return token, nil
}
