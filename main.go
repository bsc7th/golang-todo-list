package main

import "fmt"

func main() {
	var taskOne = "Set up the Go development in environment"
	var taskOneA = "Install Go on my machine via HomeBrew by running brew install go"
	var taskOneB = "Setup Neovim and install Go via Mason"
	var taskOneC = "Verify if everything is working"

	var taskTwo = "Create a Simple Todo App"
	var taskTwoA = "Create a Directory for my project by running 'mkdir go-todo-app' in the terminal"
	var taskTwoB = "Navigate to go-todo-app by running z go-todo-app" // I install zoxide that's why I'm using the command z instead of cd

	fmt.Println("###### Golang Learning To-Dos: ######")

	fmt.Println(taskOne)
	fmt.Println(taskOneA)
	fmt.Println(taskOneB)
	fmt.Println(taskOneC)

	fmt.Println()  // new line
	fmt.Println(taskTwo)
	fmt.Println(taskTwoA)
	fmt.Println(taskTwoB)

	fmt.Println() // new line
	fmt.Println("Rewards")
	fmt.Println("Reward myself with a cheesecake")
}

