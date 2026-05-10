# Task Tracker CLI

A simple command-line task management application built with Go. Track your tasks with descriptions, statuses, and timestamps.

**Project URL**: https://roadmap.sh/projects/task-tracker

## Features

- **Add Tasks**: Create new tasks with descriptions
- **Update Tasks**: Modify task descriptions or change their status
- **List Tasks**: View all tasks in a formatted table
- **Delete Tasks**: Remove tasks from your list (coming soon)
- **Persistent Storage**: Tasks are saved in JSON format

## Task Statuses

- `todo` - Task is pending
- `in-progress` - Task is currently being worked on
- `done` - Task is completed

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Saim1123/task_tracker_cli.git
cd task_tracker_cli
```

2. Build the application:
```bash
go build -o task-tracker
```

## Usage

Run the application:
```bash
./task-tracker
```

You'll be presented with a menu:
```
Select Action:
1. Add Task
2. Update Task
3. List Tasks
4. Delete Task
Enter choice (1-4):
```

### Adding a Task

1. Select option `1`
2. Enter your task description
3. Task will be saved with status `todo`

### Updating a Task

1. Select option `2`
2. Enter the task ID
3. Choose what to update:
   - Description: Enter new description
   - Status: Select from todo, in-progress, or done

### Listing Tasks

1. Select option `3`
2. View all tasks in a formatted table showing:
   - ID
   - Description
   - Status
   - Created date

## Project Structure

```
task_tracker_cli/
├── main.go                    # Entry point and menu interface
├── models/
│   └── task.go               # Task model and status definitions
├── services/
│   └── task_services.go      # Business logic for task operations
├── utils/
│   └── file.go               # File handling utilities
├── go.mod                    # Go module definition
└── task.json                 # Task storage (auto-generated)
```

## Requirements

- Go 1.25.0 or higher

## Data Storage

Tasks are stored in `task.json` in the project directory. The file is automatically created on first run.

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the MIT License.
