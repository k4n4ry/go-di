package mock

import "github.com/knry0329/go-di/model"

type MockUserRepository struct{}

func (*MockUserRepository) FindUserForId(id int) (*model.User, error) {
	return &model.User{
		Id:   int64(id),
		Name: "taro",
	}, nil
}
