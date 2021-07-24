package main

func changePtr(p *int) {
	v := 3
	p = &v // в этой строке
	println(p)
	println(*p)
}

func f25() {
	//v := 5
	//p := &v
	//println(*p)
	//
	//changePtr(p)
	//println(p)
	//println(*p)

	v1 := 1
	v2 := 2

	p1 := &v1
	p2 := &v2

	p1 = p2

	println(*p1, *p2)

}


