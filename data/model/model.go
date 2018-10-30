package model

import (
	"time"
)

type Roles int

const (
	RoleAdmin Roles = iota
	RoleUser
)

type Purchase struct {
	ID             Key       `json:"id"`
	Email          string    `json:"email"`
	PaymentSource  string    `json:"paymentSource"`
	ConfirmationID string    `json:"confirmationId"`
	Product        string    `json:"product"`
	PurchasedOn    time.Time `json:"purchasedOn"`
}

// To stay compatible with the engine
type Account struct {
	ID Key
}
type User struct {
	ID    Key
	Email string
}

// APIRequest represents a single API call
type APIRequest struct {
	ID         Key       `bson:"_id" json:"id"`
	AccountID  Key       `bson:"accountId" json:"accountId"`
	UserID     Key       `bson:"userId" json:"userId"`
	URL        string    `bson:"url" json:"url"`
	Requested  time.Time `bson:"reqon" json:"requested"`
	StatusCode int       `bson:"sc" json:"statusCode"`
	RequestID  string    `bson:"reqid" json:"reqId"`
}
