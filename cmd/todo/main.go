package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abdivasiyev/todo-cli/internal/todo"
)

func main() {
	h := todo.NewHandler()

	if err := h.Load(); err != nil {
		fmt.Printf("error while loading todos: %v\n", err)
		os.Exit(1)
	}

	lenArgs := len(os.Args)

	if lenArgs < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	switch strings.TrimLeft(os.Args[1], "-") {
	// list command
	case "list":
		todos := h.List()
		for i := range todos {
			fmt.Printf("[%d] %s - %s; %d\n", todos[i].ID, todos[i].Title, todos[i].Description, todos[i].Status)
		}
		// get command
	case "get":
		if lenArgs < 3 {
			fmt.Println("id not specified")
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, _ := strconv.ParseInt(idStr, 10, 32)

		foundTodo, err := h.Get(int(id))
		if err != nil {
			fmt.Printf("can not get todo: %v", err)
			os.Exit(1)
		}
		fmt.Printf("[%d] %s - %s; %d\n", foundTodo.ID, foundTodo.Title, foundTodo.Description, foundTodo.Status)
	case "create":
		var title, description string
		if lenArgs < 3 {
			fmt.Println("title not specified")
			os.Exit(1)
		}
		title = os.Args[2]
		if lenArgs == 4 {
			description = os.Args[3]
		}

		createdTodo := h.Create(title, description)
		fmt.Printf("[%d] %s - %s; %d\n", createdTodo.ID, createdTodo.Title, createdTodo.Description, createdTodo.Status)
	case "edit":
		var description string
		if lenArgs < 3 {
			fmt.Println("id not specified")
			os.Exit(1)
		}
		if lenArgs < 4 {
			fmt.Println("title not specified")
			os.Exit(1)
		}
		id, _ := strconv.ParseInt(os.Args[2], 10, 32)
		title := os.Args[3]

		if lenArgs == 5 {
			description = os.Args[4]
		}

		editedTodo, err := h.Edit(int(id), title, description)
		if err != nil {
			fmt.Printf("can not get todo: %v", err)
			os.Exit(1)
		}
		fmt.Printf("[%d] %s - %s; %d\n", editedTodo.ID, editedTodo.Title, editedTodo.Description, editedTodo.Status)
	case "delete":
		if lenArgs < 3 {
			fmt.Println("id not specified")
			os.Exit(1)
		}
		id, _ := strconv.ParseInt(os.Args[2], 10, 32)

		deletedTodo, err := h.Delete(int(id))
		if err != nil {
			fmt.Printf("can not get todo: %v", err)
			os.Exit(1)
		}
		fmt.Printf("[%d] %s - %s; %d\n", deletedTodo.ID, deletedTodo.Title, deletedTodo.Description, deletedTodo.Status)
	case "status":
		if lenArgs < 3 {
			fmt.Println("id not specified")
			os.Exit(1)
		}
		if lenArgs < 4 {
			fmt.Println("status not specified")
			os.Exit(1)
		}
		id, _ := strconv.ParseInt(os.Args[2], 10, 32)
		status := os.Args[3]

		updatedTodo, err := h.UpdateStatus(int(id), strings.ToLower(status))
		if err != nil {
			fmt.Printf("can not get todo: %v", err)
			os.Exit(1)
		}
		fmt.Printf("[%d] %s - %s; %d\n", updatedTodo.ID, updatedTodo.Title, updatedTodo.Description, updatedTodo.Status)
	default:
		fmt.Println("invalid command argument")
		os.Exit(1)
	}

	if err := h.Save(); err != nil {
		fmt.Printf("error while saving file: %v", err)
		os.Exit(1)
	}
}
