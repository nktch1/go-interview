package main

func ff9(x int) int {
	defer func(xx int) {
		println("defer", x, xx)
		xx++
	}(x)
	x = 2

	println("end")
	return x
}

func f9() {
	r := 0
	println(ff9(r))
}

// defer не может менять значение в return,
// но замыкание может хранить в себе значение входного параметра,
// поэтому если важно хранить состояние, в замыкание нужно передать параметр,
// чтобы иметь именно копию значения, а не ссылку
