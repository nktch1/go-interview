package main

import (
	"fmt"
	"strings"
)

const LRUSize = 5

func update(slice []string, str string) []string {
	for idx, el := range slice {
		if el == str {
			copy(slice[idx:], slice[idx+1:])
			slice[len(slice)-1] = el

			return slice
		}
	}

	if len(slice) < LRUSize {
		slice = append(slice, str)
		return slice
	}

	copy(slice[0:], slice[1:])
	slice[len(slice)-1] = str

	return slice
}

func LRUCache(strArr []string) string {

	var (
		store = make(map[string]struct{}) // для проверки наличия строки в кэше
		slice []string                    // кэш
	)

	for _, el := range strArr {
		_, isExists := store[el]

		if !isExists {
			store[el] = struct{}{}
		}

		slice = update(slice, el)
	}

	return strings.Join(slice, "-")

}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(LRUCache([]string{"A", "B", "C", "D", "E", "D", "Q", "Z", "C"}))
	fmt.Println(LRUCache([]string{"A", "B", "A", "C", "A", "B"}))
}
