package main

import "fmt"

func f23() {
	a := []int{0,1,2,3,4,5,6,7}

	//fmt.Println(a)
	//var s = make([]int, 10)
	//copy(s, a[:])
	//
	//a[0] = 1126
	//a = append(a, 1126)
	//
	//
	//fmt.Println("copy ", s)
	//fmt.Println(a)
	//
	//s0 := a[1:]
	//fmt.Println(s0, a)
	//s1 := s0[:2:4]
	//fmt.Println(s1, a)
	//s2 := append(s1, s0[4:]...)
	//fmt.Println(s2, a)
	//s2 = append(s2, 7)
	//fmt.Println(s2, a)

	fmt.Println(cutIfPointers(a, 4))
}

func rm23(slice []int, idx int) []int {
	return append(slice[:idx], slice[idx+1:]...)
}

func cutIfPointers(slice []int, idx int) []int {
	fmt.Println(slice, slice[idx:], slice[idx+1:])
	copy(slice[1:], slice[1+1:])
	//slice[idx] = -1000
	s := slice[1:]
	fmt.Println(s)
	fmt.Println(s[:cap(s)])



	return slice
}
