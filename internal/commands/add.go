package commands

import (
	"cli_task_manager/internal/db_utils"
	"cli_task_manager/internal/db_utils/bolt_utils"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func AddCmd(api db_utils.CreateAPI) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Adds a task to your task list",
		Run: func(cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			err := db_utils.CreateTask(api, task)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), "Error adding task %s due to %s", task, err)
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Added %s to your task list.\n", task)
		},
	}
}

func init() {
	RootCmd.AddCommand(AddCmd(&bolt_utils.DBClient))
}
