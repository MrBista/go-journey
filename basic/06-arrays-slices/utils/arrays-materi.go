package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallArrayLearnSection() {
	fmt.Printf("\n")
	helper.CalledFunction("CallArraySection")

	arrFuncDescribe()
	operasiArray()

	arrayMultiDimensional()
}

func arrFuncDescribe() {
	// array biasa
	var numbers [5]int
	fmt.Printf("Array kosong: %v\n\n", numbers)

	// deklarasi dengan nilai awal
	var fruits [3]string = [3]string{"apel", "jeruk", "mangga"}
	fmt.Printf("Array buah: %v\n", fruits)

	// deklari dengan ukuran otomatis
	colors := [...]string{"merah", "hijau", "biru", "kuning"}
	fmt.Printf("Array warna: %v (panjang: %d)\n", colors, len(colors))

	// deklarasi dengan index spesifik
	sparse := [5]int{1: 100, 2: 20}
	// yang ga diisi akan otomatis nol
	fmt.Printf("value spare adalah %v", sparse)
}

func operasiArray() {
	scores := [5]int{85, 100, 23, 1, 1}

	fmt.Printf("Nilai pertama: %d \n", scores[0])

	fmt.Printf("Nilai terakhir: %d \n", scores[len(scores)-1])

	scores[1] = 1
	fmt.Printf("Nilai scores setelah di ubah $%v", scores)

	fmt.Println("Iterasi dengan index: ")

	// for i := 0; i < len(scores); i++ {
	// 	fmt.Printf("Index %d: %d \n", i, scores[i])
	// }

	for i, val := range scores {
		fmt.Printf("score dengan index ke %v memiliki value %v\n", i, val)
	}

	for _, val := range scores {
		fmt.Printf("%d, ", val)
	}
}

func arrayMultiDimensional() {
	var matrix [3][4]int

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][2] = 5

	fmt.Printf("\nMatrix, %v\n", matrix)

	grid := [2][3]int{
		{1, 2, 3},
		{2, 1, 2},
	}

	fmt.Println("Grid:")

	for i, row := range grid {
		for j, value := range row {
			fmt.Printf("grid[%d][%d] = %d ", i, j, value)
		}

		fmt.Println()
	}

	// Array 3D
	cube := [2][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}
	fmt.Printf("Cube: %v\n", cube)
}
