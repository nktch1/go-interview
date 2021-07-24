package main

import (
	"fmt"
	"runtime"
	"time"
)

func f7() {
	t := time.Now()
	numcpu := runtime.NumCPU()
	println(numcpu)

	//runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(numcpu)

	var (
		ch1 = make(chan int)
		ch2 = make(chan float64)
	)

	go func() {
		for i := 0; i < 1000000; i++ {
			ch1 <- i
		}

		//ch1 <- -1
		ch2 <- 0.0
	}()

	go func() {
		total := 0.0

		for {
			t1 := time.Now().UnixNano()
			for i := 0; i < 100000; i++ {
				m := <-ch1

				if m == -1 {
					ch2 <- total
				}
			}

			t2 := time.Now().UnixNano()
			dt := float64(t2 - t1) / 1000000.0
			total += dt

			println(dt)
		}
	}()

	fmt.Println("Total", <-ch1, <-ch2, time.Now().Sub(t) )

	// стэк горутины 8K (а у классических тредов - 160KB in 32-bit mode and 320KB in 64-bit mode)
	// runtime.GOMAXPROCS
}