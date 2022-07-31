package commands

import (
	"bytes"
	"io/ioutil"
	"task/internal/db"
	"testing"
)

func Test_ListCommand(t *testing.T) {
	cmd := AddCmd(&db.MockDB)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hi" {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi", string(out))
	}
}
