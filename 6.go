package main

type S6 struct {}

func ff6(x *interface{}) {}

func f6() {
	//ff6(S{}) не заработает разумеется, поинтеры на интерфейсы это гиблое дело
}
