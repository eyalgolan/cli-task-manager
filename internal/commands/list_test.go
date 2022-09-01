package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"task/internal/db_utils"
	"testing"
)

func Test_ListCommand(t *testing.T) {
	err := db_utils.CreateTask(&db_utils.MockDB, "task")
	if err != nil {
		t.Fatal(err)
	}
	cmd := ListCmd(&db_utils.MockDB)
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
