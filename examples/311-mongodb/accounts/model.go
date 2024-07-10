package accounts

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Accounts []Account

// Account represents the structure of a MongoDB document
type Account struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	AccountHolder string             `bson:"account_holder"`
	AccountID     string             `bson:"account_id"`
	Balance       float64            `bson:"balance"`
	AccountType   string             `bson:"account_type"`
	LastUpdated   time.Time          `bson:"last_updated"`
}

const AccountTypeChecking = "checking"
const AccountTypeSavings = "savings"
