package db

import (
	bolt "go.etcd.io/bbolt"
	"time"
)

func (cfg DBConfig) Init() (*bolt.DB, error) {
	return bolt.Open(cfg.DBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
}

func (cfg DBConfig) CreateBucket(db *bolt.DB, bucket string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}
