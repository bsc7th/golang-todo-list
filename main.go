package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rivo/tview"
)

// Task structures
type Task struct {
	Title    string   `json:"title"`
	SubTasks []string `json:"subTasks"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func NewTask(title string, subTasks []string) Task {
	return Task{Title: title, SubTasks: subTasks}
}

func createDevEnvSetupTask() Task {
	return NewTask("Set up the Go development environment", []string{
		"Install Go via HomeBrew",
		"Setup Neovim and Go",
		"Verify everything is working",
	})
}

func createGoSyntaxLearningTask() Task {
	return NewTask("Learn Go Syntax", []string{
		"Variables and constants",
		"Functions",
		"Control structures",
		"Data structures",
		"Concurrency model",
	})
}

func createTodoAppCreationTask() Task {
	return NewTask("Create a Simple Todo App", []string{
		"Create project dir",
		"Initialize Go module",
		"Write and run Go code",
	})
}

// API handler
func showTasks(w http.ResponseWriter, r *http.Request) {
	taskList := TaskList{
		Tasks: []Task{
			createDevEnvSetupTask(),
			createGoSyntaxLearningTask(),
			createTodoAppCreationTask(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(taskList); err != nil {
		http.Error(w, "Unable to encode tasks to JSON", http.StatusInternalServerError)
	}
}

func main() {
	// Start API server
	go func() {
		http.HandleFunc("/", showTasks)
		fmt.Println("📡 Server running at http://localhost:4175")
		if err := http.ListenAndServe(":4175", nil); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	// Create tasks
	tasks := []Task{
		createDevEnvSetupTask(),
		createGoSyntaxLearningTask(),
		createTodoAppCreationTask(),
	}

	app := tview.NewApplication()
	list := tview.NewList()

	for i, task := range tasks {
		// Capture index to use in the handler
		index := i
		list.AddItem(task.Title, "Press Enter to view subtasks", 0, func() {
			showSubTasksModal(app, task.Title, tasks[index].SubTasks, list)
		})
	}

	list.AddItem("Quit", "Exit the app", 'q', func() {
		app.Stop()
	})

	list.SetBorder(true).SetTitle("📋 My Todo List").SetTitleAlign(tview.AlignLeft)

	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}

// Show subtasks in a modal popup
func showSubTasksModal(app *tview.Application, title string, subtasks []string, previousPage tview.Primitive) {
	modal := tview.NewModal().
		SetText("Subtasks for:\n\n" + title + "\n\n" + strings.Join(subtasks, "\n")).
		AddButtons([]string{"Back"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.SetRoot(previousPage, true)
		})

	app.SetRoot(modal, true).SetFocus(modal)
}
