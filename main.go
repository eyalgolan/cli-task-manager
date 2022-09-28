package main

import (
	"github.com/eyalgolan/cli-task-manager/internal/db_utils/bolt_utils"
	"log"
)

func main() {
	err := bolt_utils.Init(bolt_utils.DBClient, "tasks")
	if err != nil {
		log.Fatalf("unable to init db_utils. Error: %s", err)
	}
	//err = commands.RootCmd.Execute()
	//if err != nil {
	//	log.Fatalf("unable to execture command. Error: %s", err)
	//}
}
