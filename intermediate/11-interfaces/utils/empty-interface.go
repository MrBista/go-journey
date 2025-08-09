package utils

import "fmt"

func printAnything(value interface{}) {
	// interface yang kosong bisa menerima value apapun
	/*
		- Bisa menampung type apapun
		- Mirip dengan Object di Java atau any di TypeScript
		- Sering digunakan untuk generic programming sebelum Go 1.18
	*/
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func EmptyInterfaceLearn() {
	printAnything(2)
	printAnything("BabaYaga")
	printAnything([]int{1, 3, 45, 2})
	printAnything(true)
}
