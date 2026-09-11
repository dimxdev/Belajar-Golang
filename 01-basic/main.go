package main

import (
	"Belajar-Go/src"
	"fmt"
)

func main() {
	declare()
	fmt.Println()
	
	fmt.Println("=== Materi Variabel Tipedata ===")
	src.VariabelTipedata()
	fmt.Println()
	
	fmt.Println("=== Materi Function ===")
	resultAddNumbers := src.AddNumbers(15, 80)
	fmt.Println(resultAddNumbers)
	fmt.Println(src.Pembagian(10, 7))
	name, age, address, err := src.GetUserInfo()
	fmt.Printf("name: %s\nage: %d\naddress: %s\nerror:%v", name, age, address, err)
	fmt.Println() 
	
	fmt.Println("=== Materi Struct ===")
	src.UseStruct()
	fmt.Println()
	
	fmt.Println("=== Materi Slice ===")
	src.UseSlice()
	fmt.Println()

	fmt.Println("=== Materi Map ===")
	src.UseMap()
	fmt.Println()

	fmt.Println("=== Materi If Else ===")
	src.UsePercabangan()
	fmt.Println()

	fmt.Println("=== Materi For ===")
	src.UsePerulangan()
	fmt.Println()
	
	fmt.Println("=== Materi Error Handling ===")
	src.ErrorHandling1()
	src.ErrorHandling2()
	src.ErrorHandling3()
	fmt.Println()

	fmt.Println("=== Materi Method ===")
	src.PrintSapa()
	src.PrintUbahNama()
	src.MainAge()
	fmt.Println()

	fmt.Println("=== Materi Pointer ===")
	src.MainPointer()
	fmt.Println()
	
	fmt.Println("=== Materi Pointer Di Function ===")
	src.MainUbahNilai()
	src.MainUbahNamaFood()
	fmt.Println()

	fmt.Println("=== Materi Pointer Di Struct ===")
	src.MainKucing()
	src.MainPekerja()
	src.MainAnjing()
	fmt.Println()

	fmt.Println("=== Materi Pointer Di Method ===")
	src.MainMahasiswa()
	fmt.Println()

	fmt.Println("=== Materi Goroutine Di Method ===")
	src.MainSapa()
	src.MainMasak()
	src.MainHitungKuadrat()
	src.MainKirimPesan()
	fmt.Println()
}