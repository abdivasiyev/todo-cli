# Todo app

1. Todo list
2. Get todo by keyword -> id, key
3. Create todo
4. Delete todo
5. Edit todo by keyword
6. Change todo status

### Todo structure

- id - int
- title - string
- description - string
- status - TodoStatus
- created_at - time

### Flags

- list - returns todo list
- get {id} - returns todo by id
- create {title} {description} - returns created todo
- edit {id} {title} {description} - returns updated todo
- delete {id} - deletes todo and returns deleted
- status {id} {status} - returns ok

### File format

{id};{title};{description};{status};{created_at}

### Statuses

0 - todo
1 - in_progress
2 - done
3 - stopped
