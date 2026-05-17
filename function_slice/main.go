package main

import "fmt"

func main() {
	// --- BAGIAN 1: Membuat Slice dengan fungsi make() ---
	// make([]tipe_data, length, capacity)
	// Membuat slice dengan panjang (len) 3, tapi memiliki kapasitas maksimal (cap) 5.
	// Karena tipenya 'int', 3 elemen pertama akan otomatis diisi dengan angka 0.
	s := make([]int, 3, 5)
	fmt.Println(s) // Output: [0 0 0]
	fmt.Println("Len :", len(s)) // Output: 3
	fmt.Println("Cap :", cap(s)) // Output: 5

	// --- BAGIAN 2: Menambahkan elemen dengan append() ---
	// append digunakan untuk memasukkan data baru ke bagian akhir slice.
	// Karena kapasitas 's' tadi adalah 5 dan isinya baru 3, kita bisa menambahkan 2 data lagi
	// tanpa mengubah ukuran memori di belakang layar.
	s = append(s, 10, 20)
	fmt.Println("Setelah di append :", s) // Output: [0 0 0 10 20]

	// --- BAGIAN 3: Menyalin isi Slice dengan copy() ---
	// Membuat slice baru 'slice2' dengan panjang 3
	slice2 := make([]int, 3)
	
	// copy(tujuan, sumber)
	// Fungsi ini menyalin data dari 's' ke 'slice2'.
	// Karena 'slice2' hanya punya panjang 3, maka ia HANYA BISA menyalin 3 elemen pertama dari 's' (yaitu 0, 0, 0).
	// Nilai kembalian fungsi copy adalah jumlah elemen yang berhasil disalin.
	slice3 := copy(slice2, s)

	fmt.Println("copied ", slice2) // Output: [0 0 0]
	fmt.Println("jumlah elemen yang tersalin :", slice3) // Output: 3

	// --- BAGIAN 4: Membuat Slice langsung dan memotongnya ---
	// Perhatikan: Tanda [] yang kosong (tanpa angka di dalamnya) berarti ini adalah SLICE, bukan Array.
	angka := []int{1, 2, 3, 4, 5}
	
	// Memotong slice dari index 1 hingga sebelum index 4
	slice4 := angka[1:4]
	fmt.Println("Slice 4 :", slice4) // Output: [2 3 4]
}