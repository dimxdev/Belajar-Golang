package materi

import (
	"errors"
	"fmt"
)

func Bagi(a int, b int) (int, error) {
	if b == 0 {
		// return 0, errors.New("tidak bisa dibagi dengan 0")
		return 0, fmt.Errorf("%d tidak bisa dibagi dengan 0", a)
	}
	return a / b, nil
}

func ErrorHandling1() {
	hasil, err := Bagi(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Hasil:", hasil)
}

func CariUser(id int) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("id yang dimasukkan tidak valid (%d)", id)
	}
	if id != 1 {
		return  "", fmt.Errorf("user dengan id %d tidak ditemukan", id)
	}
	return "Dimas", nil
}

func ErrorHandling2() {
	hasil, err := CariUser(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hasil)
}

func cekStok(stok int) error {
	if stok <= 0 { 
		return errors.New("Stok Habis")
	}
	if stok <= 5 {
		return  errors.New("Stok menipis")
	}
	return nil
}

func ErrorHandling3() {
	stoks := []int{0, 2, 10, -3}

	for _, stok := range stoks {
		err := cekStok(stok)
		if err != nil {
			fmt.Println("Stok", stok, "->", err)
			continue
		}
		fmt.Println("Stok", stok, "-> Aman")
	}
}