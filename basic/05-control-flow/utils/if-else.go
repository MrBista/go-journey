package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallIfElseLearnSection() {
	helper.CalledFunction("CallIfElseLearnSection")
	isAdult := isAdult(1)

	if isAdult {
		fmt.Println("Hai aku sudah gede")
	}

	valGreat := getGreate(40)

	fmt.Println(valGreat)

	shortStatementIf()
}

func isAdult(age int) bool {
	return age >= 18
}

func getGreate(nilai int) string {

	if nilai > 97 {
		return "Greade A +"
	} else if nilai > 90 {
		return "Grade A"
	} else if nilai > 85 {
		return "Grade B"
	} else if nilai > 75 {
		return "Grade C"
	} else if nilai > 65 {
		return "Grade D"
	} else {
		return "Your Failed at exam"
	}
}

func shortStatementIf() {
	if age := 10; age > 5 {
		fmt.Println("Hai umur mu ", age)
	} else {
		fmt.Println("Umur ku berapa ya", age)
	}

}
