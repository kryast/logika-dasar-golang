package main

import "fmt"

func main() {
	/** buatlah output
	1
	22
	333
	4444
	55555
	dari input 5
	*/

	var input = 5

	for i := 1; i <= input; i++ { // Loop untuk setiap baris
		for j := 1; j <= i; j++ { // Loop untuk mencetak angka sesuai dengan nomor baris
			fmt.Print(i) // Mencetak angka baris yang sama
		}
		fmt.Println() // Pindah ke baris berikutnya
	}
}
