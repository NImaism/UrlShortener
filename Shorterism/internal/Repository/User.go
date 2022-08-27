package Repository

import (
	database "Shorterism/Database"
	"Shorterism/Model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	RegisterUser(User Model.User) error
	Login(User Model.LoginModel) (bool, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return userRepository{}
}

func (userRepository) RegisterUser(User Model.User) error {
	Ctx := context.TODO()
	var Result Model.User

	UsersCol := database.GetCl(database.Data, "users")

	UsersCol.FindOne(Ctx, bson.D{{"Email", User.Email}}).Decode(&Result)

	if Result.Email == User.Email {
		return errors.New("one account is active with this email")
	}

	_, err := UsersCol.InsertOne(Ctx, User)
	if err != nil {
		return err
	}
	return nil
}

func (userRepository) Login(User Model.LoginModel) (bool, error) {
	Ctx := context.TODO()
	UsersCol := database.GetCl(database.Data, "users")

	var Result Model.User
	err := UsersCol.FindOne(Ctx, bson.D{{"Email", User.Email}, {"Pass", User.Pass}}).Decode(&Result)
	if err != nil {
		return false, err
	}

	if Result.Email == User.Email {
		return true, nil
	}
	return false, nil
}
