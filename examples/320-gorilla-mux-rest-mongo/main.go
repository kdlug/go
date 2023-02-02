package main

// mongo quick guide:
// docker pull mongo
// docker run -d --name mongo -p 27017:27017 mongo
// docker exec -it mongo mongosh
// use test
// show collections
// db.task.find()
import (
	model "github.com/kdlug/restapi/model"

	"time"

	controller "github.com/kdlug/restapi/controller"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {

	articles := []interface{}{
		&model.Article{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     "Article #1",
			Content:   "Content",
		},
		&model.Article{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     "Article #2",
			Content:   "Content",
		},
	}

	model.Init()

	model.CreateArticles(articles)
}

func main() {

	controller.HandleRequests()
}
