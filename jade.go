package jade

import (
	"fmt"
)

// Task: Jade Task
//
// Status: Pending, Completed, Failed
type Task struct {
	// Id: Jade Task ID
	Id int `json:"id"`
	// Name: Jade Task Name
	Name string `json:"name"`
	// Status: Jade Task Status
	Status Status `json:"status"`
}

// Status: Pending, Completed, Failed
type Status string

// String returns the string representation of the Status value.
//
// No parameters.
// Returns a string.
func (s Status) String() string {
	return string(s)
}

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

var chanStatus = make(chan Task, 10)
var chanTask = make(chan Task, 10)

// SetStatus sets the status of the Task to the provided status value.
//
// Takes a Status parameter.
func (t *Task) SetStatus(status Status) {
	t.Status = status
	chanTask <- *t
}

// GetStatus returns the status of the Task.
//
// No parameters.
// Returns a Status value.
func (t Task) GetStatus() Status {
	return t.Status
}

// GetId returns the ID of the Task.
//
// No parameters.
// Returns an integer.
func (t Task) GetId() int {
	return t.Id
}

// GetName returns the name of the Task.
//
// No parameters.
// Returns a string.
func (t Task) GetName() string {
	return t.Name
}

// NewTask creates a new Task and sends it to the task channel.
//
// Takes a Task parameter.
// Returns a Task.
func NewTask(task Task) Task {
	chanTask <- task

	return task
}

// process is a Go function that continuously listens for tasks from the chanTask channel.
// It uses a select statement to handle multiple channels.
// When a task is received, it checks the status of the task using the GetStatus method.
// Depending on the status, it sends the task to the chanStatus channel.
// If the status is not recognized, it prints "default".
// This function does not have any parameters or return types.
func process() {
	for {
		select {
		case task := <-chanTask:
			switch task.GetStatus() {
			case StatusPending:
				chanStatus <- task
			case StatusCompleted:
				chanStatus <- task
			case StatusFailed:
				chanStatus <- task
			default:
				fmt.Println("default")
			}
		}

	}
}

// Init initializes the system by starting the process goroutine.
//
// No parameters.
// No return values.
func Init() {
	go process()
}

// listen returns a read-only channel of Task objects.
//
// No parameters.
// Returns a channel of Task objects.
func listen() <-chan Task {
	return chanStatus
}

// Listen returns a read-only channel of Task objects.
//
// No parameters.
// Returns a channel of Task objects.
func Listen() <-chan Task {
	return listen()
}
