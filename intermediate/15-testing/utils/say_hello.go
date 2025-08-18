package utils

import "fmt"

func SayHello(name string) string {
	return "Hello " + name
}

func DividedFunc(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("Tidak bisa membagi dengan 0")
	}

	valResult := num1 / num2

	return valResult, nil

}
