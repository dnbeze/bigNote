package db

import (
	"context"
	"fmt"
	"github.com/dnbeze/bigNote/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitDB() (*mongo.Database, error) {
	MongoURI, err := utils.GetEnv("MONGO_URI")
	if err != nil {
		log.Fatalf("failed to load MONGOURI")
	}
	MongoDB, err := utils.GetEnv("MONGO_DB")
	if err != nil {
		log.Fatalf("failed to load MONGOURI")
	}
	opts := options.Client().ApplyURI(MongoURI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	fmt.Println("Connected to Database")

	return client.Database(MongoDB), nil
}
