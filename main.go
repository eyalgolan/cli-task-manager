package main

import (
	"task/internal/commands"
	"task/internal/db"
)

func main() {
	createDB()
	err := commands.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func createDB() {
	dbClient, err := db.Init()
	if err != nil {
		panic(err)
	}
	err = dbClient.CreateBucket("tasks")
	if err != nil {
		panic(err)
	}
	defer dbClient.CloseDB()
}
