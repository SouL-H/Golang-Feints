package main

import "fmt"

func main() {
	if v, err := justDoIt(); err == nil {
		fmt.Println("Fonksiyonda hata yok ve dönen değer:", v)
	}else{
			fmt.Println(err)
	}

}

func justDoIt() (string, error) {
	return "Fonksiyon değeri", fmt.Errorf("Fonksiyonda hata var.")
}
