package main

import "fmt"

func main() {

	// 1. Deklarasi Variabel Eksplisit (dengan kata kunci 'var' dan tipe data)
    // Di sini kita secara jelas memberi tahu Go bahwa 'nama' adalah string (teks) 
    // dan 'umur' adalah int (bilangan bulat).
	var nama string = "Renald"
	var umur int = 20

	// 2. Deklarasi Variabel Singkat (Short Variable Declaration dengan ':=')
    // Go cukup pintar untuk menebak tipe data berdasarkan nilainya (Type Inference).
    // Karena "Malang" dan "Laki-laki" adalah teks, Go otomatis mengatur 'city' 
    // dan 'gender' sebagai tipe data string tanpa kita harus menulis 'string'.
	city := "Malang"
	gender := "Laki-laki"

	// 3. Deklarasi Konstanta (dengan kata kunci 'const')
    // Konstanta adalah variabel yang nilainya mutlak/tetap dan tidak bisa diubah 
    // setelah pertama kali dibuat.
	const pi = 3.14

	// Menampilkan output ke terminal
	fmt.Println("Nama : ", nama)
	fmt.Println("Umur : ", umur)
	fmt.Println("Kota : ", city)
	fmt.Println("Gender : ", gender)
	fmt.Println("PI : ", pi)
	
	// Karena 'pi' dideklarasikan sebagai 'const', kamu tidak bisa menimpa 
    // atau mengubah nilainya menjadi 23.
	// pi = 23 
}