package main

import "fmt"

type S5 struct {
	m string
}

func ff5() *S5 {
	return &S5{m: "hello"}
}

func f5() {
	p := *ff5()
	fmt.Println(p)
}
