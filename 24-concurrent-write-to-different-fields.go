package main

import (
	"fmt"
	"sync"
	"time"
)

func f24() {
	type s struct {
		slice1 []int
		slice2 []int
	}

	var (
		oS = s{}
		wg = &sync.WaitGroup{}
		//mu = &sync.Mutex{}
	)

	wg.Add(2)
	for i := 0; i <2; i++ {
		go func(o *s, idx int) {
			defer wg.Done()

			for j := 0; j < 100; j++ {

				//mu.Lock()
				if idx % 2 == 0 {
					o.slice1 = append(o.slice1, j)
				} else {
					o.slice2 = append(o.slice2, j)
				}
				//mu.Unlock()

				time.Sleep(100 * time.Millisecond)
			}
		}(&oS, i)
	}

	wg.Wait()

	fmt.Println(oS.slice1)
	fmt.Println(oS.slice2)
}