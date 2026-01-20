package main

import (
	"go-todo-cli/cmd"
	"go-todo-cli/internal/app"
	"go-todo-cli/internal/config"
)

func main() {
	config.SetupConfig()
	app.New()
	cmd.Execute()
}
