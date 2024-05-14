# Jade

Jade is a Go package that provides features for managing tasks goroutine with different statuses such as pending, completed, and failed.

## Installation

To use Jade in your Go project, you can simply install it using `go get`:

```sh
go get github.com/refaldyrk/jade
```

## Usage

Import Jade package in your Go code:

```go
import "github.com/refaldyrk/jade"
```

### Creating a Task

You can create a new task using the `NewTask` function:

```go
task := jade.Task{
    Id:     1,
    Name:   "task-1",
    Status: jade.StatusPending,
}

newTask := jade.NewTask(task)
```

### Changing Task Status

You can change the status of a task using the `setStatus` method:

```go
newTask.setStatus(jade.StatusCompleted)
```

### Listening for Task Status Changes

You can listen for task status changes using the `StartListening` function:

```go
listen := jade.Listen()
go func() {
	for t := range listen {
        switch t.GetStatus() {
            case jade.StatusPending:
                fmt.Println("Task ", t.Name, "Is", t.GetStatus())
            case jade.StatusCompleted:
                fmt.Println("Task ", t.Name, "Is", t.GetStatus())
            case jade.StatusFailed:
				fmt.Println("Task ", t.Name, "Is", t.GetStatus())
		}
    }
}()
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or create a pull request.

## Example
In Folder Example Hehe