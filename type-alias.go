package main

type Yazi string
type Sonuc int
func (k Yazi) Yaz() {
	println(k)
}
func (y Sonuc) Hesapla(){
	print("Girilen değerin karekökü:",(y*y))

}
func main() {
	var gelen Yazi
	gelen = "Lorem"
	//Bu şekilde stringe yeni yetenekler kazandırabildik.
	gelen.Yaz()
	var snc Sonuc
	snc=6
	snc.Hesapla()

	// byte uint8'in alias'ıdır.

}
