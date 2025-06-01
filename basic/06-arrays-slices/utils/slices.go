package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallSliceLearnSection() {
	fmt.Printf("\n")
	helper.CalledFunction("CallSliceLearnSection")

	// deklarasiSlice()
	// makeSlice()
	// slicingOperations()
	mainMateri()

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

func mainMateri() {
	fmt.Println("=== TUTORIAL SLICE GO ===\n")

	// 1. Membuat slice kosong
	fmt.Println("1. Slice Kosong:")
	var emptySlice []int
	fmt.Printf("Slice kosong: %v, Length: %d, Capacity: %d\n\n",
		emptySlice, len(emptySlice), cap(emptySlice))

	// 2. Membuat slice dengan make
	fmt.Println("2. Slice dengan make:")
	numbers := make([]int, 3, 5)
	fmt.Printf("make([]int, 3, 5): %v, Length: %d, Capacity: %d\n\n",
		numbers, len(numbers), cap(numbers))

	// 3. Slice literal
	fmt.Println("3. Slice Literal:")
	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Fruits: %v, Length: %d\n\n", fruits, len(fruits))

	// 4. Menambah elemen dengan append
	fmt.Println("4. Menambah Elemen:")
	fruits = append(fruits, "grape")
	fmt.Printf("Setelah append grape: %v\n", fruits)

	fruits = append(fruits, "mango", "pineapple")
	fmt.Printf("Setelah append mango & pineapple: %v\n\n", fruits)

	// 5. Mengakses elemen
	fmt.Println("5. Mengakses Elemen:")
	fmt.Printf("Elemen pertama: %s\n", fruits[0])
	fmt.Printf("Elemen kedua: %s\n", fruits[1])
	fmt.Printf("Elemen terakhir: %s\n\n", fruits[len(fruits)-1])

	// 6. Slicing
	fmt.Println("6. Slicing:")
	scores := []int{85, 92, 78, 96, 88, 74, 91}
	fmt.Printf("Semua scores: %v\n", scores)
	fmt.Printf("Index 1-3: %v\n", scores[1:4])
	fmt.Printf("3 pertama: %v\n", scores[:3])
	fmt.Printf("3 terakhir: %v\n", scores[len(scores)-3:])

	// 7. Mengubah elemen
	fmt.Println("\n7. Mengubah Elemen:")
	scores[0] = 90
	fmt.Printf("Setelah mengubah elemen pertama: %v\n", scores)

	// 8. Iterasi slice
	fmt.Println("\n8. Iterasi Slice:")
	fmt.Println("Dengan index:")
	for i, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", i, fruit)
	}

	fmt.Println("\nTanpa index:")
	for _, fruit := range fruits {
		fmt.Printf("- %s\n", fruit)
	}

	// 9. Copy slice
	fmt.Println("\n9. Copy Slice:")
	originalSlice := []int{1, 2, 3}
	copiedSlice := make([]int, len(originalSlice))
	copy(copiedSlice, originalSlice)

	fmt.Printf("Original: %v\n", originalSlice)
	fmt.Printf("Copied: %v\n", copiedSlice)

	// Mengubah original tidak mempengaruhi copy
	originalSlice[0] = 100
	fmt.Printf("Setelah mengubah original[0] = 100:\n")
	fmt.Printf("Original: %v\n", originalSlice)
	fmt.Printf("Copied: %v\n", copiedSlice)

	// 10. Slice multidimensi
	fmt.Println("\n10. Slice 2D (Matrix):")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matrix:")
	for i, row := range matrix {
		fmt.Printf("Baris %d: %v\n", i, row)
	}

	fmt.Printf("Elemen [1][2]: %d\n", matrix[1][2]) // 6
}
