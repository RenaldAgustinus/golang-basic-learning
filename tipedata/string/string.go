package main

import "fmt"

func main() {
	// Deklarasi string dengan kutip ganda
	nama := "Triady"
	pesan := "Selamat datang di aplikasi kami!"

	// Deklarasi string dengan backtick (raw string)
	paragraf := `Halo, ini adalah contoh multi-line
        string di Go.`

	// Menampilkan nilai string
	fmt.Println("Nama:", nama)
	fmt.Println("Pesan:", pesan)
	fmt.Println("Paragraf:\n", paragraf)
}