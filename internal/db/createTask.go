package db

import (
	bolt "go.etcd.io/bbolt"
)

func (c DBClient) createTaskFunc(task string) (*int, error) {
	var id int
	err := c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		// not catching the error because when inside of an Update transaction, the only way an error can only occur is
		// if the transaction is closed or if it's not writable transaction. Both shouldn't be possible here.
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	return &id, err
}
