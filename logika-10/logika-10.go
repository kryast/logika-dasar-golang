package main

import "fmt"

func main() {

	/**
	soal :
	buatlah output
	**********
	*        *
	*        *
	*        *
	*        *
	*        *
	*        *
	*        *
	*        *
	**********
	dari input angka 10
	*/

	// Mendefinisikan variabel 'kotak' dengan nilai 10
	kotak := 10

	// Loop pertama untuk mengiterasi setiap baris dari 1 hingga 'kotak'
	for i := 1; i <= kotak; i++ {
		// Loop kedua untuk mengiterasi setiap kolom dari 1 hingga 'kotak'
		for j := 1; j <= kotak; j++ {
			// Memeriksa apakah kita berada di baris pertama, baris terakhir, atau kolom pertama
			if i == 1 || i == kotak || j == 1 {
				// Jika syarat terpenuhi, cetak bintang (*)
				fmt.Print("*")
			} else {
				// Jika tidak, cetak spasi untuk membentuk area kosong di dalam kotak
				fmt.Print(" ")
			}
		}
		// Mencetak bintang (*) di akhir baris untuk menutup sisi kanan dari kotak
		fmt.Println("*")
	}
}
