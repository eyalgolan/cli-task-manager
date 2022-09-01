package main

import (
	"log"
	"task/internal/commands"
	"task/internal/db_utils/bolt_utils"
)

func main() {
	err := bolt_utils.Init(&bolt_utils.DBClient)
	if err != nil {
		log.Fatalf("unable to init db_utils. Error: %s", err)
	}
	err = bolt_utils.CreateBucket(&bolt_utils.DBClient, "tasks")
	if err != nil {
		log.Fatalf("unable to create db_utils buckets. Error: %s", err)
	}
	err = commands.RootCmd.Execute()
	if err != nil {
		log.Fatalf("unable to execture command. Error: %s", err)
	}
}
