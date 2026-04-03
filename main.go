package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Saim1123/task_tracker_cli/services"
)

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
		services.AddTask()
	case 2:
		services.UpdateTask()
	case 3:
		services.ListTasks()
	case 4:
		services.DeleteTask()
	default:
		fmt.Println("Invalid choice")
	}
}
