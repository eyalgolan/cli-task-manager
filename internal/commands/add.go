package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"task/internal/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		dbClient, err := db.Init()
		if err != nil {
			fmt.Printf("error in db init: %s", err)
			os.Exit(1)
		}
		defer dbClient.CloseDB()
		_, err = dbClient.CreateTask(task)
		if err != nil {
			fmt.Printf("Error adding task %s due to %s", task, err)
			os.Exit(1)
		}
		fmt.Printf("Added %s to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
