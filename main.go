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
	name, age, address, err := materi.GetUserInfo()
	fmt.Printf("name: %s\nage: %d\naddress: %s\nerror:%v", name, age, address, err)
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
	
	fmt.Println("=== Materi Error Handling ===")
	materi.ErrorHandling1()
	materi.ErrorHandling2()
	materi.ErrorHandling3()
	fmt.Println()

	fmt.Println("=== Materi Method ===")
	materi.PrintSapa()
	materi.PrintUbahNama()
	fmt.Println()

	fmt.Println("=== Materi Pointer ===")
	materi.MainPointer()
	fmt.Println()
}