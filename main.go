package main

import (
	"fmt"
	"net/http"
)

// Define constants for task titles and subtasks
const (
	DevEnvSetupTaskTitle  = "Set up the Go development environment"
	GoSyntaxLearningTitle = "Learn Go Syntax"
	TodoAppCreationTitle  = "Create a Simple Todo App"
)

const (
	// Subtasks for DevEnvSetupTask
	SubtaskInstallGo          = "Install Go on my machine via HomeBrew by running brew install go"
	SubtaskSetupNeovim        = "Setup Neovim and install Go via Mason"
	SubtaskVerifyEnvironment  = "Verify if everything is working"

	// Subtasks for GoSyntaxLearning
	SubtaskUnderstandVariables = "Understand variables, data types, and constants"
	SubtaskLearnFunctions      = "Learn how to declare and use functions"
	SubtaskControlStructures   = "Get familiar with control structure (if, switch, for loops)"
	SubtaskExploreDataStructures = "Explore Go's Data Structure"
	SubtaskLearnConcurrency    = "Dive into Go's Concurrency Model"
	SubtaskErrorHandling       = "Understand Error Handling in Go"
	SubtaskExploreStandardLib  = "Explore Go's Standard Library"
	SubtaskPracticeUnitTests   = "Practice Writing Unit Tests in Go"
	SubtaskLearnModules        = "Learn About Go Modules and Dependency Management"
	SubtaskMemoryManagement    = "Understand Go’s Memory Management"
	SubtaskExploreWebFrameworks = "Learn About Go’s Web Frameworks"
	SubtaskPerformanceOptimization = "Learn About Go’s Performance Optimization"

	// Subtasks for TodoAppCreation
	SubtaskCreateProjectDir     = "Create a Directory for my project by running 'mkdir go-todo-app' in the terminal"
	SubtaskNavigateToProject    = "Navigate to go-todo-app by running 'cd go-todo-app'"
	SubtaskInitializeModule     = "Initialize the Go Module by running this command 'go mod init go-todo-app'"
	SubtaskCreateGoFile         = "Create a Go File by running 'touch main.go'"
	SubtaskImportMainPackage    = "Import package main"
	SubtaskImportFmtPackage     = "Import 'fmt'"
	SubtaskWriteGoCode          = "Start writing Go code"
	SubtaskRunGoProgram         = "Run the Go Program 'go run main.go'"
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
		fmt.Printf("  - %s\n", subTask)
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

// Task creation functions to keep main clean and organized
func createDevEnvSetupTask() Task {
	return NewTask(DevEnvSetupTaskTitle, []string{
		SubtaskInstallGo,
		SubtaskSetupNeovim,
		SubtaskVerifyEnvironment,
	})
}

func createGoSyntaxLearningTask() Task {
	return NewTask(GoSyntaxLearningTitle, []string{
		SubtaskUnderstandVariables,
		SubtaskLearnFunctions,
		SubtaskControlStructures,
		SubtaskExploreDataStructures,
		SubtaskLearnConcurrency,
		SubtaskErrorHandling,
		SubtaskExploreStandardLib,
		SubtaskPracticeUnitTests,
		SubtaskLearnModules,
		SubtaskMemoryManagement,
		SubtaskExploreWebFrameworks,
		SubtaskPerformanceOptimization,
	})
}

func createTodoAppCreationTask() Task {
	return NewTask(TodoAppCreationTitle, []string{
		SubtaskCreateProjectDir,
		SubtaskNavigateToProject,
		SubtaskInitializeModule,
		SubtaskCreateGoFile,
		SubtaskImportMainPackage,
		SubtaskImportFmtPackage,
		SubtaskWriteGoCode,
		SubtaskRunGoProgram,
	})
}

func main() {
	// Initialize the task list
	taskList := TaskList{}
	// This function takes two parameters:
	// 1. Address: The address to which the server should bind (in this case, ":3000").
	// 2. Handler: A request handler (currently set to nil, meaning default multiplexer is used).

	// When the application starts, it will return a '404' error for any unhandled requests.
	// This indicates that the server is running, but there are no routes defined to handle the request.
	// If the server were not running, attempting to access the address would result in "Site can't be reached" error.
	http.ListenAndServe(":3000", nil)

	// Add tasks to the task list
	taskList.AddTask(createDevEnvSetupTask())
	taskList.AddTask(createGoSyntaxLearningTask())
	taskList.AddTask(createTodoAppCreationTask())

	// Print all tasks in the list
	taskList.PrintAll()
}
