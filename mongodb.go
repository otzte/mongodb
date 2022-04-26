package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var collection *mongo.Collection
var ctx, cancel = context.WithTimeout(context.Background(),30*time.Second)

func Init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}	
	err = client.Ping(ctx, nil)
	if err != nil {
	  log.Fatal(err)
	}
	collection = client.Database("Adorithm").Collection("users")
}

func insertOne() string {
	var document interface{}
	document = bson.D{
        {"rollNo", 175},
        {"maths", 80},
        {"science", 90},
        {"computer", 95},
    }
	collection.InsertOne(ctx,document)
	return "success"
}

func main(){
	Init()
	insertOne()
	// releases resources if connection completes before timeout elapses
	defer cancel()
}