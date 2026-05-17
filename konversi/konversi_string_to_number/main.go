package main

import (
	"fmt"
	"strconv" // Package ini wajib di-import untuk melakukan konversi dari/ke string
)

func main() {
	// Mendeklarasikan variabel 'score' berupa angka bulat (integer) bernilai 95
	var score int = 95

	// Mengubah angka (int) menjadi teks (string) menggunakan fungsi Itoa (Integer to ASCII).
	// Di Golang, kita tidak bisa langsung menggunakan string(score) karena hasilnya akan 
	// menjadi karakter Unicode, bukan angka "95" dalam bentuk teks.
	var text string = strconv.Itoa(score) // int to string

	// Menampilkan hasil teks ke layar
	fmt.Println("Nilai ujian:", text)
}