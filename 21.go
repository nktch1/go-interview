package main

import (
	"fmt"
	"sort"
)

type Ints []int

func rm(slice []int, idx int) []int {
	copy(slice[idx:], slice[idx+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}

func ins(slice []int, value, index int) []int {
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func f21() {
	s := []int{1,2,3,4,5,6}

	fmt.Println(s)
	sort.Sort(Ints(s))
	fmt.Println(s)

	ss := rm(s, 4)
	fmt.Println(ss)

	sss := ins(ss, 2, 4)
	fmt.Println(sss)
}

func (in Ints) Len() int { return len(in) }
func (in Ints) Less(i, j int) bool { return in[i] > in[j] }
func (in Ints) Swap(i, j int) { in[i], in[j] = in[j], in[i] }


