package mock_utils

import "task/internal/db_utils"

type MockHappyFlowClient struct {
}

var MockHappyFlowDB = MockHappyFlowClient{}

func (c *MockHappyFlowClient) CreateTask(task string) error {
	return nil
}

func (c *MockHappyFlowClient) DeleteTask(key int) error {
	return nil
}

func (c *MockHappyFlowClient) AllTasks() ([]db_utils.Task, error) {
	return []db_utils.Task{{
		Key:   0,
		Value: "task",
	}}, nil
}
