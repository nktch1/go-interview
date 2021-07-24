package main

func ff1() {
	defer println("f1 start")
	ff2()
	defer println("f1 finish")
}

func ff2() {
	defer println("f2 start")
	ff3()
	defer println("f2 finish")
}

func ff3() {
	defer println("f3 start")
	panic(0)
	defer println("f3 finish")
}

func f8() {
	ff1() // стэк разворачивается в обратную сторону, функции не завершилиись корректно
	// -> в обратном порядке будут выведены сообщения о старте функции (f3 started -> f2 started -> f1 started)
}
