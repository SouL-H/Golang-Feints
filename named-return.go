package main

func Araba() (marka string, model int) {
	return "BMW", 1995
}
func ArabaV2() (marka string, model int) {
	marka = "Tofaş"
	model = 2010
	return
}
func main() {
	v1, v2 := Araba()
	println("Deger 1:,", v1, "Deger 2: ", v2)
	v3, v4 := ArabaV2()
	println("Deger 3:,", v3, "Deger 4: ", v4)
}
