package main

import "fmt"

func main() {
	// Mendeklarasikan variabel 'a' dengan tipe data integer (bilangan bulat) dan nilai 10.
	var a int = 10
	
	// Melakukan konversi tipe data (Type Casting).
	// Nilai dari variabel 'a' (int) dikonversi menjadi tipe data float64 (bilangan desimal/pecahan),
	// kemudian hasilnya disimpan ke dalam variabel baru bernama 'b'.
	var b float64 = float64(a) // ubah dari int ke float64

	// Menampilkan hasil ke layar/terminal.
	fmt.Println("Nilai a:", a)
	fmt.Println("Nilai b:", b)
}