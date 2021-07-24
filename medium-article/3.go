package main

import "time"

func worker(res chan int){
	//ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		res <- 42
	}()
}

func f3() {
	timeStart := time.Now()
	ch := make(chan int)

	worker(ch)
	worker(ch)

	//_, _ = <-worker(ch), <-worker(ch)
	_, _ = <-ch, <-ch

	println(int(time.Since(timeStart).Seconds())) // что выведет - 3 или 6?
}
