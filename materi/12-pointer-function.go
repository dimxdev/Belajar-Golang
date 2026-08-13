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


type Food struct {
	nama string 
	rate int
}

func ubahNamaFood(f *Food) {
	(*f).nama = "sate" // sama aja kaya yg bawah cuman ini versi ribetnya
	f.rate = 10
}

func MainUbahNamaFood() {
	food1 := Food{
		nama: "Gule",
		rate: 9,
	}

	ubahNamaFood(&food1)
	fmt.Println(food1)
}