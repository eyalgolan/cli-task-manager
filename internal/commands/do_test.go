package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"task/internal/db"
	"testing"
)

func Test_DoCmd(t *testing.T) {
	err := db.CreateTask(&db.MockDB, "task")
	if err != nil {
		t.Fatal(err)
	}
	cmd := DoCmd(&db.MockDB, &db.MockDB)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"1"})
	err = cmd.Execute()
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
