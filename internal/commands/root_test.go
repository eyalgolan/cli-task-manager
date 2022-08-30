package commands

import "testing"

func Test_RootCommand(t *testing.T) {
	err := RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
