run:
	go run cmd/todo/main.go

build:
	go build -o todo ./cmd/todo

.PHONY: run build
.DEFAULT_GOAL:=run