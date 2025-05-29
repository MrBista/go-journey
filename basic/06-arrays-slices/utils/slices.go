package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallSliceLearnSection() {
	fmt.Printf("\n")
	helper.CalledFunction("CallSliceLearnSection")

	deklarasiSlice()
	makeSlice()
	slicingOperations()

	fmt.Println()
}

func deklarasiSlice() {
	var nilSlice []int

	fmt.Printf("Nil slice: %v, len: %d, cap: %d, nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)

	emptySlice := []int{}
	fmt.Printf("Empty slice: %v, len: %d, cap: %d, nil: %t\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)

	// slice literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Literal slice: %v, len: %d, cap: %d\n", numbers, len(numbers), cap(numbers))

	angka := []int{1, 2, 3, 4, 5}
	fmt.Printf("Valud angka %v \n", angka)

}

func makeSlice() {
	slice1 := make([]int, 5)

	fmt.Printf("make([]int, 5): %v, len: %d, cap: %d\n",
		slice1, len(slice1), cap(slice1))

	slice2 := make([]string, 3, 10)
	fmt.Printf("make([]string, 3, 10): %v, len: %d, cap: %d\n",
		slice2, len(slice2), cap(slice2))

	// 3. Zero capacity (error jika cap < len)
	// slice3 := make([]int, 5, 3) // ERROR: len > cap

	slice2[0] = "Hello"
	slice2[1] = "World"
	slice2[2] = "Go"

	fmt.Printf("Setelah diisi: %v\n", slice2)

	type Person struct {
		Name string
		Age  int
	}

	people := make([]Person, 2, 10)

	people[0] = Person{"Bisboy", 20}
	people[1] = Person{"Bisma", 10}
	// people[2] = Person{"Bang udah bang", 1}

	fmt.Printf("value dari people variabel %v", people)
}

func slicingOperations() {
	fmt.Println()
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Original %v\n", original)
}
