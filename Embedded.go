package main

import (
	"fmt"
	"sync"
)

//Embedded:isim vermeden bir tipi diğer bir tipin içine ekleme.

type Arac struct {
	marka string
}

func (a Arac) String() string {
	return fmt.Sprintf("Araç tipindeki string fonksiyon: %v", a.marka)
}

//Aracın tüm özellikleri aktarılır.
type Araba struct {
	Arac
}

//println ile basmak istdiğimizde String() fonksiyonu çalışacaktır.
//Nativ interface olduğu için otomatikman Araba ekrana yazdırmada String interfacesini çağırır.
func main() {
	a := Araba{}
	a.marka = "tofaş"
	fmt.Println(a)
}

//Go dilinde klaik katılım ve polimorfik hareketler klasik OOP gibi çalışmamaktaı.
//Onun yerine bir tipi başka bir tipin içine aktararak kullanabiliyoruz.
//var olarak atamak önemli type problem yaratıyor.
var Sayac struct {
	sync.Mutex
	Sayi int
}
func EmpleddedLockExample(){
	Sayac.Lock()
	Sayac.Sayi++
	Sayac.Unlock()
}
