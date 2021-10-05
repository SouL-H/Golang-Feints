package main

//Parametre olarak fonksiyon alan fonksiyonlara higher-Order fonksiyon denir.

func ArabaBoya(fn func(string),renk string){
	fn(renk)
}

func main(){
f:=func(renk string){
	println("Arabanın rengi",renk)
}

ArabaBoya(f,"mavi")

}