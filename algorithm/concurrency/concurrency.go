package concurrency

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers = 100
	bufferSize = 100
)

type Task struct {
	ID    int
	Value string
}

func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		time.Sleep(time.Second)
		results <- fmt.Sprintf("Worker %d processed task %d", id, task.ID)
	}
}

func Concurrency() {
	tasks := make(chan Task, bufferSize) // Channel buffered
	results := make(chan string, bufferSize)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	go func() {
		for i := 1; i <= 200; i++ {
			task := Task{ID: i, Value: fmt.Sprintf("Task %d", i)}

			for {
				if len(tasks) < cap(tasks) {
					tasks <- task
					break
				} else {
					fmt.Println("Buffer đầy, chờ...")
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
		close(tasks)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All tasks processed")
}
