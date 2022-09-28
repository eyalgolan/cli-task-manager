package commands

import (
	"bytes"
	"cli_task_manager/internal/db_utils"
	"cli_task_manager/internal/db_utils/mock_utils"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_ListCommand(t *testing.T) {
	err := db_utils.CreateTask(&mock_utils.MockHappyFlowDB, "task")
	if err != nil {
		t.Fatal(err)
	}
	cmd := ListCmd(&mock_utils.MockHappyFlowDB)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	err = cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := fmt.Sprintf("%d. task", 1)
	if !strings.Contains(string(out), expectedOutput) {
		t.Fatalf("expected output to contain \"%s\" got \"%s\"", expectedOutput, string(out))
	}
}
