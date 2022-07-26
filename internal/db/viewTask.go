package db

import bolt "go.etcd.io/bbolt"

func (c Client) AllTasks() ([]Task, error) {
	var tasks []Task
	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	return tasks, err
}
