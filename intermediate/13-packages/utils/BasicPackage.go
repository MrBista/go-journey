package utils

import (
	"fmt"
	m "go-journey/intermediate/13-packages/math"
)

func CallPackageLearn() {
	valAdd := m.Add(2, 34)
	fmt.Printf("Hasil penjumlahan %v", valAdd)
}
