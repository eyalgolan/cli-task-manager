package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"task/internal/db"
	"testing"
)

func Test_ListCommand(t *testing.T) {
	err := db.CreateTask(&db.MockDB, "task")
	if err != nil {
		t.Fatal(err)
	}
	cmd := ListCmd(&db.MockDB)
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
