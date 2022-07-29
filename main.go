package main

import (
	"fmt"
	"os"
	"task/internal/commands"
	"task/internal/db"
)

func main() {
	createDB()
	err := commands.RootCmd.Execute()
	if err != nil {
		fmt.Printf("unable to execture command. Error: %s", err)
		os.Exit(1)
	}
}

func createDB() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	err = db.CreateBucket("tasks")
	if err != nil {
		panic(err)
	}
}
