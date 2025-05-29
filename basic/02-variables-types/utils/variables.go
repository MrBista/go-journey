package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func ListOfVariabels() {
	// menggunakan var
	helper.CalledFunction("ListOfVariabels")

	calledWithVar()

	calledWithShortVariabel()

	// Menggunakan := (Short Variable Declaration)

}

func calledWithVar() {
	var nama string = "John"
	var umur int = 25
	var tinggi float64 = 175.5

	fmt.Println("Hai aku "+nama+" umur aku ", umur, " dan tinggi aku adalah ", tinggi)

	// Deklarasi tanpa nilai awal (akan mendapat zero value)
	var saldo int
	var status bool
	var pesan string

	fmt.Println("saldo ku ", saldo, " status saldo ku ", status, " pesan ", pesan)

	// deklarasi multiple variabel
	var (
		y int  = 10
		x int  = 1
		z bool = false
	)

	fmt.Println("x=", x, "y=", y, "z=", z)

	// Type inference (Go menentukan tipe otomatis)

	var harga = 50000
	var pi = 31.14
	var aktif = true

	fmt.Println("Harga=", harga, "pi=", pi, "aktif=", aktif)
}

func calledWithShortVariabel() {
	nama := "Bisboy"
	umur := 23
	tinggi := 170

	fmt.Println("Hai aku ", nama, " aku berumur ", umur, " tinggi aku adalah ", tinggi)
}
