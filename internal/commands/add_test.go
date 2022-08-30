package commands

import (
	"bytes"
	"io/ioutil"
	"strings"
	"task/internal/db"
	"testing"
)

func Test_AddCommand(t *testing.T) {
	cmd := AddCmd(&db.MockDB)
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
