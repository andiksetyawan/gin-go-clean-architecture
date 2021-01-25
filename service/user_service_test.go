package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"solid-go/entity"
	"solid-go/repository"
	"testing"
)

//example unit test Mock

var userRepositoryMock = repository.UserRepositoryMock{Mock: mock.Mock{}}
var userServiceMock = NewUserService(&userRepositoryMock)

func TestUserService_FindUserByID_Error(t *testing.T) {
	guid := "d49135e2-7b07-4252-bfd0-ebbff33224d2"
	userRepositoryMock.Mock.On("FindUserByID", guid).Return(nil)

	user, err := userServiceMock.FindUserByID(guid)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserService_FindUserByID_Success(t *testing.T) {
	guid := "d49135e2-7b07-4252-bfd0-ebbff33224d2"
	user := entity.User{
		Guid:     "d49135e2-7b07-4252-bfd0-ebbff33224d2",
		Email:    "andiksetyawandev@gmail.com",
		Password: "12348123",
		Name:     "Andik Setyawan",
		Address:  "Jakarta",
		Photo:    "photo.jpg",
	}
	userRepositoryMock.Mock.On("FindUserByID", guid).Return(user)

	res, err := userServiceMock.FindUserByID(guid)
	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Equal(t, user.Guid, res.Guid)
	assert.Equal(t, user.Email, res.Email)
	assert.Equal(t, user.Password, res.Password)
	assert.Equal(t, user.Name, res.Name)
	assert.Equal(t, user.Address, res.Address)
	assert.Equal(t, user.Photo, res.Photo)
}
