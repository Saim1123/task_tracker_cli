package main

import (
	"bufio"
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
	id          int
	description string
	status      Status
	created_at  time.Time
	updated_at  time.Time
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select Action:")
	fmt.Println("1. Add Task")
	fmt.Println("2. Update Task")
	fmt.Println("3. List Tasks")
	fmt.Println("4. Delete Task")

	fmt.Print("Enter choice (1-4): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, error := strconv.Atoi(input)
	if error != nil {
		fmt.Println("Invalid choice")
		return
	}

	switch choice {
	case 1:
		fmt.Println("Add Task")
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
