package service

import (
	"gin-go-clean-architecture/entity"
	"gin-go-clean-architecture/repository"
	"log"
)

type UserService interface {
	CreateUser(user entity.User) (*entity.User, error)
	FindUser() (*[]entity.User, error)
	FindUserByID(guid string) (*entity.User, error)
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

type userService struct {
	repository repository.UserRepository
}

func (u userService) CreateUser(user entity.User) (*entity.User, error) {

	//upload to minio

	res, err := u.repository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	log.Println("service signup", res)
	return nil, err
}

func (u userService) FindUser() (*[]entity.User, error) {
	users, err := u.repository.FindUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u userService) FindUserByID(guid string) (*entity.User, error) {
	user, err := u.repository.FindUserByID(guid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
