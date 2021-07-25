package main

import "time"

const TaskCount = 1000

func f2() {

	var (
		counter int
		done    = make(chan struct{}, TaskCount)
		mu      = make(chan struct{}, 1)
		t       = time.Now()
	)

	mu <- struct{}{} // проинициализировать "мьютекс"

	for i := 0; i < TaskCount; i++ {
		go func() {
			<-mu //
			counter++

			done <- struct{}{}
			mu <- struct{}{}
		}()
	}

	var readyTasksCounter int
	for range done {
		readyTasksCounter++

		if readyTasksCounter == TaskCount {
			break
		}
	}

	println(readyTasksCounter, time.Now().Sub(t).String())
}

