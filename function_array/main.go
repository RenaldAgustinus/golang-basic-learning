package main

import "fmt"

func main() {
	// --- BAGIAN 1: Fungsi len() dan Perulangan (Looping) Array ---
	number := [5]int{10, 20, 30, 40, 50}

	// len() digunakan untuk mendapatkan panjang atau jumlah kapasitas elemen array
	fmt.Println("Jumlah elemen : ", len(number))
	
	fmt.Println("Index ke-1 :", number[1])
	number[1] = 80 // Mengubah nilai pada index ke-1
	fmt.Println("Index ke-1 :", number[1])

	fmt.Println("Ini adalah array :")

	// 'for range' adalah cara paling elegan di Golang untuk mengulang (iterasi) isi array.
	// Pada setiap putaran, ia otomatis memberikan 2 nilai: 'index' (posisi) dan 'value' (isinya).
	for index, value := range number {
		fmt.Println("isi index ke :", index, " = ", value)
	}

	fmt.Println() // Hanya untuk memberi jarak enter (baris kosong) di terminal

	// --- BAGIAN 2: Perbandingan Array ---
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	// Di Golang, kita bisa langsung membandingkan dua array menggunakan == atau !=
	// Syarat utamanya: Tipe data dan kapasitas (panjang) kedua array harus sama persis!
	// Array dianggap sama (==) JIKA DAN HANYA JIKA semua elemen di index yang sama nilainya sama.
	fmt.Println("arr1 == arr2 :", arr1 == arr2) // Pasti false, karena isinya (1,2,3) beda dengan (4,5,6)
	fmt.Println("arr1 != arr2 :", arr1 != arr2) // Pasti true, karena memang isinya tidak sama
}