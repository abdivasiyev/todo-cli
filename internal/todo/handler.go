package todo

type Handler struct {
	list   []Todo
	lastID int
}

func NewHandler() *Handler {
	return &Handler{list: make([]Todo, 0), lastID: 0}
}
