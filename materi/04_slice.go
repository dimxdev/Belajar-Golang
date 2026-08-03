package materi

import "fmt"

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

	siswa = append(siswaB, "Nopal") //nimpa wkwk
	fmt.Println(siswa)

	campur := []any{"Budi", 123, true, 3.14}
	fmt.Println(campur) 
}