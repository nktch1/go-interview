package main

import "fmt"

type Stack []rune

func (s *Stack) Push(val rune) {
	*s = append(*s, val)
}

func (s *Stack) Pop() rune {
	popped := (*s)[len(*s)-1]

	*s = (*s)[:len(*s)-1]

	return popped
}

func (s *Stack) Peak() rune {
	return (*s)[len(*s)-1]
}

func (s *Stack) Print() {
	for _, el := range *s {
		fmt.Print(el, " ")
	}

	fmt.Println()
}
