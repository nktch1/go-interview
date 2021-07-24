package main

import "fmt"

var o = fmt.Print

type A int
func (A) g() {}
func (A) m() int { return 1 }

type B int
func (B) g() {}
func (B) f() {}

type C struct {
	A
	B
}

func (C) m() int { return 9 }

type Q int
func (Q) ff() {}

func f22() {
	var c interface{} = C{}

	_, bf := c.(interface{ff()})
	_, bg := c.(interface{g()})

	i := c.(interface{m() int})

	fmt.Println(bf, bg, i.m())


}
