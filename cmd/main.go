package main

import (
	"github.com/mitchellh/go-homedir"
	"path/filepath"
	"task/internal/commands"
	"task/internal/db"
)

func main() {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(homeDir, "tasks.db")
	cfg := Config{db.DBConfig{DBPath: dbPath}}
	dbClient, err := cfg.DBConfig.Init()
	if err != nil {
		panic(err)
	}
	err = dbClient.CreateBucket("tasks")
	if err != nil {
		panic(err)
	}
	err = commands.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
