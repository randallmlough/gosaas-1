package bolt

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/dstpierre/gosaas/data/model"
)

type Users struct {
	DB *bolt.DB
}

const (
	bucketPurchases string = "purchases"
	bucketEmails    string = "emails"
)

func (u *Users) Purchase(email, product, source, confirmID string) (*model.Purchase, error) {
	data := model.Purchase{
		Email:          email,
		PaymentSource:  source,
		ConfirmationID: confirmID,
		Product:        product,
		PurchasedOn:    time.Now(),
	}

	err := u.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketPurchases))

		id, _ := b.NextSequence()
		data.ID = id

		buf, err := json.Marshal(data)
		if err != nil {
			return err
		}

		if err := b.Put(itob(id), buf); err != nil {
			return err
		}

		be := tx.Bucket([]byte(bucketEmails))
		if err != nil {
			return err
		}

		return be.Put([]byte(email), []byte(email))
	})

	return &data, err
}

func (u *Users) GetSignInLink(email string) (string, error) {
	link := ""
	err := u.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketEmails))

		exists := b.Get([]byte(email))
		if exists == nil {
			return fmt.Errorf("email does not exists")
		}

		link = email
		return nil
	})
	return link, err
}

func (u *Users) GetDownload(id model.Key) (*model.Purchase, error) {
	var data model.Purchase
	err := u.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketPurchases))

		buf := b.Get(itob(id))
		if buf == nil {
			return fmt.Errorf("unable to find this order")
		}

		return json.Unmarshal(buf, &data)
	})
	return &data, err
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// Auth just a stub to remains compatible with the Auth engine
func (u *Users) Auth(accountID, token string, pat bool) (*model.Account, *model.User, error) {
	return &model.Account{}, &model.User{}, nil
}
