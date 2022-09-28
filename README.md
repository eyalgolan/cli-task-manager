# Cli Task Manager

[![Build](https://github.com/eyalgolan/cli-task-manager/actions/workflows/build.yml/badge.svg)](https://github.com/eyalgolan/cli-task-manager/actions/workflows/build.yml)

This is a CLI tool that can be used to manage your TODOs in the terminal.

I started this small project by following [this Gophercises exercise](https://github.com/gophercises/task), and later expanded it to make it more testable.

## Usage

Basic usage of the tool:
```sybase
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
You have completed the "review talk proposal" task.

$ task list
You have the following tasks:
1. some task description
```