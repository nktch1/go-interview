package main

import "reflect"

type T struct {
	x string
	X string
}

func (T) M(){}
func (T) m(){}

func f17() {
	v := reflect.ValueOf(T{})
	println(v.NumField(), v.NumMethod())

	// рефлексия может видеть приватные поля,
	// но не может видеть приватные методы
}
