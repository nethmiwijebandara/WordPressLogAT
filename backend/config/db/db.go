package db

import (
	"backend/model"
	"context"
	"fmt"
	"log"
	"time"

	"../model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDBCollectionUnzip(fileContentList string) (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://neth98:neth123@test.dq1l7.mongodb.net/FilesDB?retryWrites=true&w=majority"))
	checkError(err)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}
	collectionUnzip := client.Database("FilesDB").Collection(fileContentList)

	return collectionUnzip, nil

}

func GetDBCollectionLogin(loginList string) (*mongo.Collection, error) {
	// log.Fatal("err")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://neth98:neth123@test.dq1l7.mongodb.net/GoLogin?retryWrites=true&w=majority"))
	checkError(err)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}
	collectionLogin := client.Database("GoLogin").Collection("users")
	// log.Fatal(collectionLogin)
	return collectionLogin, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateAccessLog(accessLog model.AccessLog) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://neth98:neth123@test.dq1l7.mongodb.net/FilesDB?retryWrites=true&w=majority"))

	checkError(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	accesslogCollection := client.Database("Logs").Collection("actionlogs")

	result, insertErr := accesslogCollection.InsertOne(ctx, accessLog)

	if insertErr != nil {
		log.Fatal(insertErr)
	}
	newID := result.InsertedID

	fmt.Print("Inserted ID ", newID)
}


func CreateErrorLog(errorLog model.ErrorLog) {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://neth98:neth123@test.dq1l7.mongodb.net/FilesDB?retryWrites=true&w=majority"))

	checkError(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	accesslogCollection := client.Database("Logs").Collection("errorlogs")

	result, insertErr := accesslogCollection.InsertOne(ctx, errorLog)

	if insertErr != nil {
		log.Fatal(insertErr)
	}
	newID := result.InsertedID

	fmt.Print("Inserted ID ", newID)
}

