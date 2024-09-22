package main

import "fmt"

func main() {

	/**
	Soal : Array[1,5,2,4,5,6,7,2,3,5,1,9,6,8,4,6,8,4,2,5]

	pertannyaan :
	1. tampilkan angka secara tunggal dan total angka apabila di tunggalkan.
	2. tampilkan total dari masing-masing angka
	*/

	// Mendeklarasikan dan menginisialisasi array dengan elemen yang diberikan
	dataArray := [20]int{1, 5, 2, 4, 5, 6, 7, 2, 3, 5, 1, 9, 6, 8, 4, 6, 8, 4, 2, 5}

	// Mendeklarasikan slice kosong untuk menyimpan elemen unik
	var arrayBaru []int

	// Loop untuk mengiterasi setiap elemen dalam dataArray
	for i := 0; i < len(dataArray); i++ {
		// Variabel untuk menandakan apakah elemen sudah ditemukan
		var ketemu = false

		// Loop untuk memeriksa apakah elemen sudah ada dalam arrayBaru
		for a := 0; a < len(arrayBaru); a++ {
			// Jika elemen di dataArray sama dengan elemen di arrayBaru
			if dataArray[i] == arrayBaru[a] {
				// Tandai bahwa elemen sudah ditemukan
				ketemu = true
			}
		}

		// Jika elemen belum ditemukan dalam arrayBaru
		if !ketemu {
			// Tambahkan elemen ke arrayBaru
			arrayBaru = append(arrayBaru, dataArray[i])

			// Variabel untuk menghitung jumlah kemunculan elemen
			totalKetemu := 0

			// Loop untuk menghitung berapa kali elemen muncul dalam dataArray
			for j := 0; j < len(dataArray); j++ {
				// Jika elemen di dataArray sama dengan elemen yang sedang diproses
				if dataArray[i] == dataArray[j] {
					// Tambah jumlah kemunculan
					totalKetemu = totalKetemu + 1
				}
			}

			// Menampilkan jumlah kemunculan elemen
			fmt.Println("Angka ", dataArray[i], "sebanyak ", totalKetemu)
		}
	}
}
