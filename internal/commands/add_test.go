package commands

import (
	"bytes"
	"fmt"
	"github.com/eyalgolan/cli-task-manager/internal/db_utils/mock_utils"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_AddCommandHappyFlow(t *testing.T) {
	cmd := AddCmd(&mock_utils.MockHappyFlowDB)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"task"})
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := "Added task to your task list."
	if !strings.Contains(string(out), expectedOutput) {
		t.Fatalf("expected output to contain \"%s\" got \"%s\"", expectedOutput, string(out))
	}
}

func Test_AddCommandErrorFlow(t *testing.T) {
	cmd := AddCmd(&mock_utils.MockErrorFlowDB)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	taskName := "test"
	cmd.SetArgs([]string{taskName})
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := fmt.Sprintf(
		"Error adding task %s due to %s",
		taskName,
		fmt.Sprintf("unable to create task %s", taskName))
	if string(out) != expectedOutput {
		t.Fatalf("expected output to contain \"%s\" got \"%s\"", expectedOutput, out)
	}
}
