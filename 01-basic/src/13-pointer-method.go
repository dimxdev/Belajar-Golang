package src

import "fmt"

type Mahasiswa struct {
	Name string
	age  int
}

func (m *Mahasiswa) Brithday() {
	// (*m).age++
	m.age++ // sama kek diatas
}

func (m *Mahasiswa) Update(Name string, age int) {
	m.Name = Name
	m.age = age
}

func MainMahasiswa() {
	mahasiswa1 := Mahasiswa{
		Name: "Dimas",
		age:  20,
	}

	// (&mahasiswa1).Brithday()
	mahasiswa1.Brithday() // sama kek diatas jadi auto kalo method
	fmt.Println(mahasiswa1)
	
	mahasiswa1.Update("Budi", 54)
	fmt.Println(mahasiswa1)
}

