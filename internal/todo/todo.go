package todo

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ErrTodoNotFound  = errors.New("todo not found")
	ErrInvalidStatus = errors.New("invalid todo status")
)

type TodoStatus int

const (
	StatusTodo TodoStatus = iota
	StatusInProgress
	StatusDone
	StatusStopped
)

type Todo struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	Status      TodoStatus
}

var (
	statusMapping = map[string]TodoStatus{
		"todo":        StatusTodo,
		"in_progress": StatusInProgress,
		"done":        StatusDone,
		"stopped":     StatusStopped,
	}
)

func (h *Handler) List() []Todo {
	return h.list
}

func (h *Handler) Get(id int) (Todo, error) {
	for _, todo := range h.list {
		if todo.ID == id {
			return todo, nil
		}
	}

	return Todo{}, ErrTodoNotFound
}

func (h *Handler) Create(title, description string) Todo {
	h.lastID++
	newTodo := Todo{
		ID:          h.lastID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		Status:      StatusTodo,
	}

	h.list = append(h.list, newTodo)

	return newTodo
}

func (h *Handler) Edit(id int, title, description string) (Todo, error) {
	var foundIndex int = -1

	for i := range h.list {
		if h.list[i].ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return Todo{}, ErrTodoNotFound
	}

	h.list[foundIndex].Title = title
	h.list[foundIndex].Description = description

	return h.list[foundIndex], nil
}

func (h *Handler) Delete(id int) (Todo, error) {
	var foundIndex int = -1

	for i := range h.list {
		if h.list[i].ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return Todo{}, ErrTodoNotFound
	}

	found := h.list[foundIndex]
	h.list = append(h.list[:foundIndex], h.list[foundIndex+1:]...)
	return found, nil
}

func (h *Handler) UpdateStatus(id int, status string) (Todo, error) {
	var foundIndex int = -1

	for i := range h.list {
		if h.list[i].ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return Todo{}, ErrTodoNotFound
	}

	stts, ok := statusMapping[status]
	if !ok {
		return Todo{}, ErrInvalidStatus
	}

	h.list[foundIndex].Status = stts
	return h.list[foundIndex], nil
}

func (h *Handler) Save() error {
	if err := os.Remove("./todolist.csv"); err != nil {
		return err
	}

	f, err := os.OpenFile("./todolist.csv", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, todo := range h.list {
		if todo.ID == 0 {
			continue
		}
		todoStr := fmt.Sprintf("%d;%s;%s;%d;%s\n", todo.ID, todo.Title, todo.Description, todo.Status, todo.CreatedAt.Format(time.RFC3339))
		_, err = f.Write([]byte(todoStr))
		if err != nil {
			return err
		}
	}

	return f.Sync()
}

func (h *Handler) Load() error {
	f, err := os.OpenFile("./todolist.csv", os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ";")

		if len(items) < 5 {
			continue
		}

		id, _ := strconv.ParseInt(items[0], 10, 32)
		status, _ := strconv.ParseInt(items[3], 10, 32)

		t, _ := time.Parse(time.RFC3339, items[4])

		todo := Todo{
			ID:          int(id),
			Title:       items[1],
			Description: items[2],
			Status:      TodoStatus(status),
			CreatedAt:   t,
		}
		h.lastID = int(id)
		h.list = append(h.list, todo)
	}
	return nil
}
