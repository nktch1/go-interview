package main

type I1 interface {
	Read() int
}

type I2 interface {
	Read() int
}

type S1 struct {
	//m string
}

func (s *S1) Read() int { return 0 }

type S2 struct{}

func (s *S2) Read() int { return 0 }

func f12() {

	var (
		s1  = &S1{}
		s2  = S2{}
		s11 = S1(s2)
	)

	var (
		a1 I1 = s1
		a2 I1 = &s11
	)

	println(a1 == a2)

	// ресивер должен соотвествовать типу созданной переменной (если "func (s *Str) ..."
	// в реализации интерфейсов, то s := &Str{})
}
