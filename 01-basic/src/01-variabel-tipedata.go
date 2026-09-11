package src

import "fmt"

func VariabelTipedata() {
	// Penulisan Variabel
		// 1. var + nama + tipe data
		var name string = "Dimas Brotowali Hidayatullah"
		var age int = 20
		fmt.Println("nama:", name)
		fmt.Println("umur:", age)

		// 2. short declaration (nama + :=)
		city := "Jepara"
		ipk := 3.45
		fmt.Println("asal kota:", city)
		fmt.Println("ipk:", ipk)

		// 3. nulis 2 variabel sekaligus
		a, b := 1, 3
		fmt.Println(a)
		fmt.Println(b)
		
	// Tipe Data
	var isActive bool = true
	var score int = 98
	var price float64 = 1500.678
	var nickName string = "dimxdev"
	fmt.Println(isActive, score, price, nickName)
}