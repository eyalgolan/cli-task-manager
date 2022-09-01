package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"task/internal/db_utils"
	"task/internal/db_utils/bolt_utils"
)

func DoCmd(doApi db_utils.DeleteApi, viewApi db_utils.ViewApi) *cobra.Command {
	return &cobra.Command{
		Use:   "do",
		Short: "Marks a task as complete",
		Run: func(cmd *cobra.Command, args []string) {
			ids := parseInputIds(args)
			tasks, err := db_utils.AllTasks(viewApi)
			if err != nil {
				log.Fatalf("error getting all tasks: %s", err)
			}
			for _, id := range ids {
				if id <= 0 || id > len(tasks) {
					fmt.Fprintf(cmd.OutOrStdout(), "invalid task number %d\n", id)
					continue
				}
				task := tasks[id-1]
				err = db_utils.DeleteTask(doApi, task.Key)
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
	RootCmd.AddCommand(DoCmd(&bolt_utils.DBClient, &bolt_utils.DBClient))
}
