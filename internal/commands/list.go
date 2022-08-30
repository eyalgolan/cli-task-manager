package commands

import (
	"fmt"
	"log"
	"task/internal/db"

	"github.com/spf13/cobra"
)

func ListCmd(api db.ViewApi) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists all of your tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := db.AllTasks(api)
			if err != nil {
				log.Fatalf("error getting all tasks: %s", err)
			}
			if len(tasks) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "Nothing to do!")
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), "You have the following tasks:\n")
			for i, task := range tasks {
				fmt.Fprintf(cmd.OutOrStdout(), "%d. %s\n", i+1, task.Value)
			}
		},
	}
}

func init() {
	RootCmd.AddCommand(ListCmd(&db.DBClient))
}
