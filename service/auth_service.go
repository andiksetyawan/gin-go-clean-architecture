package service

import (
	"errors"
	"gin-go-clean-architecture/entity"
	"gin-go-clean-architecture/repository"
	"gin-go-clean-architecture/storage"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type AuthService interface {
	Login(email string, password string) (*string, error)
	SignUp(user entity.User, photo io.Reader) (*string, error)
}

type authService struct {
	repository repository.UserRepository
	jwt        JWTService
	storage    storage.Storage
}

func NewAuthService(repository repository.UserRepository, jwt JWTService, storage storage.Storage) AuthService {
	return &authService{repository: repository, jwt: jwt, storage: storage}
}

func (a *authService) Login(email string, password string) (*string, error) {
	user, err := a.repository.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	//compare bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	//generate jwt
	token := a.jwt.GenerateToken(user.Guid, user.Name, false)
	return token, nil
}

func (a *authService) SignUp(user entity.User, photo io.Reader) (*string, error) {
	if user, _ := a.repository.FindUserByEmail(user.Email); user != nil {
		return nil, errors.New("account is exist")
	}
	//upload to minio
	fileID := uuid.New().String()
	err := a.storage.Upload(fileID, photo, "user")
	if err != nil {
		return nil, err
	}

	user.Photo = fileID

	//new guid
	user.Guid = uuid.New().String()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	_, err = a.repository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	//generate jwt
	token := a.jwt.GenerateToken(user.Guid, user.Name, false)
	return token, nil
}
