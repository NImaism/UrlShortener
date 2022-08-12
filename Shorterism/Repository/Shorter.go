package Repository

import (
	database "Shorterism/Database"
	"Shorterism/Model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type LinkRepository interface {
	NewLink(Data Model.Url) error
	Link(Name string) (Model.Url, error)
}

type linkRepository struct{}

func NewLinkShorter() LinkRepository {
	return linkRepository{}
}

func (linkRepository) NewLink(Data Model.Url) error {
	Ctx := context.TODO()
	LinksCol := database.GetCl(database.Data, "links")

	var Result Model.Url
	LinksCol.FindOne(Ctx, bson.D{{"Name", Data.Name}}).Decode(&Result)
	if Result.Name == Data.Name {
		return errors.New("one Link is active with this Name")
	}

	_, err := LinksCol.InsertOne(Ctx, Data)
	if err != nil {
		return err
	}

	return nil
}

func (linkRepository) Link(Name string) (Model.Url, error) {
	Ctx := context.TODO()
	LinksCol := database.GetCl(database.Data, "links")

	var Result Model.Url
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
