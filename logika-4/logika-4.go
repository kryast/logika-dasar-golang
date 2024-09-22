package main

import "fmt"

func main() {

	// Soal 2 : carilah bilangan prima dari 1 - 100

	// Deklarasi slice kosong untuk menampung bilangan prima
	var prima []int

	// Loop dari 1 hingga 100 untuk mencari bilangan prima
	for i := 1; i <= 100; i++ {
		// Variabel untuk menghitung jumlah pembagi
		angka := 0

		// Loop untuk memeriksa setiap angka dari 1 hingga i
		for j := 1; j <= i; j++ {
			// Jika i dapat dibagi j tanpa sisa, berarti j adalah pembagi dari i
			if i%j == 0 {
				angka = angka + 1 // Tambah jumlah pembagi
			}
		}

		// Jika jumlah pembagi tepat 2, maka i adalah bilangan prima
		if angka == 2 {
			// Tambahkan bilangan prima ke dalam slice
			prima = append(prima, i)
		}
	}

	// Tampilkan hasil bilangan prima yang ditemukan
	fmt.Println("Bilangan Prima :", prima)
}
