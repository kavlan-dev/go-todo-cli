package app

import (
	"encoding/json"
	"fmt"
	"go-todo-cli/internal/config"
	"go-todo-cli/internal/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var taskManager *todoList

type todoList struct {
	Tasks  []models.Task `json:"tasks"`
	NextId int           `json:"next_id"`
}

func LoadTasks(cmd *cobra.Command, args []string) {
	file := config.TasksPath
	taskManager = &todoList{
		Tasks:  make([]models.Task, 0),
		NextId: 1,
	}

	if data, err := os.ReadFile(file); err == nil {
		json.Unmarshal(data, taskManager)
	}
}

func SaveTasks(cmd *cobra.Command, args []string) {
	data, err := json.MarshalIndent(taskManager, "", "  ")
	if err != nil {
		fmt.Printf("Ошибка сохранения задачи: %v\n", err)
		return
	}

	os.WriteFile(config.TasksPath, data, 0644)
}

func AddTask(cmd *cobra.Command, args []string) {
	description := strings.Join(args, " ")

	task := models.Task{
		Id:        taskManager.NextId,
		Content:   description,
		Done:      false,
		CreatedAt: time.Now(),
	}

	if err := validateTask(task); err != nil {
		fmt.Printf("Ошибка сохранения задачи: %v\n", err)
		return
	}

	taskManager.Tasks = append(taskManager.Tasks, task)
	taskManager.NextId++

	fmt.Printf("Добавлена задача #%d: %s\n", task.Id, task.Content)
}

func ListTasks(cmd *cobra.Command, args []string) {
	if len(taskManager.Tasks) == 0 {
		fmt.Println("Задачи не найдены.")
		return
	}

	fmt.Printf("%-4s %-10s %-50s %s\n", "ID", "СТАТУС", "ОПИСАНИЕ", "СОЗДАН")
	fmt.Println(strings.Repeat("-", 90))

	for _, task := range taskManager.Tasks {
		status := "В процессе"
		if task.Done {
			status = "Выполнен"
		}

		fmt.Printf("%-4d %-10s %-50s %s\n",
			task.Id,
			status,
			truncate(task.Content, 50),
			task.CreatedAt.Format("2026-01-02"))
	}
}

func CompleteTask(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Не верный id задачи: %s\n", args[0])
		os.Exit(1)
	}

	for i := range taskManager.Tasks {
		if taskManager.Tasks[i].Id == id {
			if taskManager.Tasks[i].Done {
				fmt.Printf("Задача #%d уже была выполнена\n", id)
				return
			}

			now := time.Now()
			taskManager.Tasks[i].Done = true
			taskManager.Tasks[i].CompletedAt = &now

			fmt.Printf("Выполнена задача #%d: %s\n", id, taskManager.Tasks[i].Content)
			return
		}
	}

	fmt.Printf("Задача #%d не найдена\n", id)
	os.Exit(1)
}

func DeleteTask(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Не верный id задачи: %s\n", args[0])
		os.Exit(1)
	}

	for i, task := range taskManager.Tasks {
		if task.Id == id {
			taskManager.Tasks = append(taskManager.Tasks[:i], taskManager.Tasks[i+1:]...)
			fmt.Printf("Удалена задача #%d: %s\n", id, task.Content)
			return
		}
	}

	fmt.Printf("Задача #%d не найдена\n", id)
	os.Exit(1)
}

func ClearAllTasks(cmd *cobra.Command, args []string) {
	taskManager.Tasks = []models.Task{}
	taskManager.NextId = 1
	fmt.Println("Все задачи очищены")
}

func CompleteAllTasks(cmd *cobra.Command, args []string) {
	now := time.Now()
	for i := range taskManager.Tasks {
		if !taskManager.Tasks[i].Done {
			taskManager.Tasks[i].Done = true
			taskManager.Tasks[i].CompletedAt = &now
		}
	}

	fmt.Println("Все задачи отмечены как выполненные")
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

func validateTask(task models.Task) error {
	if len(task.Content) > int(config.MaxTaskLength) {
		return fmt.Errorf("текст задачи не должен превышать %d символов\n", config.MaxTaskLength)
	}

	if strings.TrimSpace(task.Content) == "" {
		return fmt.Errorf("новый текст задачи не может быть пустым")
	}

	for _, t := range taskManager.Tasks {
		if strings.EqualFold(t.Content, task.Content) {
			return fmt.Errorf("задача с таким заголовком уже существует")
		}
	}

	return nil
}
