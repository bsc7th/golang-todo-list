package main

import "fmt"

func main() {
	var taskOne = "Set up the Go development in environment"
	var taskOneA = "Install Go on my machine via HomeBrew by running brew install go"
	var taskOneB = "Setup Neovim and install Go via Mason"
	var taskOneC = "Verify if everything is working"

	var taskTwo = "Learn Go Syntax"
	var taskTwoA = "Understand variables, data types, and constants"
	var taskTwoB = "Learn how to declare  and use functions"
	var taskTwoC = "Get familiar with control structure (if, switch for loops)"
	var taskTwoD = "Explore Go's Data Structure"
	var taskTwoE = "Dive into Go' Concurrency Model"
	var taskTwoF = "Understand Error Handling in Go"
	var taskTwoG = "Explore Go's Standard Library"
	var taskTwoH = "Practice Writing Unit Tests in Go"
	var taskTwoI = "Learn About Go Modules and Dependency Management"
	var taskTwoJ = "Understand Go’s Memory Management"
	var taskTwoK = "Learn About Go’s Web Frameworks"
	var taskTwoL = "Learn About Go’s Performance Optimization"

	var taskThree = "Create a Simple Todo App"
	var taskThreeA = "Create a Directory for my project by running 'mkdir go-todo-app' in the terminal"
	var taskThreeB = "Navigate to go-todo-app by running z go-todo-app" // I install zoxide that's why I'm using the command z instead of cd
	var taskThreeC = "Initialize the Go Module by running this command 'go mod init go-todo-app'"
	var taskThreeD = "Create a Go File by running 'touch main.go'"
	var taskThreeE = "Import package main"
	var taskThreeF = "Import 'fmt'"
	var taskThreeG = "Start writing Go code"
	var taskThreeH = "Run the Go Program 'go run main.go'"

	fmt.Println("###### Golang Learning To-Dos: ######")

	fmt.Println(taskOne)
	fmt.Println(taskOneA)
	fmt.Println(taskOneB)
	fmt.Println(taskOneC)

	fmt.Println() // new line
	fmt.Println(taskTwo)
	fmt.Println(taskTwoA)
	fmt.Println(taskTwoB)
	fmt.Println(taskTwoC)

	fmt.Println()  // new line
	fmt.Println(taskThree)
	fmt.Println(taskThreeA)
	fmt.Println(taskThreeB)
	fmt.Println(taskThreeC)
	fmt.Println(taskThreeD)
	fmt.Println(taskThreeE)
	fmt.Println(taskThreeF)
	fmt.Println(taskThreeG)
	fmt.Println(taskThreeH)
}

