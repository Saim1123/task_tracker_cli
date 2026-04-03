package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Saim1123/task_tracker_cli/models"
	"github.com/Saim1123/task_tracker_cli/utils"
)

func AddTask() {
	err := utils.EnsureFile("task.json")
	if err != nil {
		fmt.Println("Failed to ensure task.json")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Add Todo: ")

	description, err := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	data, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("Error reading task.json", err)
		return
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing task.json:", err)
		return
	}

	newId := len(tasks) + 1

	task := models.Task{
		ID:          newId,
		Description: description,
		Status:      models.StatusTodo,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	tasks = append(tasks, task)

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error converting tasks to JSON:", err)
		return
	}

	err = os.WriteFile("task.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to task.json:", err)
		return
	}

	fmt.Println("Task added successfully")
}

func UpdateTask() {}

func DeleteTask() {}

func ListTasks() {}
