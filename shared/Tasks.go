package shared

import (
	"time"
)

type Task struct {
	Title   string
	Message string
	Rest    time.Duration
}

var (
	TASKS      = []Task{}
	TASK_INDEX = 0
)

func (t Task) AppendTask() {
	// Append a copy of the task to TASKS
	TASKS = append(TASKS, t)
}

func NewTask(title string, message string, rest time.Duration) Task {
	// Default rest period if nil == 5 minutes
	rest = time.Duration(rest * time.Second)
	// Makes new task with inputed fields
	t := Task{Title: title, Message: message, Rest: rest}

	t.AppendTask()

	return t // return object
}

// Returns the current task
func GetTask() Task {
	return TASKS[TASK_INDEX]
}

// Returns the next task in the list
func NextTask() Task {
	TASK_INDEX++
	return TASKS[TASK_INDEX]
}
