package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Mongodb struct {
	database *mongo.Client
}

func New() *Mongodb {
	ctx, done := context.WithTimeout(context.Background(), time.Second*10)
	defer done()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}

	return &Mongodb{database: client}
}

func (m *Mongodb) GetCl(name string) *mongo.Collection {
	userCollection := m.database.Database("Shorterism").Collection(name)
	return userCollection
}

// Connect To database Server
//func Connect() {
//	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://nimaism:pass@nimaism.ucoxrje.mongodb.net/?retryWrites=true&w=majority"))
//	if err != nil {
//		log.Fatalln(err)
//	}
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	err = client.Connect(ctx)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	err = client.Ping(ctx, readpref.Primary())
//	if err != nil {
//		log.Fatalln(err)
//	}
//	Data = client
//}
