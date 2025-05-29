package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallRangeLearnSection() {
	helper.CalledFunction("CallRangeLearnSection")

	describeForLoopRange()
}

func describeForLoopRange() {

	// dengan Slice
	numbers := []int{10, 2, 3, 4, 1, 2, 100}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// hanya per item / object
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// Hanya menggunakan index
	// for index := range numbers {
	//     fmt.Println("Index:", index)
	// }

	text := "Hello"

	for index, char := range text {
		fmt.Printf("Index %d, Char: %c\n", index, char)
	}

	// dengan Map hash-map/ key value pair

	colors := map[string]string{
		"red":   "merah",
		"blue":  "biru",
		"green": "hijau",
	}

	// call key to get value its like hash table / hash map in java
	// fmt.Println("colors", colors["red"])

	for key, value := range colors {
		fmt.Println("Key", key, "memiliki value", value)
	}

}
