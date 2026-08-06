package materi

import (
	"fmt"
	"strings"
)

func UsePerulangan() {
	// 1. mirip loop biasa (C style)
	for i := 0; i < 7; i++ {
		fmt.Println(i)
	}

	// 2. gaya seperti while
	i := 0
	for i < 7 {
		fmt.Println(i)
		i++
	}

	// 3. infinite loop
	// for {
    // 	fmt.Println("Ini bakal loop terus")
	// }

	// 4. switch
	var hari string = "senin"
	// fmt.Print("masukkan hari: ")
	// fmt.Scan(&hari)
	hari = strings.ToLower(hari)

	switch hari {
	case "senin":
		fmt.Println("ini hari senin")
	case "selasa":
		fmt.Println("ini hari selasa")
	case "rabu":
		fmt.Println("ini hari rabu")
	default:
		fmt.Println(hari, "gadiajak wlee") 
	}

}