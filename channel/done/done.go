package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, worker worker) {
	for n := range worker.in {
		fmt.Printf("Worker %d recieved %c \n", id, n)
		worker.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() { wg.Done() },
	}
	go doWorker(id, w)
	return w
}

type worker struct {
	in   chan int
	done func()
}

func chanDemo() {
	waitGroup := sync.WaitGroup{}
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &waitGroup)
	}

	waitGroup.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i

	}

	for i, worker := range workers {
		worker.in <- 'A' + i

	}

	waitGroup.Wait()
}

func main() {
	chanDemo()
}
