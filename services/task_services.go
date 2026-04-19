package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func UpdateTask() {
	err := utils.EnsureFile("task.json")
	if err != nil {
		fmt.Println("Failed to ensure task.json")
		return
	}

	data, err := os.ReadFile("task.json")
	if err != nil {
		fmt.Println("Error reading task.json:", err)
		return
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing task.json:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Task ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID:", err)
		return
	}

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found")
		return
	}

	fmt.Println("What to update?")
	fmt.Println("1. Description")
	fmt.Println("2. Status")

	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)

	choice, err := strconv.Atoi(choiceStr)
	if err != nil {
		fmt.Println("Invalid choice:", err)
		return
	}

	switch choice {
	case 1:
		fmt.Print("Enter new description: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		tasks[index].Description = description

	case 2:
		fmt.Println("Select new status:")
		fmt.Println("1. todo")
		fmt.Println("2. in-progress")
		fmt.Println("3. done")

		statusStr, _ := reader.ReadString('\n')
		statusStr = strings.TrimSpace(statusStr)

		statusInt, err := strconv.Atoi(statusStr)
		if err != nil {
			fmt.Println("Invalid status:", err)
			return
		}

		switch statusInt {
		case 1:
			tasks[index].Status = models.StatusTodo
		case 2:
			tasks[index].Status = models.StatusInProgress
		case 3:
			tasks[index].Status = models.StatusDone
		default:
			fmt.Println("Invalid status option")
			return
		}

	default:
		fmt.Println("Invalid choice")
		return
	}

	tasks[index].Updated_at = time.Now()

	updatedData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	err = os.WriteFile("task.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Task updated successfully ✅")
}

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
