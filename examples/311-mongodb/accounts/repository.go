package accounts

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "accounts"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{collection: db.Collection(collectionName)}
}

func (r *Repository) InsertOne(ctx context.Context, document interface{}) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, document)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), err
}

// @TODO verify if some of docs can be inserted - and error will occur - what should be returned
func (r *Repository) InsertMany(ctx context.Context, documents []interface{}) ([]primitive.ObjectID, error) {
	var insertedIDs []primitive.ObjectID

	result, err := r.collection.InsertMany(ctx, documents)

	if err != nil {
		insertedIDs = append(insertedIDs, primitive.NilObjectID)

		return insertedIDs, err
	}

	for _, insertedID := range result.InsertedIDs {
		switch id := insertedID.(type) {
		case primitive.ObjectID:
			insertedIDs = append(insertedIDs, id)
		default:
			fmt.Errorf("inserted ID is of an unknown type")
		}
	}

	return insertedIDs, err
}

func (r *Repository) FindOneByAccountTypeAndBalanceGte(ctx context.Context, accountType string, balance int) (Account, error) {
	//filter := bson.D{
	//	{Key: "balance", Value: bson.D{{Key: "$gte", Value: 1000}}},
	//	{Key: "account_type", Value: "checking"},
	//}
	filter := bson.M{
		"balance":      bson.M{"$gte": balance},
		"account_type": accountType,
	}

	return r.FindOneWithFilter(ctx, filter)
}

func (r *Repository) FindOneWithFilter(ctx context.Context, filter interface{}) (Account, error) {

	var result Account
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Account{}, fmt.Errorf("no document found: %w", err)
		} else {
			return Account{}, err
		}
	}

	isDebug := true //ctx.Value("isDebug").(bool)

	if isDebug {
		docJSON, err := bson.MarshalExtJSON(result, true, true)
		if err == nil {
			fmt.Println(string(docJSON))
		}
	}

	return result, nil
}

func (r *Repository) Find(ctx context.Context, filter interface{}) (Accounts, error) {
	//findOpts := &options.FindOptions{}
	var accounts Accounts

	if filter == nil { //default filter if not set, it's needed
		filter = bson.M{}
	}

	cursor, err := r.collection.Find(ctx, filter) //, findOpts)
	if err != nil {
		return accounts, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var account Account
		err = cursor.Decode(&account)
		if err != nil {
			return accounts, err
		}
		accounts = append(accounts, account)
	}

	return accounts, cursor.Err()
}
