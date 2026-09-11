package src

import "fmt"

func AddNumbers(a int, b int) int {
	return a + b
}

func Pembagian(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa dibagi nol")
	}
	return a / b, nil 
}

func GetUserInfo() (string, int, string, error) {
	return "Dimas", 20, "Jepara", nil
}

