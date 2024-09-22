package main

import "fmt"

func main() {

	/** buatlah output
	5
	45
	345
	2345
	12345
	dari input 5
	*/

	var input = 5 // masukkan input 5 terlebih dahulu

	for i := 1; i <= input; i++ { // buat loop pertama sebagai pembatas dengan output maksimal angka 5
		for j := input - i + 1; j <= input; j++ { // buat loop kedua di dalam loop pertama untuk mengisi angka
			fmt.Print(j) // print digunakan untuk output nya ke kanan
		}
		fmt.Println() // println digunakan untuk output nya ke bawah
	}

	/** buatlah output
	1
	12
	123
	1234
	12345
	dari input 5
	*/

	var input2 = 5
	for a := 1; a <= input2; a++ {
		for b := 1; b <= a; b++ {
			fmt.Print(b)
		}
		fmt.Println()
	}
}
