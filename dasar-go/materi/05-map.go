package materi

import "fmt"

func UseMap() {
	// bikin map: [key]-value
	nilai := map[string]int{
		"Dimas": 100,
		"Azis":  98,
	}

	fmt.Println(nilai)
	fmt.Println(nilai["Dimas"])
	
	// nambah/update isi map
	nilai["Nabila"] = 99
	nilai["Dimas"] = 97
	fmt.Println(nilai)

	// cek apakah key ada?
	value, exists := nilai["Andi"]
	if !exists {
		fmt.Println("Andi tidak terdaftar")
	} else {
		fmt.Println(value)
	}

	// loop map
	for key, value := range nilai {
		fmt.Printf("%s: %d\n", key, value)
	}

}