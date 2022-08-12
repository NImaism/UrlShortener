package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Data *mongo.Client

func Connect() {
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

	Data = client
}

// Connect To Database Server
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

func GetCl(db *mongo.Client, name string) *mongo.Collection {
	userCollection := db.Database("Shorterism").Collection(name)
	return userCollection
}
