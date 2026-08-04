package materi

import "fmt"

type Barangs struct {
	Name string
	Stock int
	Harga int
}

func UseSlice() {
	siswa := []string{"Dimas", "Julian", "Rifqi"}
	fmt.Println(siswa)
	fmt.Println(siswa[0])
	fmt.Println(len(siswa))
	
	// operasi slice
	siswa = append(siswa, "Safril")
	fmt.Println(siswa)
	potongan := siswa[1:3] //ambil index 1 sampai sebelum 3
	fmt.Println(potongan)
	
	siswaB := []string{"siswo", "Faiz", "Rizqi"}
	fmt.Println(siswaB)
	gabungan := append(siswa, siswaB...)
	fmt.Println(gabungan)

	// loop slice
	fruits := []string{"nanas", "apel", "mangga", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("%d. %s\n", i, fruit)
	} 
	for _, fruit := range fruits {
		fmt.Println(fruit)
	} 
	for i := range fruits {
		fmt.Println(i)
	} 
	

	siswa = append(siswaB, "Nopal") //nimpa wkwk
	fmt.Println(siswa)

	campur := []any{"Budi", 123, true, 3.14}
	fmt.Println(campur)
	
	// latihan struct + slice
	 barangs := []Barangs{
		{Name: "Keyboard", Stock: 5, Harga: 100000},
		{Name: "Laptop", Stock: 10, Harga: 1000000},
	 }

	 for _, barang := range barangs {
		fmt.Println(barang.Name, "-", barang.Stock, "-", barang.Harga)
	 }
}