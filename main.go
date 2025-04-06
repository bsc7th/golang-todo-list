package main

import "fmt"

// Define constants for task titles
const (
    DevEnvSetupTaskTitle   = "Set up the Go development environment"
    GoSyntaxLearningTitle  = "Learn Go Syntax"
    TodoAppCreationTitle   = "Create a Simple Todo App"
)

// Define a Task struct for better structure
type Task struct {
    Title    string
    SubTasks []string
}

// Constructor for creating a new task
func NewTask(title string, subTasks []string) Task {
    return Task{
        Title:    title,
        SubTasks: subTasks,
    }
}

func main() {
    // Create tasks using the constructor for clarity
    taskList := []Task{
        NewTask(DevEnvSetupTaskTitle, []string{
            "Install Go on my machine via HomeBrew by running brew install go",
            "Setup Neovim and install Go via Mason",
            "Verify if everything is working",
        }),
        NewTask(GoSyntaxLearningTitle, []string{
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
        }),
        NewTask(TodoAppCreationTitle, []string{
            "Create a Directory for my project by running 'mkdir go-todo-app' in the terminal",
            "Navigate to go-todo-app by running 'cd go-todo-app'",
            "Initialize the Go Module by running this command 'go mod init go-todo-app'",
            "Create a Go File by running 'touch main.go'",
            "Import package main",
            "Import 'fmt'",
            "Start writing Go code",
            "Run the Go Program 'go run main.go'",
        }),
    }

    fmt.Println("###### Golang Learning To-Dos: ######")
    for _, task := range taskList {
        printTask(task)
    }
}

// Helper function to print a task and its subtasks
func printTask(task Task) {
    fmt.Println(task.Title)
    for _, subTask := range task.SubTasks {
        fmt.Println("  -", subTask) // Added indentation for clarity
    }
    fmt.Println() // Blank line between tasks
}
