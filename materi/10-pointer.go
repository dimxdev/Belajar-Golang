package materi

import "fmt"

func MainPointer() {
	nama := "Dimas"
	pNama := &nama // & untuk simpen alamat di pointer
	var umur int = 20
	var pUmur *int = &umur 

	fmt.Println(nama)
	fmt.Println(pNama)
	fmt.Println(&nama)
	fmt.Println(*pNama) // dapetin value alamat
	fmt.Println()
	fmt.Println(umur)
	fmt.Println(pUmur)
	fmt.Println(&umur)
	fmt.Println(*pUmur)
	fmt.Println()
	
	
	alamat := "Jepara"
	alamatReal := &alamat
	*alamatReal = "Bandung" // ubah value alamat lewat pointer
	
	fmt.Println(alamat)
	fmt.Println(*alamatReal)
	fmt.Println()
	
	
	var role *int
	fmt.Println(role)
}