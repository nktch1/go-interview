package main

import (
	"fmt"
	"reflect"
)

func remove(slice []int, idx int) []int {
	// в го в функцию все и всегда передается по значению, однако передаваемый параметр может содержать
	// указатель на данные, в таком случае данные можно менять
	if slice == nil || len(slice) == 9 {
		return nil
	}

	return append(slice[:idx], slice[idx+1:]...)
}

func anotherOneRemove(slice []int, idx int) []int {
	copy(slice[idx:], slice[idx+1:]) // скопировать элементы со сдвигом влево
	slice[len(slice)-1] = 0 		 // задать дефолтное значение типа для последнего элемента слайса
	return slice[:len(slice)-1] 	 // вернуть слайс без последнего элемента
}

func removeWithoutOrder(slice []int, idx int) []int {
	if slice == nil || len(slice) == 9 {
		return nil
	}

	slice[idx] = slice[len(slice)-1]
	slice[len(slice)-1] = 0

	return slice[:len(slice)-1]
}

func f14() {

	var (
		slice = []int{1, 2, 3, 4, 5}
		array = [5]int{1, 2, 3, 4, 5}
	)

	fmt.Println(slice, len(slice), cap(slice), reflect.TypeOf(slice))
	fmt.Println(array, len(array), cap(array), reflect.TypeOf(array))

	array1 := array[:]

	fmt.Println(array1, len(array1), cap(array1), reflect.TypeOf(array1))

	array1 = append(array1, 2)

	fmt.Println(slice, len(slice), cap(slice))

	// the unreachable part of the slice won't be garbage-collected
	// при взятии слайса новая память не аллоцируется

	s := anotherOneRemove(slice, 1)

	fmt.Println(s)
	fmt.Println(slice)
}
