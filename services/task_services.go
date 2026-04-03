package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
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

func ListTasks() {
	err := utils.EnsureFile("task.json")
	if err != nil {
		fmt.Println("Failed to ensure task.json")
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

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Println("")
	fmt.Fprintln(w, "ID\tDESCRIPTION\tSTATUS\tCREATED")
	fmt.Fprintln(w, "--\t-----------\t------\t-------")

	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
			task.ID,
			task.Description,
			task.Status,
			task.Created_at.Format("2006-01-02 15:04"),
		)
	}

	w.Flush()
}
