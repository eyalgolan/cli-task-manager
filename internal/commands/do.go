package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"task/internal/db"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		ids := parseInputIds(args)
		tasks, err := db.AllTasks(&db.DBClient)
		if err != nil {
			log.Fatalf("error getting all tasks: %s", err)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("invalid task number %d\n", id)
				continue
			}
			task := tasks[id-1]
			err = db.DeleteTask(&db.DBClient, task.Key)
			if err != nil {
				fmt.Printf("failed to mark %d as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked %d as completed\n", id)
			}
		}
	},
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
	RootCmd.AddCommand(doCmd)
}
