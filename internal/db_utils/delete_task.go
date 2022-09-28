package db_utils

type DeleteApi interface {
	DeleteTask(key int) error
}

func DeleteTask(api DeleteApi, key int) error {
	return api.DeleteTask(key)
}
