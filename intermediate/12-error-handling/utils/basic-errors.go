package utils

import (
	"errors"
	"fmt"
)

func BasicErrorLearn() {

	// valPembagi, err := devide(10, 2)
	valPembagi, err := devide(10, 0)
	if err != nil {
		fmt.Println("Terjadi kesalah terdapat error", err)
		return
	}
	fmt.Printf("Devided test %v", valPembagi)

	errAge := validateAge(1000)

	if errAge != nil {
		fmt.Println(errAge)
	}
}

func devide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("tidak dapat membagi dengan 0")
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("umur tidak valid: %d (harus >= 0)", age)
	}
	if age >= 150 {
		return fmt.Errorf("umur tidak realistis: %d", age)
	}

	return nil
}
