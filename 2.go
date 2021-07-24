package main

import "log"

func f2() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%+v", r)
		}
	}()

	panic(1)
	panic(2)
}
