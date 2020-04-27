package service

type UserService interface {
	GetUserInfo(id string) (bool, error)
}
