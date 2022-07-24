package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse the argument %s", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Printf("Following tasks marked as done:\n")
		for _, id := range ids {
			fmt.Printf("%d\n", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
