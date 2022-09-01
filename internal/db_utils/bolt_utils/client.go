package bolt_utils

import (
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
	"path/filepath"
	"time"
)

type Client struct {
	DB *bolt.DB
}

var DBClient = Client{}

func Init(c Client, bucket string) error {
	homeDir, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "get the user's home dir")
	}
	dbPath := filepath.Join(homeDir, "tasks.db_utils")

	c.DB, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return errors.Wrap(err, "open db_utils")
	}

	return c.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}
