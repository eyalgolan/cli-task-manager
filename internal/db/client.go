package db

import (
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
	"path/filepath"
	"time"
)

var dbClient *bolt.DB

func Init() error {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(homeDir, "tasks.db")
	dbClient, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return errors.Wrap(err, "open db")
	}
	return nil
}

func CreateBucket(bucket string) error {
	return dbClient.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}

//func (c Client) CloseDB() {
//	err := c.DB.Close()
//	if err != nil {
//		panic(err)
//	}
//}
