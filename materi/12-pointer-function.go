package materi

import "fmt"

func ubahNilai(nilaiAwal *int) {
	*nilaiAwal = 100
}

func MainUbahNilai() {
	nilai := 10
	// pNilai := &nilai
	// ubahNilai(pNilai)
	ubahNilai(&nilai)
	fmt.Println(nilai)
}