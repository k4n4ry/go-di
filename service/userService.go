package service

import "github.com/knry0329/go-di/repository"

type UserService interface {
	GetUserName(id int) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUserName(id int) (string, error) {
	m, err := s.repo.FindUserForId(id)
	if err != nil {
		return "", err
	}
	return m.Name, nil
}
