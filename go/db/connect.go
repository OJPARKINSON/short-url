package db

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect() (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().
		ApplyURI("mongodb://" + os.Getenv("CONNECTION_URL") + "/?compressors=snappy,zlib,zstd"))
	if err != nil {
		fmt.Println("Error decoding body: ", err)
		return nil, err
	}

	return client, nil
}

func ConnectToCollection() (*mongo.Collection, error) {
	client, err := Connect()
	if err != nil {
		fmt.Println("Error decoding body: ", err)
		return nil, err
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	return collection, nil
}
