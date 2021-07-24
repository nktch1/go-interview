package main

import "fmt"

func f3() {
	s1 := make([]int, 1, 1) // слайс создан
	s2 := s1 // ссылка на слайс

	s1 = append(s1, 1)
	s2 = append(s2, 2)

	// если capacity больше чем количество элементов в результирующем массиве,
	// то новой аллокаци не происходит и состояние сохраняется

	fmt.Println(s1, s2)
}