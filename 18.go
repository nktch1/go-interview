package main

func ff18(n int) func()int {
	return func() int {
		n++
		return n
	}
}

func f18() {
	f := ff18(5)

	defer func() {
		println(f())
	}()

	defer println(f())

	f()
}
