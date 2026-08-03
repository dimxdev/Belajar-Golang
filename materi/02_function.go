package materi

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



