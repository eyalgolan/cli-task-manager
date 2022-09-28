package bolt_utils

import (
	"cli_task_manager/internal/db_utils"
	bolt "go.etcd.io/bbolt"
)

func (c *Client) AllTasks() ([]db_utils.Task, error) {
	var tasks []db_utils.Task
	err := c.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			tasks = append(tasks, db_utils.Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	return tasks, err
}
