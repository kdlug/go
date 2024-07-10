package main

import (
	"context"
	"fmt"
	"mongodb-example/accounts"
	"mongodb-example/mongodb"
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

	account := accounts.Account{
		AccountHolder: "John Doe",
		AccountID:     "MDB99115881",
		Balance:       1785.50,
		AccountType:   "checking",
		LastUpdated:   time.Date(2024, time.March, 20, 14, 45, 0, 0, time.UTC),
	}

	accountsRepo := accounts.NewRepository(db)

	accountsRepo.InsertOne(ctx, account)

	accountList := []accounts.Account{account, account}

	// Convert []Account na []interface{}
	docs := make([]interface{}, len(accountList))
	for i, a := range accountList {
		docs[i] = a
	}

	accountsRepo.InsertMany(ctx, docs)

	result, err := accountsRepo.FindOneByAccountTypeAndBalanceGte(ctx, accounts.AccountTypeChecking, 1000)
	if err != nil {
		fmt.Println("Error finding one:", err)
	} else {
		fmt.Println("Found one:", result)
	}

	results, err := accountsRepo.Find(ctx, nil)
	if err != nil {
		fmt.Println("Error finding all:", err)
	} else {
		fmt.Println("Found all:", results)
	}
	fmt.Println("Found all:", results)

}
