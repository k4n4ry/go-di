package repository

import "github.com/knry0329/go-di/model"

type UserRepository interface {
	FindUserForId(id int) (*model.User, error)
}

//interfaceで戻す
func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (*userRepository) FindUserForId(id int) (*model.User, error) {
	// var u model.User

	return nil, nil
}
