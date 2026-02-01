package main

import (
	"go-todo-cli/cmd"
	"go-todo-cli/internal/config"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	cmd.Execute()
}
