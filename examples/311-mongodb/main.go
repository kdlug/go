package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"mongodb-example/accounts"
	"mongodb-example/mongodb"
	"strings"
	"time"
)

type contextKey string

const debugKey contextKey = "isDebug"

func main() {
	ctx := context.WithValue(context.Background(), debugKey, true)

	// Connect to MongoDB
	db, err := mongodb.GetDatabase()
	if err != nil {
		panic("initialize mongodb client")
	}

	// Account doc
	account := accounts.Account{
		AccountHolder: "John Doe",
		AccountID:     "MDB310054629",
		Balance:       1785.50,
		AccountType:   accounts.AccountTypeSavings,
		LastUpdated:   time.Date(2024, time.March, 20, 14, 45, 0, 0, time.UTC),
	}

	accountsRepo := accounts.NewRepository(db)

	// Insert one
	fmt.Println("Insert one", strings.Repeat("-", 20))
	id, err := accountsRepo.InsertOne(ctx, account)
	if err != nil {
		fmt.Println("Error inserting one:", err)
	} else {
		fmt.Println("Inserted one:", id.String())
	}

	// Insert many
	accountList := []accounts.Account{account, account}

	// Convert []Account na []interface{}
	docs := make([]interface{}, len(accountList))
	for i, a := range accountList {
		docs[i] = a
	}

	fmt.Println("Insert many", strings.Repeat("-", 20))
	accountsRepo.InsertMany(ctx, docs)

	// Find one
	fmt.Println("Find one", strings.Repeat("-", 20))
	result, err := accountsRepo.FindOneByAccountTypeAndBalanceGte(ctx, accounts.AccountTypeChecking, 1000)
	if err != nil {
		fmt.Println("Error finding one:", err)
	} else {
		fmt.Println("Found one:", result)
	}

	// Find many
	fmt.Println("Find many", strings.Repeat("-", 20))
	results, err := accountsRepo.Find(ctx, nil)
	if err != nil {
		fmt.Println("Error finding all:", err)
	} else {
		fmt.Println("Found all:", results)
	}
	fmt.Println("Found all:", results)

	fmt.Println("Find by Account ID", strings.Repeat("-", 20))
	acc, err := accountsRepo.FindOneByAccountID(ctx, "MDB99115881")
	if err != nil {
		fmt.Println("Error finding one:", err)
	} else {
		fmt.Println("Found one:", acc)
	}
	fmt.Println("Increase balance by account ID", strings.Repeat("-", 20))
	update, err := accountsRepo.IncreaseBalanceByID(ctx, acc.ID, 100)
	if err != nil {
		fmt.Println("Error updating one:", err)
	} else {
		fmt.Println("Updated one:", update)
	}

	fmt.Println("Update many", strings.Repeat("-", 20))
	err = accountsRepo.AddMinBalanceToSavingsAccounts(ctx, 100)
	if err != nil {
		fmt.Println("Error updating many:", err)
	} else {
		fmt.Println("Updated many")
	}

	err = accountsRepo.DeleteOneByID(ctx, acc.ID)
	if err != nil {
		fmt.Println("Error deleting one:", err)
	} else {
		fmt.Println("Deleted one")
	}

	fmt.Println("Transaction", strings.Repeat("-", 20))
	transactionTransferFunds(err, db, accountsRepo, ctx)

	summary, err := accountsRepo.ShowAccountTypeSummary(ctx, 1000)

	fmt.Println("Summary", strings.Repeat("-", 20))
	fmt.Println(summary)

}
func transactionTransferFunds(err error, db *mongo.Database, accountsRepo *accounts.Repository, ctx context.Context) {
	clientSession, err := db.Client().StartSession()
	if err != nil {
		panic(err)
	}
	defer clientSession.EndSession(context.Background())

	// Transaction
	txnBody := func(sessCtx mongo.SessionContext) (interface{}, error) {
		fromAccount, err := accountsRepo.FindOneByAccountID(ctx, "MDB99115881")
		if err != nil {
			return nil, err
		}

		from := bson.M{
			"_id": fromAccount.ID,
		}
		withdrawal := bson.M{
			"$inc": bson.M{
				"balance": -100,
			},
		}

		toAccount, err := accountsRepo.FindOneByAccountID(ctx, "MDB310054629")

		if err != nil {
			return nil, err
		}
		to := bson.M{
			"_id": toAccount.ID,
		}
		deposit := bson.M{
			"$inc": bson.M{
				"balance": 100,
			},
		}

		if fromAccount.Balance < 100 {
			return nil, fmt.Errorf("insufficient funds")
		}

		err = accountsRepo.UpdateOne(ctx, from, withdrawal)
		if err != nil {
			return nil, err
		}

		err = accountsRepo.UpdateOne(ctx, to, deposit)
		if err != nil {
			return nil, err
		}

		return "Transferred funds", nil
	}

	transactionResult, err := clientSession.WithTransaction(context.Background(), txnBody)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(transactionResult)
		toAccount, _ := accountsRepo.FindOneByAccountID(ctx, "MDB310054629")
		fromAccount, _ := accountsRepo.FindOneByAccountID(ctx, "MDB99115881")

		fmt.Println(fromAccount)
		fmt.Println(toAccount)
	}
}
