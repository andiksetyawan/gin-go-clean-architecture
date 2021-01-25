package repository

import (
	"errors"
	"gin-go-clean-architecture/entity"
	"github.com/stretchr/testify/mock"
	"log"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindUserByID(id string) (*entity.User, error) {
	arguments := repository.Mock.Called(id)
	log.Println(arguments)
	log.Println(arguments.Get(0))
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
