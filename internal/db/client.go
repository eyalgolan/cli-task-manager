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

var DBClient = Client{}

type InitAPI interface {
	Init() error
	CreateBucket(bucket string) error
}

func Init(api InitAPI) error {
	return api.Init()
}

func CreateBucket(api InitAPI, bucket string) error {
	return api.CreateBucket(bucket)
}

func (c *Client) Init() error {
	homeDir, err := homedir.Dir()
	if err != nil {
		return errors.Wrap(err, "get the user's home dir")
	}
	dbPath := filepath.Join(homeDir, "tasks.db")

	c.DB, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return errors.Wrap(err, "open db")
	}
	return nil
}

func (c *Client) CreateBucket(bucket string) error {
	return c.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})
}
