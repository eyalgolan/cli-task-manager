package main

import (
	"log"
	"task/internal/commands"
	"task/internal/db"
)

func main() {
	createDB()
	err := commands.RootCmd.Execute()
	if err != nil {
		log.Fatalf("unable to execture command. Error: %s", err)
	}
}

func createDB() {
	err := db.Init(&db.DBClient)
	if err != nil {
		log.Fatalf("unable to init db. Error: %s", err)
	}
	err = db.CreateBucket(&db.DBClient, "tasks")
	if err != nil {
		log.Fatalf("unable to create db buckets. Error: %s", err)
	}
}
