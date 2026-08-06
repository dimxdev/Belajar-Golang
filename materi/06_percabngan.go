package materi

import "fmt"

func UsePercabangan() {
	var umur int = 12
	// fmt.Print("masukkan umur: ")
	// fmt.Scan(&umur)

	
	if umur >= 17 {
		fmt.Println("udah dewasa")
	}
	
	if umur >= 17 {
		fmt.Println("udah dewasa")
	} else {
		fmt.Println("masi anak anak")
	}
	
	if umur >= 17 {
		fmt.Println("udah dewasa")
	} else if umur >= 10 {
		fmt.Println("lumayan lah")
	} else {
		fmt.Println("masi bocil")
	}

	// Deklarasi variabel langsung di If else
	if nilai := 99; nilai <= 80 {
		fmt.Println("busyukkk")
	} else {
		fmt.Println("bagus2")
	}

}