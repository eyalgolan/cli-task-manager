package commands

import (
	"cli_task_manager/internal/db_utils"
	"cli_task_manager/internal/db_utils/bolt_utils"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

type doCmdApis interface {
	db_utils.DeleteApi
	db_utils.ViewApi
}

func DoCmd(apis doCmdApis) *cobra.Command {
	return &cobra.Command{
		Use:   "do",
		Short: "Marks a task as complete",
		Run: func(cmd *cobra.Command, args []string) {
			ids := parseInputIds(args)
			tasks, err := db_utils.AllTasks(apis)
			if err != nil {
				log.Fatalf("error getting all tasks: %s", err)
			}
			for _, id := range ids {
				if id <= 0 || id > len(tasks) {
					fmt.Fprintf(cmd.OutOrStdout(), "invalid task number %d\n", id)
					continue
				}
				task := tasks[id-1]
				err = db_utils.DeleteTask(apis, task.Key)
				if err != nil {
					fmt.Printf("failed to mark %d as completed. Error: %s\n", id, err)
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "Marked %d as completed\n", id)
				}
			}
		},
	}
}

func parseInputIds(args []string) []int {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Failed to parse the argument %s", arg)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}
func init() {
	RootCmd.AddCommand(DoCmd(&bolt_utils.DBClient))
}
