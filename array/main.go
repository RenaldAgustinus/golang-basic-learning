package main

import "fmt"

func main() {
	// Mendeklarasikan array bernama 'angka' dengan panjang/kapasitas persis 3 elemen.
	// Tipe datanya adalah integer (int), dan langsung diisi dengan nilai: 10, 20, 30.
	var angka [3]int = [3]int{10, 20, 30}
	
	// Menampilkan seluruh isi array sekaligus
	fmt.Println(angka)      // [10 20 30]
	
	// Menampilkan elemen array pada indeks ke-1.
	// CATATAN PENTING: Indeks array selalu dimulai dari 0, bukan 1.
	// Indeks 0 nilainya 10
	// Indeks 1 nilainya 20
	// Indeks 2 nilainya 30
	fmt.Println(angka[1]) // 20
	
	// Mengubah isi array pada indeks ke-1.
	// Nilai yang sebelumnya 20 ditimpa/diganti menjadi 80.
	angka[1] = 80
	
	// Menampilkan kembali nilai pada indeks ke-1 untuk melihat perubahannya
	fmt.Println("Ini adalah value index ke-1 :", angka[1])
}