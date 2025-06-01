package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

type Person struct {
	Name   string
	Age    int
	Height float64
}

func CallStructsLearnSection() {
	helper.CalledFunction("CallStructsLearnSection")
	deklarationStruct()
	modifikasiAndAksesStruct()

}

func deklarationStruct() {
	fmt.Println("1. CARA MEMBUAT STRUCT")
	// 1. literal
	person1 := Person{
		Name:   "Bisma",
		Age:    20,
		Height: 170.0,
	}

	fmt.Println("Struct literal person", person1)

	// 2. menggunakan var, baik kalau kita belum pasti mengisi setiap propertynya
	var person2 Person

	person2.Age = 21
	person2.Name = "Gusti"
	person2.Height = 180

	fmt.Println("Struct dengan var", person2)

	// 3. mengunakan urutan posisi, menurutku sangat tidak rekomend
	person3 := Person{"Bratha", 20, 170.0}

	fmt.Println("Struct dengan urutan ", person3)

	fmt.Println()

}

func modifikasiAndAksesStruct() {
	// mengakses setiap field atau attribute sama dengan javascipt dengan '.'
	fmt.Println("2. AKSES DAN MODIFIKASI STRUCT")
	// Akses struct
	person1 := Person{
		Name:   "Bisma",
		Age:    23,
		Height: 170,
	}

	fmt.Printf("struct pertama memiliki nama %v dengan umur %v dan tinggi %v\n", person1.Name, person1.Age, person1.Height)

}
