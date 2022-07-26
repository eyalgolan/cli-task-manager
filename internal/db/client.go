package db

import (
	bolt "go.etcd.io/bbolt"
	"time"
)

type DBClient struct {
	db *bolt.DB
}

func (cfg DBConfig) Init() (*DBClient, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return &DBClient{db}, nil
}

func (c DBClient) CreateBucket(bucket string) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}
