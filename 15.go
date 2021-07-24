package main

func ff15(num ...int) {
	num[0] = 18
}

func f15() {

	// 1. func Sum(nums ...int) int {
	// 2. fmt.Println(Sum(primes...))
	// 3. [...]string{"Moe", "Larry", "Curly"} - длина равная количеству элементов при инициализации
	// 4. go test ./...

	i := []int{5,6,7}
	ff15(i...)
	println(i[0])
}