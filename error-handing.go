package main

func main() {
	execute()
}

//recover ile erroru catch ediyoruz. Erroru panic ile fırlatmak lazım.
func execute() {
	defer func() {
		if e := recover(); e != nil {
			println("hata yakalandı:", e)
		}
	}()

	panic("exception fırlat")
}
