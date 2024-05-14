package example

import (
	"fmt"
	"github.com/refaldyrk/jade"
	"time"
)

func main() {
	go jade.Init()

	t := jade.NewTask(jade.Task{
		Id:     1,
		Name:   "task-sample",
		Status: jade.StatusPending,
	})

	time.Sleep(1 * time.Second)

	t.SetStatus(jade.StatusCompleted)
	fmt.Println(t)

	time.Sleep(1 * time.Second)

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

	go func() {
		time.Sleep(2 * time.Second)
		for i := 0; i < 50; i++ {
			t := jade.NewTask(jade.Task{
				Id:     i,
				Name:   "task-" + fmt.Sprint(i),
				Status: jade.StatusFailed,
			})

			t.SetStatus(jade.StatusPending)
			time.Sleep(1 * time.Second)
			t.SetStatus(jade.StatusCompleted)
		}
	}()

	for {
	}
}
