package utils

import (
	"fmt"
	"go-journey/basic/helper"
)

func CallConstantLearnSection() {
	helper.CalledFunction("CallConstantLearnSection")

	// constanta sederhana
	const pi = 3.14159
	const perusahaan = "PT. Maju Jaya"
	const maxUser = 1000

	// multiple constanta
	const (
		StatusAktif   = 1
		StatusPending = 2
		StatusBatal   = 3
	)

	// konstanta dengan tipe eksplisit
	const (
		bunga float64 = 0.05
		pajak int     = 10
	)

	// konstanta dengan iota (auto-increment)

	const (
		Senin  = iota // 0
		Selasa        // 1
		Rabu          // 2
		Kamis         // 3
		Jumat         // 4
		Sabtu         // 5
		Minggu        // 6

	)

	fmt.Println("Perusahaan =", perusahaan, "Nilai Pi =", pi, "maximalUser =", maxUser, "Status Aktif=", StatusAktif, "pajak=", pajak, "ValMinggu=", Minggu)
}
