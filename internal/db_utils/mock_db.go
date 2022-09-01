package db_utils

type MockClient struct {
}

var MockDB = MockClient{}

func (c *MockClient) CreateTask(task string) error {
	return nil
}

func (c *MockClient) DeleteTask(key int) error {
	return nil
}

func (c *MockClient) AllTasks() ([]Task, error) {
	return []Task{{
		Key:   0,
		Value: "task",
	}}, nil
}
