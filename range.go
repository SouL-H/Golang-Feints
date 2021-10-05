package main

//Key-value  çiftini range ile gezebiliriz.
func main() {
	RangeExample()
}

func RangeExample() {
	list := map[string]string{
		"A": "Siyah",
		"B": "Mor",
	}

	for k, v := range list {
		println("Key: ", k, "Value: ", v)
	}
	i := 0
	//Sadece mevcut olup olmadığını kontrol etmek için.
	//Amaç dönen değeri almadan kullanılabilirliği.
	for range list {
		i++
		println(i)
	}
}
