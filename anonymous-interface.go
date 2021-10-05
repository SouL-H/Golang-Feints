package main

import "fmt"

//Başka bir paketin içindeki private fonksiyonu çağırmak için kullanılabilir.

type Arac struct {
	marka string
}

func (a Arac) String() string {
	return fmt.Sprintf("Araç tipindeki string fonksiyonu: %v", a.marka)
}
func (a Arac) yil() int {
	return 2015
}

type Araba struct {
	Arac
}

func main() {
	a := Araba{}
	a.marka = "BMW"

	//Anonim interfaceye a değerini attığımız için yil değişkenine ulaşabiliyoruz.
	var s interface {
		yil() int
	} = a

	fmt.Println(s.yil())

}
