package main

import "fmt"

type A interface {
	Add() bool
	Del() bool
}
type TA struct {
}

func (TA) Add() bool {
	return true
}

//Bunlardan birini eklemediğimizde maindeki if kontrolüne takılacaktır.
func (TA) Del() bool {
	return false
}

//TA tipine add ve del eklediğimizde implemente etmiş oluyoruz.
func main() {

	var t interface{}
	t = TA{}
	//interface check ettiğimiz yer.
	v, ok := t.(A)

	if !ok {
		println("Interface implemente edilmemiş")
	}
	fmt.Printf("%T\n", v)
}
