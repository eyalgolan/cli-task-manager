package db

import bolt "go.etcd.io/bbolt"

func DeleteTask(key int) error {
	return dbClient.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		return b.Delete(itob(key))
	})
}
