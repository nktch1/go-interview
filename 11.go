package main

func f11() {
	var i interface{}

	if i == nil {
		println("nil")
		return
	}

	println("not nil")
}
