package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Mengubah Boolean menjadi String (bool -> string)
	truth := true
	// FormatBool mengubah nilai true/false menjadi teks "true" atau "false"
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean ke string:", str)

	// 2. Mengubah String menjadi Boolean (string -> bool)
	// ParseBool mencoba membaca teks dan mengubahnya jadi nilai boolean.
	// Teks yang valid misalnya: "1", "t", "T", "true", "TRUE", "True" (untuk true)
	// atau "0", "f", "F", "false", "FALSE", "False" (untuk false).
	// 
	// CATATAN PENTING: Sama seperti Atoi sebelumnya, ParseBool mengembalikan 2 nilai (hasil dan error).
	// Karena kita sedang tidak ingin mengurus pesan error-nya, kita gunakan tanda garis bawah (_) 
	// atau "Blank Identifier" untuk membuang/mengabaikan nilai error tersebut agar Go tidak protes.
	val, _ := strconv.ParseBool("true")
	fmt.Println("String ke boolean:", val)
}