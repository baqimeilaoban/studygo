package main

//接口的实现

type animal interface {
	move()
	eat()
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	name string
}
