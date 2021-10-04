package main

import "fmt"

func main() {
	mapShot()
}

//map'da mevcut değer olup olmadğını kolayca öğrenebiliriz.
func mapDeger() {
	list := map[string]int{
		"A": 1,
		"B": 2,
	}

	v, ok := list["A"]

	if ok {
		fmt.Println("Mevcut değer map'da bulunuyor, Değer: ", v)
	}
}
//Anonim 
func mapShot() {
	arabalar := map[int]struct {
		Marka string
		Model string
		Yil   int
	}{
		1: {"tofaş", "doğan", 1995},
		2: {"tofaş", "şahin", 2000},
	}
	fmt.Printf("İçeriği: %v\n", arabalar)
}
