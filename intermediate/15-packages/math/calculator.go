package math

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Tidak dapat membagi dengan 0")
	}

	return a / b, nil
}
