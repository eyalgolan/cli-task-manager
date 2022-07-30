package db

type MockClient struct {
}

var MockDB = MockClient{}

func (c *MockClient) CreateTask(task string) error {
	return nil
}
