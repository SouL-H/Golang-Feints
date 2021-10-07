package main

import ("fmt"
		"github.com/jinzhu/copier")

func main() {
	fmt.Println("Hello")
	personList:= []Person{
		{Name:"Bora",Surname: "Kasmer",Age:43},
		{Name:"Bill",Surname: "Gates",Age:55},
		{Name:"Micheal",Surname: "Jordan",Age:52},
		{Name:"Marc",Surname: "Zekuber",Age:37},
		{Name:"Burak",Surname: "Senyurt",Age:45},
		{Name:"Engin",Surname: "Polat",Age:41},

	}

	famousList:=[]Person{}

	copier.Copy(&famousList,personList[1:4])
	famousList[0].Name="Ataturk"
	famousList[1].Name="Ataturk"
	famousList[2].Name="Ataturk"

	for _, per := range famousList{
		fmt.Printf("Famous Name:%s / Surname: %s / Age: %d \n",per.Name,per.Surname,per.Age)
	}

}

func justDoIt() (string, error) {
	return "Fonksiyon değeri", fmt.Errorf("Fonksiyonda hata var.")
}

type Person struct{
	Name string
	Surname string
	Age int
}
