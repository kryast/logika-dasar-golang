package main

import "fmt"

func main() {

	/**
	soal :
	buatlah output
	**********
	**********
	**********
	**********
	**********
	**********
	**********
	**********
	**********
	**********
	dari input angka 5
	*/

	kotak := 10

	for i := 0; i <= kotak; i++ {
		for j := 0; j <= kotak; j++ {
			fmt.Print("*")
		}
		fmt.Println("*")
	}
}
