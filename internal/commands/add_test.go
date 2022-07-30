package commands

import (
	"task/internal/db"
	"testing"
)

func Test_AddCommand(t *testing.T) {
	err := AddCmd(&db.MockDB)
	if err != nil {
		t.Fatal(err)
	}
}
