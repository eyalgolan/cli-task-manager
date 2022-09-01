package mock_utils

import (
	"fmt"
	"github.com/pkg/errors"
)

type MockErrorFlowClient struct {
}

var MockErrorFlowDB = MockErrorFlowClient{}

func (c *MockErrorFlowClient) CreateTask(task string) error {
	return errors.New(fmt.Sprintf("unable to create task %s", task))
}

func (c *MockErrorFlowClient) DeleteTask(key int) error {
	return nil
}
