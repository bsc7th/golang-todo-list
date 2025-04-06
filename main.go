package main

import (
	"fmt"
	"net/http"
	"strings"
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
func (t Task) Print() string {
	var result strings.Builder
	result.WriteString(t.Title + "\n")
	for _, subTask := range t.SubTasks {
		result.WriteString(fmt.Sprintf("  - %s\n", subTask))
	}
	return result.String()
}

// Method to add a task to a task list
func (tl *TaskList) AddTask(task Task) {
	tl.Tasks = append(tl.Tasks, task)
}

// Method to print all tasks in the list as a string
func (tl TaskList) PrintAll() string {
	var result strings.Builder
	result.WriteString("###### Golang Learning To-Dos: ######\n")
	for i, task := range tl.Tasks {
		result.WriteString(task.Print())
		if i < len(tl.Tasks)-1 { // Avoid adding extra newline after the last task
			result.WriteString("\n") // Add a new line after each main task
		}
	}
	return result.String()
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

// Greeting handler function to show tasks
func showTasks(w http.ResponseWriter, r *http.Request) {
	// Initialize the task list
	taskList := TaskList{}

	// Add tasks to the task list
	taskList.AddTask(createDevEnvSetupTask())
	taskList.AddTask(createGoSyntaxLearningTask())
	taskList.AddTask(createTodoAppCreationTask())

	// Print all tasks as a string and write to the HTTP response
	tasks := taskList.PrintAll()
	fmt.Fprintln(w, tasks)
}

func main() {
	// Set up HTTP handler to show tasks
	http.HandleFunc("/", showTasks)

	// Start the server on port 3000
	http.ListenAndServe(":3000", nil)
}
