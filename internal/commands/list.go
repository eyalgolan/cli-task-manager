package commands

import (
	"fmt"
	"os"
	"task/internal/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		dbClient, err := db.Init()
		if err != nil {
			fmt.Printf("error in db init: %s", err.Error())
			os.Exit(1)
		}
		tasks, err := dbClient.AllTasks()
		if err != nil {
			fmt.Printf("error getting all tasks: %s", err.Error())
			os.Exit(1)
		}
		defer dbClient.CloseDB()
		if len(tasks) == 0 {
			fmt.Println("Nothing to do!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
