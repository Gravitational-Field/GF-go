package main

import "fmt"

type Animal interface {
	Eat()
}

type Bird struct {
	Name string
}

type BigBird struct {
	Bird
}

func (b Bird)Eat() {
	fmt.Println(b.Name+"吃虫")
}


type Dog struct {
	Name string
}

func (d Dog)Eat() {
	fmt.Println(d.Name+"吃肉")
}


func EatWhat(a Animal)  {
	a.Eat()
}


func IsDog(a Animal) bool {
	if v, ok := a.(Dog); ok {
		fmt.Println(v)
		return true
	}
	return false
}

func WhatType(a Animal) {
	switch a.(type) {
	case Dog:
		fmt.Println("Dog")
	case Bird:
		fmt.Println("Bird")
	default:
		fmt.Println("error")
	}
}


func main() {
	b := new(Bird)
	d := new(Dog)

	dd := Dog{}

	EatWhat(b)  //吃虫
	EatWhat(d)  //吃肉

	bb := new(BigBird)
	EatWhat(bb) //吃虫

	fmt.Println(IsDog(*d)) //{}   true
	fmt.Println(IsDog(dd)) //{}   true

	WhatType(*b) //Bird
	WhatType(*d)  //Dog
	WhatType(*bb)  //error

	/*空接口*/
	var any interface{}

	any = 1
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)

}
