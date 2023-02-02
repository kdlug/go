package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://localhost:27017"
const dbname = "blog"

var ctx = context.TODO()
var db *mongo.Database
var collection *mongo.Collection

func Init() {
	db, _ = GetDatabase()
}

func GetDatabase() (*mongo.Database, error) {
	var err error
	var database *mongo.Database

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected and pinged.")

	database = client.Database(dbname)

	return database, err
}
