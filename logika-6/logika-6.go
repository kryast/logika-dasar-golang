package main

import "fmt"

func main() {

	/** Soal : Ahmad menabung dengan uang 750.000 di bulan pertama dan mendapatkan
	bunga sebesar 6% di bulan berikutnya.
	setiap bulan berikutnya, Ahmad menyisihkan 3% dari tabungannya
	dan maksimal yang disisihkan nya 3% dari 800.000.
	berapakah total tabungan Ahmad selama 1 tahun ?
	*/

	// Inisialisasi jumlah uang Ahmad
	uangAhmad := 750000

	// Loop untuk menghitung saldo tabungan Ahmad selama 12 bulan
	for i := 1; i <= 12; i++ {
		// Inisialisasi variabel bunga dan sisih untuk setiap bulan
		bunga := 0
		sisih := 0

		// Perhitungan hanya dimulai dari bulan kedua (i > 1)
		if i > 1 {
			// Hitung bunga sebagai 6% dari uangAhmad
			bunga = uangAhmad * 6 / 100

			// Memeriksa apakah total uang setelah bunga lebih dari 800000
			if (uangAhmad + bunga) > 800000 {
				// Jika lebih dari 800000, sisihkan 3% dari 800000
				sisih = 800000 * 3 / 100
			} else {
				// Jika tidak, sisihkan 3% dari total uang setelah bunga
				sisih = (uangAhmad + bunga) * 3 / 100
			}
		}

		// Update total uang Ahmad setelah penambahan bunga dan pengurangan sisih
		uangAhmad = (uangAhmad + bunga) - sisih

		// Menampilkan informasi bulanan: saldo tabungan, bunga, sisih, dan uang bulan ini
		fmt.Println("Tabungan awal :", uangAhmad, "\n Bunga :", bunga, "\n Sisih :", sisih, "\n Uang bulan ke-", i, ": ", (uangAhmad+bunga)-sisih)
	}
}
