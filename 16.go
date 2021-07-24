package main

import "fmt"

func f16() {
	num := []int{1,2,3,4,5,6,7,8,9,10}

	s1 := num[:1:9]
	fmt.Println(s1, len(s1), cap(s1))

	// вторым значение задается capacity
	// в пределеах длины изначального массива

	// s2 := num[::4] не скомпилируется
	// fmt.Println(s2)
}
