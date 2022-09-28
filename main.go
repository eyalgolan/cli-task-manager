package main

import (
	"log"
	"task/internal/db_utils/bolt_utils"
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
