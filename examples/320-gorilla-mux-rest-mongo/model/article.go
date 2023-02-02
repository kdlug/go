package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
}

// type Repository struct {
//var collection *mongo.Collection

// }

func GetCollection() *mongo.Collection {
	return db.Collection("article")

}

func GetAllPosts() ([]*Article, error) {
	filter := bson.D{{}}

	return filterPosts(filter)
}

func filterPosts(filter interface{}) ([]*Article, error) {
	collection := GetCollection()

	var articles []*Article

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return articles, err
	}

	for cur.Next(ctx) {
		var t Article
		err := cur.Decode(&t)
		if err != nil {
			return articles, err
		}

		articles = append(articles, &t)
	}

	if err := cur.Err(); err != nil {
		return articles, err
	}

	cur.Close(ctx)

	if len(articles) == 0 {
		return articles, mongo.ErrNoDocuments
	}

	return articles, nil
}

func CreateArticles(articles []interface{}) error {
	collection := GetCollection()

	_, err := collection.InsertMany(ctx, articles)
	return err
}
