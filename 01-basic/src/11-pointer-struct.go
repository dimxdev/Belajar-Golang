package src

import "fmt"

type Kucing struct {
	Nama  string
	color string
	umur  int
}

func MainKucing() {
	kucing1 := Kucing{
		Nama:  "Alex",
		color: "Red",
		umur:  3,
	}
	// var pKucing1 *Kucing = &kucing1 
	pKucing1 := &kucing1 // sama aja kek yg diatas(versi singkat)
	pKucing1.Nama = "Boni" 

	fmt.Println(kucing1)
	fmt.Println(pKucing1)

	fmt.Println((*pKucing1).color) 
	fmt.Println(pKucing1.color) // sama kek yg diatas cuman lebih simple 
	fmt.Println() 
}


type Adress struct {
	City string
	Desa string
}

type Pekerja struct {
	Nama string
	Alamat Adress
}

func MainPekerja() {
	pekerja1 := Pekerja{
		Nama: "Budi",
		Alamat: Adress{
			City: "Jepara",
			Desa: "BandungHarjo",
		},
	}
	pPekerja1 := &pekerja1
	pPekerja1.Alamat.City = "Semarang"
	
	fmt.Println(pekerja1)
	fmt.Println(pPekerja1)
	fmt.Println(*pPekerja1)
	fmt.Println() 
}


type Anjing struct {
	Name string
	age *int
}

func MainAnjing() {
	age := 30 
	Anjing1 := Anjing{
		Name: "Lutsi",
		age: &age,
	}

	fmt.Println(Anjing1)
	fmt.Println(&Anjing1)
}