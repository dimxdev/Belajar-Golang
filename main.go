package main

import (
	"Belajar-Go/materi"
	"fmt"
)

func main() {
	declare()
	fmt.Println()
	
	fmt.Println("=== Materi Variabel Tipedata ===")
	materi.VariabelTipedata()
	fmt.Println()
	
	fmt.Println("=== Materi Function ===")
	resultAddNumbers := materi.AddNumbers(15, 80)
	fmt.Println(resultAddNumbers)
	fmt.Println(materi.Pembagian(10, 7))
	fmt.Println()
	
	fmt.Println("=== Materi Struct ===")
	materi.UseStruct()
	fmt.Println()
	
	fmt.Println("=== Materi Slice ===")
	materi.UseSlice()
	fmt.Println()

	fmt.Println("=== Materi Map ===")
	materi.UseMap()
	fmt.Println()

	fmt.Println("=== Materi If Else ===")
	materi.UsePercabangan()
	fmt.Println()

	fmt.Println("=== Materi For ===")
	materi.UsePerulangan()
	fmt.Println()
}