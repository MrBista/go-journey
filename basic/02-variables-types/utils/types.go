package utils

import "fmt"

// Signed integers
var a int8 = 127    // -128 to 127
var b int16 = 32767 // -32,768 to 32,767
var c int32 = 2147483647
var d int64 = 9223372036854775807

// Unsigned integers
var e uint8 = 255    // 0 to 255
var f uint16 = 65535 // 0 to 65,535
var g uint32 = 4294967295
var h uint64 = 18446744073709551615

// Platform-dependent
var i int = 42   // 32 atau 64 bit tergantung sistem
var j uint = 100 // 32 atau 64 bit tergantung sistem

// Floating Point
var harga float32 = 19.99      // Single precision
var saldo float64 = 1000000.75 // Double precision (lebih umum)

// Boolean
var aktif bool = true
var selesai bool = false

// String
var nama string = "Budi Santoso"
var alamat string = `Jl. Merdeka No. 123
Jakarta Pusat` // Raw string literal (multiline)

func ListOfTypeVariabel() {
	fmt.Println("=========== List of Variabel Called ===========")
}
