package main

import (
	"fmt"
	"strings"
)

// Define constants for task titles
const (
	DevEnvSetupTaskTitle  = "Set up the Go development environment"
	GoSyntaxLearningTitle = "Learn Go Syntax"
	TodoAppCreationTitle  = "Create a Simple Todo App"
)

// Task struct for better structure
type Task struct {
	Title    string
	SubTasks []string
}

// TaskList to manage multiple tasks
type TaskList struct {
	Tasks []Task
}

// Constructor for creating a new task
func NewTask(title string, subTasks []string) Task {
	return Task{
		Title:    title,
		SubTasks: subTasks,
	}
}

// Method to print the task and subtasks
func (t Task) Print() {
	fmt.Println(t.Title)
	for _, subTask := range t.SubTasks {
		fmt.Println("  -", subTask)
	}
	fmt.Println()
}

// Method to add a task to a task list
func (tl *TaskList) AddTask(task Task) {
	tl.Tasks = append(tl.Tasks, task)
}

// Method to print all tasks in the list
func (tl TaskList) PrintAll() {
	fmt.Println("###### Golang Learning To-Dos: ######")
	for _, task := range tl.Tasks {
		task.Print()
	}
}

func main() {
	// Initialize the task list
	taskList := TaskList{}

	// Create tasks using the constructor
	taskList.AddTask(NewTask(DevEnvSetupTaskTitle, []string{
		"Install Go on my machine via HomeBrew by running brew install go",
		"Setup Neovim and install Go via Mason",
		"Verify if everything is working",
	}))

	taskList.AddTask(NewTask(GoSyntaxLearningTitle, []string{
		"Understand variables, data types, and constants",
		"Learn how to declare and use functions",
		"Get familiar with control structure (if, switch, for loops)",
		"Explore Go's Data Structure",
		"Dive into Go's Concurrency Model",
		"Understand Error Handling in Go",
		"Explore Go's Standard Library",
		"Practice Writing Unit Tests in Go",
		"Learn About Go Modules and Dependency Management",
		"Understand Go’s Memory Management",
		"Learn About Go’s Web Frameworks",
		"Learn About Go’s Performance Optimization",
	}))

	taskList.AddTask(NewTask(TodoAppCreationTitle, []string{
		"Create a Directory for my project by running 'mkdir go-todo-app' in the terminal",
		"Navigate to go-todo-app by running 'cd go-todo-app'",
		"Initialize the Go Module by running this command 'go mod init go-todo-app'",
		"Create a Go File by running 'touch main.go'",
		"Import package main",
		"Import 'fmt'",
		"Start writing Go code",
		"Run the Go Program 'go run main.go'",
	}))

	// Print all tasks in the list
	taskList.PrintAll()
}
