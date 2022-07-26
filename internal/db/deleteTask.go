package db

import bolt "go.etcd.io/bbolt"

func (c Client) DeleteTask(key int) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.Delete(itob(key))
	})
}
