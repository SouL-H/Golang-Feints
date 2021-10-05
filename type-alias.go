package main

type Yazi string

func (k Yazi) Yaz() {
	println(k)
}
func main() {
	var gelen Yazi
	gelen = "Lorem"
	//Bu şekilde stringe yeni yetenekler kazandırabildik.
	gelen.Yaz()
	// byte uint8'in alias'ıdır. 

}
