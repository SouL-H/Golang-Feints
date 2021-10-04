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

//Anonim Struct {}


	
}

