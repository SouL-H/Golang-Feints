package main

func main() {

Level1:
	for i := 0; i < 10; i++ {
		for x := 0; x < 10; x++ {
			if x%3 > 0 {
				break Level1
			}
		}
		if i == 10 {
			goto Level2
		}
	}
Level2:
	println("10 Level oldun oyun bitti.")
}
