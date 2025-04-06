package main

import "fmt"

func main() {
    tasks := map[string][]string{
        "Set up the Go development environment": {
            "Install Go on my machine via HomeBrew by running brew install go",
            "Setup Neovim and install Go via Mason",
            "Verify if everything is working",
        },
        "Learn Go Syntax": {
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
        },
        "Create a Simple Todo App": {
            "Create a Directory for my project by running 'mkdir go-todo-app' in the terminal",
            "Navigate to go-todo-app by running 'z go-todo-app'",
            "Initialize the Go Module by running this command 'go mod init go-todo-app'",
            "Create a Go File by running 'touch main.go'",
            "Import package main",
            "Import 'fmt'",
            "Start writing Go code",
            "Run the Go Program 'go run main.go'",
        },
    }

    fmt.Println("###### Golang Learning To-Dos: ######")
    for task, subTasks := range tasks {
        fmt.Println(task)
        for _, subTask := range subTasks {
            fmt.Println(subTask)
        }
        fmt.Println() // new line
    }
}
