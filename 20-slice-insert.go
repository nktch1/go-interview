package main

import "fmt"

func insert(slice []int, value, index int) []int {
	slice = append(slice[:index+1], slice[index:]...) // при аппенде взять и задублировать элемент
	// по искомому индексу, а потом просто переписать дубль значением из параметра

	slice[index] = value
	return slice
}

func f20() {
	s := []int{1,3,4}
	fmt.Println(s)

	sss := insert(s, 50, 2)
	fmt.Println(sss)
}