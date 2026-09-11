package src

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

type Product struct {
	Name string
	stock int
	Price float32
}

func UseStruct() {
	// Cara 1: isi berurutan
	user1 := User {
		"Dimas",
		"brotowalidimas@gmail.com",
		20,
	}

	// Cara 2: isi pakai nama field(lebih jelas)
	user2 := User {
		Name:  "Acep",
		Email: "acep@gmail.com",
		Age:   25,
	}

	var product1 Product = Product{"laptop", 15, 15000.000,}
	product2 := Product{
		Name: "handphone",
		stock: 100,
		Price: 17000.000,

	}

	fmt.Println(user1.Email)
	fmt.Println(user2)
	fmt.Println(product1.Name)
	fmt.Println(product2.Price)
}