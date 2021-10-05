package main

func main() {
	var fn func(string)
	//Metodu allocate ettik.
	fn = func(m string) { println(m) }

	fn("Yazdir.")
}
