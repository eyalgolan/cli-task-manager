package mock_utils

import "github.com/eyalgolan/cli-task-manager/internal/db_utils"

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
