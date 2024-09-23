package main

import "fmt"

func main() {
	/**
	soal :
	buatlah output
	**********
	**      **
	* *    * *
	*  *  *  *
	*   **   *
	*   **   *
	*  *  *  *
	* *    * *
	**      **
	**********
	dari input angka 10
	*/

	// Mendefinisikan variabel 'kotak' dengan nilai 10
	kotak := 10

	// Loop pertama untuk mengiterasi setiap baris dari 0 hingga 'kotak'
	for i := 0; i <= kotak; i++ {
		// Loop kedua untuk mengiterasi setiap kolom dari 0 hingga 'kotak - 1'
		for j := 0; j < kotak; j++ {
			// Memeriksa apakah kita berada di baris pertama, baris terakhir, kolom pertama,
			// diagonal menurun (j == kotak-i), atau diagonal naik (j == i)
			if i == 0 || i == kotak || j == 0 || j == kotak-i || j == i {
				// Jika syarat terpenuhi, cetak bintang (*)
				fmt.Print("*")
			} else {
				// Jika tidak, cetak spasi untuk membentuk area kosong di dalam pola
				fmt.Print(" ")
			}
		}
		// Mencetak bintang (*) di akhir baris untuk menutup sisi kanan dari pola
		fmt.Println("*")
	}
}
