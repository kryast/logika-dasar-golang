package main

import "fmt"

func main() {

	// Soal  : buatlah bilangan fibonacci 10 kali

	// Deklarasi slice kosong untuk menampung deret Fibonacci
	var fibonacci []int

	// Inisialisasi dua angka pertama dalam deret Fibonacci
	angka1 := 0 // Bilangan pertama (F(0))
	angka2 := 1 // Bilangan kedua (F(1))

	// Loop untuk menghitung deret Fibonacci dari F(0) hingga F(10)
	for i := 0; i <= 10; i++ {
		// Menangani kasus untuk i = 0 dan i = 1 secara khusus
		if i == 0 || i == 1 {
			// Tambahkan angka i ke dalam slice Fibonacci
			fibonacci = append(fibonacci, i)
		} else {
			// Hitung total sebagai penjumlahan dari dua angka sebelumnya
			total := angka1 + angka2
			// Tambahkan total ke dalam slice Fibonacci
			fibonacci = append(fibonacci, total)
			// Update angka1 dan angka2 untuk iterasi selanjutnya
			angka1 = angka2 // angka1 diupdate ke angka2 (F(n-1))
			angka2 = total  // angka2 diupdate ke total (F(n))
		}
	}
	// Menampilkan hasil deret Fibonacci
	fmt.Println("Bilangan Fibonacci :", fibonacci)
}
