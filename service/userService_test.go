package service

import (
	"testing"

	"github.com/knry0329/go-di/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Verify(t *testing.T) {
	s := &userService{repo: &mock.MockUserRepository{}}
	m, err := s.GetUserName(100)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "taro", m)
}
