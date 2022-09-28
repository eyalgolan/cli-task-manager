package commands

import (
	"bytes"
	"cli_task_manager/internal/db_utils/mock_utils"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_DoCmdHappyFlow(t *testing.T) {
	cmd := DoCmd(&mock_utils.MockHappyFlowDB)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"1"})
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedOutput := fmt.Sprintf("Marked %d as completed", 1)
	if !strings.Contains(string(out), expectedOutput) {
		t.Fatalf("expected output to contain \"%s\" got \"%s\"", expectedOutput, string(out))
	}
}
