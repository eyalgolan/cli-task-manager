package db

import bolt "go.etcd.io/bbolt"

func (c Client) AllTasks() ([]Task, error) {
	var tasks []Task
	err := c.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	return tasks, err
}
