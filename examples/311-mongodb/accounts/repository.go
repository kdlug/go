package accounts

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	collectionName = "accounts"
)

var ErrAcountNotFound = fmt.Errorf("account not found")

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

func (r *Repository) FindOneByAccountID(ctx context.Context, accountID string) (Account, error) {
	filter := bson.M{
		"account_id": accountID,
	}

	return r.FindOneWithFilter(ctx, filter)
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

func (r *Repository) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}

func (r *Repository) UpdateMany(ctx context.Context, filter interface{}, update interface{}) error {

	result, err := r.collection.UpdateMany(ctx, filter, update)

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no documents updated")
	}

	if result.ModifiedCount < result.MatchedCount {
		return fmt.Errorf("partial update %d of %d documents updated", result.ModifiedCount, result.MatchedCount)
	}

	return err
}

func (r *Repository) AddMinBalanceToSavingsAccounts(ctx context.Context, balance float64) error {
	filter := bson.M{
		"account_type": AccountTypeSavings,
	}

	update := bson.M{
		"$set": bson.M{
			"balance": balance,
		},
	}

	return r.UpdateMany(ctx, filter, update)
}

// FindOneAndUpdate finds document, update and returns updated document
func (r *Repository) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) (Account, error) {
	opts := options.FindOneAndUpdateOptions{}
	opts.SetReturnDocument(options.After)

	result := r.collection.FindOneAndUpdate(ctx, filter, update, &opts)

	switch result.Err() {
	case nil:
	case mongo.ErrNoDocuments:
		return Account{}, ErrAcountNotFound
	default:
		return Account{}, fmt.Errorf("action %s failed:  %w", "FindOneAndUpdate", result.Err().Error())
	}

	var account Account
	err := result.Decode(&account)

	return account, err
}

func (r *Repository) IncreaseBalanceByID(ctx context.Context, ID primitive.ObjectID, balance float64) (Account, error) {
	//objectId, err := primitive.ObjectIDFromHex(accountID)
	//if err != nil {
	//	return Account{}, fmt.Errorf("invalid account ID: %w", err)
	//}
	filter := bson.M{"_id": ID}
	update := bson.M{"$set": bson.M{
		"status":       AccountStatusActive,
		"last_updated": time.Now().UTC(),
	},
		"$inc": bson.M{
			"balance": balance,
		},
	}

	return r.FindOneAndUpdate(ctx, filter, update)
}

func (r *Repository) DeleteOneByID(ctx context.Context, ID primitive.ObjectID) error {
	filter := bson.M{"_id": ID}
	result, err := r.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("no documents deleted")
	}

	return nil
}

func (r *Repository) ShowAccountTypeSummary(ctx context.Context, maxBalance float64) (AccountSummary, error) {

	pipeline := mongo.Pipeline{
		{{
			"$match", bson.M{
				"balance": bson.M{"$lte": maxBalance},
			},
		}},
		{{
			"$group", bson.M{
				"_id":         "$account_type",
				"avg_balance": bson.M{"$avg": "$balance"},
			},
		}},
		// For sort stage we use bson.D, to keep order
		{{
			"$sort", bson.D{{"avg_balance", -1}},
		}},
		//{{
		//	"$limit", 5,
		//}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return AccountSummary{}, err
	}

	var accountSummary AccountSummary
	for cursor.Next(ctx) {
		var result bson.M
		err = cursor.Decode(&result)
		if err != nil {
			return AccountSummary{}, err
		}
		accountSummary.AccountType = result["_id"].(string)
		accountSummary.AvgBalance = result["avg_balance"].(float64)
	}

	return accountSummary, cursor.Err()
}
