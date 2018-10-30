package data

import (
	"github.com/dstpierre/gosaas/data/model"
)

type DB struct {
	DatabaseName string
	Connection   *model.Connection

	Users UserServices
}

type UserServices interface {
	Purchase(email, product, source, confirmID string) (*model.Purchase, error)
	GetSignInLink(email string) (string, error)
	GetDownload(id model.Key) (*model.Purchase, error)
	Auth(accountID, token string, pat bool) (*model.Account, *model.User, error)
}
