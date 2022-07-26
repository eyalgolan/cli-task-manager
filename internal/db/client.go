package db

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

func Init() (*Client, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(homeDir, "tasks.db")
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, errors.Wrap(err, "open db")
	}
	return &Client{db}, nil
}

func (c Client) CreateBucket(bucket string) error {
	return c.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}

func (c Client) CloseDB() {
	err := c.DB.Close()
	if err != nil {
		panic(err)
	}
}
