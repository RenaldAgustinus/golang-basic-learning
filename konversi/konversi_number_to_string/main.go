package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Mendeklarasikan variabel 'text' berupa teks (string) yang berisi angka "100"
	var text string = "100a"

	// Mengubah teks menjadi angka menggunakan fungsi Atoi (ASCII to Integer).
	// Fungsi ini mengembalikan DUA nilai sekaligus: 
	// 1. 'number' = hasil konversinya (jika berhasil)
	// 2. 'err' = pesan error (jika teks tidak bisa diubah jadi angka, misal teksnya "seratus")
	number, err := strconv.Atoi(text) // string to int

	// Pengecekan Error (Sangat umum di Golang)
	// Jika 'err' tidak sama dengan nil (artinya ada error yang terjadi)
	if err != nil {
		fmt.Println("Pesan Error : ", err.Error())
	} else {
		// Jika 'err' bernilai nil (artinya konversi sukses tanpa masalah)
		fmt.Println("Angka:", number)
	}
}