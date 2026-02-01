package config

import (
	"os"
	"strconv"
)

var MaxTaskLength uint
var TasksPath string

func InitConfig() error {
	maxTaskLengthStr, err := strconv.Atoi(envOrDefault("MAX_TASK_LENGTH", "200"))
	if err != nil {
		return err
	}

	MaxTaskLength = uint(maxTaskLengthStr)
	TasksPath = envOrDefault("TASKS_PATH", "tasks.json")

	return nil
}

func envOrDefault(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if value == "" {
		value = defaultValue
	}
	return value
}
