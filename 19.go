package main

func f19() {
	x := []int(nil)
	y := map[int]int(nil)
	_ = x[:]
	_ = y[3]
	z := interface{}(x)
	w := interface{}(x)
	_ = z == w
	f := z == w
	_ = z == z

	println(f)
	//
	//type s struct{}
	//type v struct{}
	//
	//vv := interface{}(s{})
	//vvv := interface{}(v{})
	//
	//println(vv == vvv)
}

// нельзя сравнивать два пустых интерфейса?