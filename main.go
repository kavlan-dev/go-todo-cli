package main

import (
	"go-todo-cli/cmd"
	"go-todo-cli/internal/config"
)

func main() {
	config.SetupConfig()
	cmd.Execute()
}
