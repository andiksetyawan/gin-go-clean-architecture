package repository

import (
	"context"
	"gin-go-clean-architecture/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository interface {
	CreateUser(user entity.User) (*entity.User, error)
	FindUserByID(id string) (*entity.User, error)
	FindUser() (*[]entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	Db *mongo.Collection
}

func NewUserRepo(db *mongo.Database) UserRepository {
	return &userRepository{
		Db: db.Collection("user"),
	}
}

func (repository *userRepository) CreateUser(user entity.User) (*entity.User, error) {
	res, err := repository.Db.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	log.Println(res)
	return nil, err
}

func (repository *userRepository) FindUserByID(id string) (*entity.User, error) {
	var user entity.User
	err := repository.Db.FindOne(context.TODO(), bson.M{"guid": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *userRepository) FindUser() (*[]entity.User, error) {
	cur, err := repository.Db.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var users []entity.User
	cur.All(context.TODO(), &users)
	return &users, nil
}

func (repository *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := repository.Db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
