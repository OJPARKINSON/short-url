package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Init() {
	collection, err := ConnectToCollection()
	if err != nil {
		fmt.Println("failed to connect to db for init script")
	}

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "shortCode", Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	result, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		fmt.Println("failed to create index in init script: ", err)
	}

	fmt.Println("creating index", result)

}
