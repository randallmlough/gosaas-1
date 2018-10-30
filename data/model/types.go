package model

import (
	"github.com/boltdb/bolt"
)

type Connection = bolt.DB
type Key = uint64

func Open(options ...string) (*bolt.DB, error) {
	conn, err := bolt.Open(options[1], 0600, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func keyToString(id Key) string {
	return string(id)
}
