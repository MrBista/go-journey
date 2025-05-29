package utils

import "fmt"

func ZeroValues() {
	fmt.Println("=========== List of Zero Values called ===========")
	var angka int       // 0
	var desimal float64 // 0.0
	var teks string     // "" (string kosong)
	var status bool     // false
	// var pointer *int    // nil

	fmt.Printf("int: %d\n", angka)       // 0
	fmt.Printf("float64: %f\n", desimal) // 0.000000
	fmt.Printf("string: '%s'\n", teks)   // ''
	fmt.Printf("bool: %t\n", status)     // false
}
