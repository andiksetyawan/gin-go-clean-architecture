package repository

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"solid-go/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindUserByID(id string) (*entity.User, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil, errors.New("error")
	} else {
		user := arguments.Get(0).(entity.User)
		return &user, nil
	}
}

func (repository *UserRepositoryMock) FindUserByEmail(email string) (*entity.User, error) {
	panic("implement me")
}

func (repository *UserRepositoryMock) CreateUser(user entity.User) (*entity.User, error) {
	panic("implement me")
}

func (repository *UserRepositoryMock) FindUser() (*[]entity.User, error) {
	panic("implement me")
}
