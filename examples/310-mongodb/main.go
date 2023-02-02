package main

// mongo quick guide:
// docker pull mongo
// docker run -d --name mongo -p 27017:27017 mongo
// docker exec -it mongo mongosh
// use test
// show collections
// db.task.find()
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://localhost:27017"

var collection *mongo.Collection

// mongo.Connect() accepts a Context and a options.ClientOptions object, which is used to set the connection string and other driver settings. You can visit the options package documentation to see what configuration options are available.

// Context is like a timeout or deadline that indicates when an operation should stop running and return.
// It helps to prevent performance degradation on production systems when specific operations are running slow.
// In this code, you’re passing context.TODO() to indicate that you’re not sure what context to use right now, but you plan to add one in the future.
var ctx = context.TODO()

// You use the primitive package to set the type of the ID of each task
// since MongoDB uses ObjectIDs for the _id field by default.
// Another default behavior of MongoDB is that the lowercase field name is used as the key for each exported
// field when it is being serialized, but this can be changed using bson struct tags.

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

func main() {
	// Create a new client and connect to the server
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	// Ensure that your MongoDB server was found and connected to successfully using the Ping method.
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected and pinged.")

	// Create a database
	collection = client.Database("test").Collection("task")

	task := &Task{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Text:      "Task #1",
		Completed: false,
	}

	newTasks := []interface{}{
		&Task{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Text:      "Task #10",
			Completed: false,
		},
		&Task{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Text:      "Task #11",
			Completed: false,
		},
	}

	if err := createTask(task); err != nil {
		log.Fatal(err)
	}
	if err := createTasks(newTasks); err != nil {
		log.Fatal(err)
	}

	tasks, _ := getAll()
	printTasks(tasks)
}

func printTasks(tasks []*Task) {
	for i, v := range tasks {
		fmt.Printf("%d: %s\n", i+1, v.Text)

	}
}

func createTask(task *Task) error {
	_, err := collection.InsertOne(ctx, task) // returns the ID of the document that was inserted.
	return err
}

func createTasks(tasks []interface{}) error {
	_, err := collection.InsertMany(ctx, tasks) // returns the IDs of the documents that was inserted.
	return err
}

func getAll() ([]*Task, error) {
	// passing bson.D{{}} matches all documents in the collection
	// BSON (Binary-encoded JSON) is how documents are represented in a MongoDB database
	// and the bson package is what helps us work with BSON objects in Go.
	filter := bson.D{{}}
	return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*Task, error) {
	// A slice of tasks for storing the decoded documents
	var tasks []*Task

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t Task
		err := cur.Decode(&t)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}
