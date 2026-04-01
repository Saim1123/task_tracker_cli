package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func ensureFile(filename string) error {
	if !fileExists(filename) {
		file, err := os.Create(filename)

		if err != nil {
			return err
		}

		defer file.Close()

		file.WriteString("[]")
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select Action:")
	fmt.Println("1. Add Task")
	fmt.Println("2. Update Task")
	fmt.Println("3. List Tasks")
	fmt.Println("4. Delete Task")

	fmt.Print("Enter choice (1-4): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid choice")
		return
	}

	switch choice {
	case 1:
		err := ensureFile("task.json")
		if err != nil {
			fmt.Println("Failed to ensure task.json")
			return
		}

		reader = bufio.NewReader(os.Stdin)

		fmt.Println("Add Todo: ")

		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		data, err := os.ReadFile("task.json")
		if err != nil {
			fmt.Println("Error reading task.json", err)
			return
		}

		var tasks []Task
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error parsing task.json:", err)
			return
		}

		newId := len(tasks) + 1

		task := Task{
			ID:          newId,
			Description: description,
			Status:      StatusTodo,
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
	case 2:
		fmt.Println("Update Task")
	case 3:
		fmt.Println("List Tasks")
	case 4:
		fmt.Println("Delete Task")
	default:
		fmt.Println("Invalid choice")
	}
}
