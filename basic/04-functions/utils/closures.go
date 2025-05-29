package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallClosureLearnSection() {
	helper.CalledFunction("CallClosureLearnSection")

	// valcount akan sebagai closure dan sebuah function yg perlu di call lagi
	valCount := counterClosure(2)
	valCount()
	fmt.Println("Val counterClosure", valCount())

	// currying perlu di call dan kirim param lagi untuk funcnya jalan

	adds := addCurying(2)

	fmt.Println("currying expale", adds(1))
}

// mengembalikan function untuk closure
func counterClosure(val int) func() int {
	startVal := val

	return func() int {
		startVal += 1
		return startVal
	}
}

func addCurying(num int) func(int) int {

	return func(num2 int) int {

		return num + num2
	}
}
