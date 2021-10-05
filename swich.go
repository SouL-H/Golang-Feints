package main

import (
	"fmt"
	"time"
)


func main(){
	switchExampleV1()
	println(switchExampleV2(351))
}


func switchExampleV1(){
	switch gun:=time.Now().Day();{
	case gun<10:
		fmt.Println("Ayın 1.çeyreği")
	case gun>10 && gun<20:
		println("Ayın ortasi")
	case gun >20:
		println("Ayin sonu")
	default:
		println("Geçersiz gün")
	}
}
//rune dedimal anlamında kullanılır.
//küçük ş'nin Ascii değeri 351'dir.
func switchExampleV2(c rune) bool{
	switch c {
	case 'ş','ü','ğ','ö','ç':
		return true
	}
	return false

}