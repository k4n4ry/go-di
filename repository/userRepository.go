package repository

import (
	"github.com/knry0329/go-di/db"
	"github.com/knry0329/go-di/model"
)

type UserRepository interface {
	FindUserForId(id int) (*model.User, error)
}

//interfaceで戻す
func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

func (*userRepository) FindUserForId(id int) (*model.User, error) {
	db := db.GetDBInstance()
	var m model.User

	db.First(&m, id)
	// var u model.User

	return &m, nil
}
