package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallDeferLearnSection() {
	helper.CalledFunction("CallDeferLearnSection")

	// oneline
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered", r)
	// 	}
	// }()

	// callmethod

	defer CatchHandler()
	callDefferFunc()

	panic("error unexpected ini")

}

func callDefferFunc() {
	// walaupun ini di call di awal tetap saja ini akan di eksekusi di akhir karena LIFO

	// ini terakhir
	defer fmt.Println("deffer called")
	// ini dulu karena dia last in
	defer fmt.Println("Coba call deffer")

	fmt.Println("Hallo dunia")
}

func CatchHandler() {
	if r := recover(); r != nil {
		fmt.Println("catch recover from diver", r)
	}
}
