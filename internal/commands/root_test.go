package commands

import (
	"testing"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := RootCmd
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
