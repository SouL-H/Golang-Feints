package main

import "fmt"


type Araba struct{
	MotorTip struct{
		Marka string
		Model string
		Tip struct{
			Elektrikli bool
			Hibrit bool
			YakitTipi struct{
				Benzin bool
				LPG bool
				Dizel bool
			}
		}
	}
}

func main(){
a:= Araba{}.MotorTip.Tip.Elektrikli
fmt.Println(a)

//Anonim Struct, sonda {} kullanımın sebebi hem değerde atanbilir otomatik eşleştirir.
//Hem de constructor için kullanılması mecburdur.
 Araba:= struct {
	 Motor struct{
		 Marka string
		 Model string
		 Tip struct{
			 Elektirikli bool
			 Hibrit bool
			 YakitTipi struct{
				 Benzin bool
				 LPG bool
				 Dizel bool
			 }
		 }
	 }
 }{}
 println("---Model ", Araba.Motor.Model)
 StuctExample()
  StructArrAnonim()



	
}


func StuctExample(){
	araba := struct {
		Marka string
		Model string
		Yil int
	}{
		"tofaş","doğan",1995,
	}
	println(araba.Model)
}

func StructArrAnonim(){
	araba:= []struct {
		Marka string
		Yil int
	}{
		{"tofaş",2000},
		{"BMW",2021},
		{"FIAT",2020},
	}
	fmt.Println(araba[0].Marka)
}
