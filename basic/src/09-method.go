package src

import "fmt"

type Employee struct {
	Name string
	age  int
	role string
}

func (e Employee) Sapa() string {
	return "Halo, saya " + e.Name
}

func PrintSapa() {
	employe1 := Employee{Name: "Dimas", age: 20, role: "BE"}
	fmt.Println(employe1.Sapa())
}

func (e *Employee) UbahNama(namaBaru string) {
	e.Name = namaBaru
}

func PrintUbahNama() {
	employe1 := Employee{Name: "Dimas", age: 20, role: "BE"}

	employe1.UbahNama("Dimas Brotowali H")
	fmt.Println(employe1)
}


type Age int 

func (a Age) isAdult() bool {
	return a >= 20
}

func MainAge() {
	age := Age(17)
	ageIsAdult := age.isAdult()

	fmt.Println(ageIsAdult)
}
