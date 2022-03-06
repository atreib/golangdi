package services

import (
	"github.com/atreib/golangdi/entities"
	"github.com/atreib/golangdi/infra/libs"
	"github.com/atreib/golangdi/infra/repositories"
)

type AuthenticateService struct {
	UserRepository *repositories.UserRepository
	JwtEncrypter   *libs.JwtEncrypter
}

func (s *AuthenticateService) Authenticate(login, password string) (*entities.Authentication, error) {
	user, err := s.UserRepository.GetUser(login, password)

	if err != nil {
		return nil, err
	}

	jwt, err := s.JwtEncrypter.Generate()

	if err != nil {
		return nil, err
	}

	return &entities.Authentication{
		User: *user,
		Jwt:  jwt,
	}, nil
}
