package db

import bolt "go.etcd.io/bbolt"

type DeleteApi interface {
	DeleteTask(key int) error
}

func DeleteTask(api DeleteApi, key int) error {
	return api.DeleteTask(key)
}

func (c *Client) DeleteTask(key int) error {
	return c.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.Delete(itob(key))
	})
}
