package db_utils

type ViewApi interface {
	AllTasks() ([]Task, error)
}

func AllTasks(api ViewApi) ([]Task, error) {
	return api.AllTasks()
}
