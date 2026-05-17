package main

import "fmt"

func main() {
	// Mendeklarasikan array dengan 6 elemen
	// Index:  0   1   2   3   4   5
	// Value: [10, 20, 30, 40, 50, 60]
	arr := [6]int{10, 20, 30, 40, 50, 60}

	// 1. Mengambil seluruh elemen array
	// Syntax [:] berarti dari index awal sampai index paling akhir.
	s1 := arr[:] // ini untuk mengambil seluruh element
	fmt.Println("ini adalah S1 :", s1) // Output: [10 20 30 40 50 60]

	// 2. Mengambil elemen dari awal HINGGA sebelum index 3
	// Syntax [:batas_akhir] berarti mulai dari index 0 sampai batas_akhir - 1.
	// Di sini mengambil index 0, 1, dan 2. (Index 3 tidak diikutkan)
	s2 := arr[:3]
	fmt.Println("ini adalah s2 :", s2) // Output: [10 20 30]

	// 3. Mengambil elemen dari index 2 hingga paling akhir
	// Syntax [batas_awal:] berarti mulai dari batas_awal sampai elemen terakhir.
	// Di sini mengambil index 2, 3, 4, 5.
	s3 := arr[2:]
	fmt.Println("ini adalah S3 :", s3) // Output: [30 40 50 60]

	// 4. Mengambil elemen dari index 1 HINGGA sebelum index 4
	// Syntax [batas_awal:batas_akhir] 
	// Di sini mengambil index 1, 2, dan 3. (Index 4 tidak diikutkan)
	s4 := arr[1:4]
	fmt.Println("ini adalah s4 :", s4) // Output: [20 30 40]
}