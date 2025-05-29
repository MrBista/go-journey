package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallLoopsLearnSection() {
	helper.CalledFunction("CallLoopsLearnSection")

	describeForLoop()
}

func describeForLoop() {
	// for loop biasa
	for i := 0; i < 5; i++ {
		fmt.Println("Iterasi ke -", i)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println("Nilai ke ", i)
		i++
	}

	// infinite loop
	j := 0
	for {

		fmt.Println("Loop infinite, counter", j)
		j++
		if j >= 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			// skip kalau ganjil
			continue
		}

		fmt.Println("Semua valu genap sampai 10 iterasi", i)
	}

}
