package store

import (
	"Shorterism/internal/database"
	"Shorterism/internal/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Mognodb struct {
	database *database.Mongodb
}

func New(Database *database.Mongodb) *Mognodb {
	return &Mognodb{database: Database}
}

func (m *Mognodb) NewLink(Data model.Url) error {
	Ctx, done := context.WithTimeout(context.Background(), time.Second*10)
	defer done()

	LinksCol := m.database.GetCl("links")

	var Result model.Url
	LinksCol.FindOne(Ctx, bson.D{{"Name", Data.Name}}).Decode(&Result)
	if Result.Name == Data.Name {
		expire := time.Unix(Result.Expire, 0)
		if !expire.IsZero() && expire.Before(time.Now()) {
			LinksCol.DeleteOne(Ctx, bson.D{{"Name", Result.Name}})
		} else {
			return errors.New("one Link is active with this Name")
		}
	}

	_, err := LinksCol.InsertOne(Ctx, Data)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mognodb) Link(Name string) (model.Url, error) {
	Ctx, done := context.WithTimeout(context.Background(), time.Second*10)
	defer done()

	LinksCol := m.database.GetCl("links")

	var Result model.Url
	LinksCol.FindOne(Ctx, bson.D{{"Name", Name}}).Decode(&Result)
	if Result.Name == Name {
		expire := time.Unix(Result.Expire, 0)
		if !expire.IsZero() && expire.Before(time.Now()) {
			LinksCol.DeleteOne(Ctx, bson.D{{"Name", Name}})
			return Result, errors.New("url Has Expired")
		}
		return Result, nil
	}
	return Result, errors.New("link Not Found")
}

func (m *Mognodb) RegisterUser(User model.User) error {
	Ctx, done := context.WithTimeout(context.Background(), time.Second*10)
	defer done()

	var Result model.User

	UsersCol := m.database.GetCl("users")

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

func (m *Mognodb) Login(User model.LoginModel) (bool, error) {
	Ctx, done := context.WithTimeout(context.Background(), time.Second*10)
	defer done()

	UsersCol := m.database.GetCl("users")

	var Result model.User
	err := UsersCol.FindOne(Ctx, bson.D{{"Email", User.Email}, {"Pass", User.Pass}}).Decode(&Result)
	if err != nil {
		return false, err
	}

	if Result.Email == User.Email {
		return true, nil
	}
	return false, nil
}
