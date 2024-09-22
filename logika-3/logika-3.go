package main

import "fmt"

func main() {

	// Soal : pisahkan angka Ganjil dan Genap dari angka 1 sd 100

	var Genap []int  // buat array untuk menampung angka genap
	var Ganjil []int // buat array untuk menampung angka ganjil

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			Genap = append(Genap, i) // append(array, nilai yg akan di masukkan ke array)
		} else {
			Ganjil = append(Ganjil, i)
		}
	}

	fmt.Println("angka genap :", Genap)
	fmt.Println("angka ganjil :", Ganjil)
}
