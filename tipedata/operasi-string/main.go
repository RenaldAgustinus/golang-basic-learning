package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Halo, Dunia!"

	// Mengubah menjadi huruf kecil
	fmt.Println("Lowercase:", strings.ToLower(text)) // halo, dunia!

	// Mengubah menjadi huruf besar
	fmt.Println("Uppercase:", strings.ToUpper(text)) // HALO, DUNIA!

	// Mengecek apakah string dimulai dengan "Halo"
	fmt.Println("StartsWith Halo?", strings.HasPrefix(text, "Halo")) // true

	// Mengecek apakah mengandung kata "Dunia"
	fmt.Println("Contains Dunia?", strings.Contains(text, "Dunia")) // true

	// Memisahkan string berdasarkan delimiter
	parts := strings.Split("apel,banana,cerry", ",")
	fmt.Println("Split:", parts) // [apel banana cerry]

	// Mengganti bagian string
	newText := strings.ReplaceAll(text, "Dunia", "Golang")
	fmt.Println("Replace:", newText) // Halo, Golang!
}