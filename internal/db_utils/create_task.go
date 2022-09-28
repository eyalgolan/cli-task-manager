package db_utils

type CreateAPI interface {
	CreateTask(task string) error
}

func CreateTask(api CreateAPI, task string) error {
	return api.CreateTask(task)
}
