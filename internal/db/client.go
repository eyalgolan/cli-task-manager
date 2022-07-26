package db

import (
	bolt "go.etcd.io/bbolt"
	"time"
)

type Client struct {
	db *bolt.DB
}

func (cfg Config) Init() (*Client, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return &Client{db}, nil
}

func (c Client) CreateBucket(bucket string) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}
