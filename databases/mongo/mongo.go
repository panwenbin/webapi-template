package mongo

import (
	"app/settings"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var Mongo *mongo.Client

var (
	MongodbUri string
)

func init() {
	settings.RequireEnvs([]string{
		"MONGODB_URI",
	})

	MongodbUri = os.Getenv("MONGODB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(MongodbUri))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}
	Mongo = client
}
