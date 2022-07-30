package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"task/internal/db"
)

func AddCmd(api db.CreateAPI) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Adds a task to your task list",
		Run: func(cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			err := db.CreateTask(api, task)
			if err != nil {
				log.Fatalf("Error adding task %s due to %s", task, err)
			}
			fmt.Printf("Added %s to your task list.\n", task)
		},
	}
}

func init() {
	RootCmd.AddCommand(AddCmd(&db.DBClient))
}
