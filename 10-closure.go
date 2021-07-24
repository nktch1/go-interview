package main

func closure() func(int)int {
	s := 0 // на стэке аллоцируется переменная,
	// которая потом может изменяться и не удаляться при этом GC

	return func(x int) int {
		s += x
		return s
	}
}

func f10() {
	a, b := closure(), closure()

	println(a(1), b(2), b(3), a(-1))

}