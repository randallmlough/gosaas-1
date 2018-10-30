package data

import (
	"github.com/dstpierre/gosaas/data/bolt"
	"github.com/dstpierre/gosaas/data/model"
)

func (db *DB) Open(driverName, dataSourceName string) error {
	conn, err := model.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	db.Users = &bolt.Users{}

	db.Connection = conn

	db.DatabaseName = "gosaas"
	return nil
}
