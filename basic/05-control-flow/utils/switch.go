package utils

import "fmt"

func CallSwitchLearnSection() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error apa ini gengs", r)
		}
	}()

	describeJamKerja := getStatusJamKerja("senin")

	fmt.Println(describeJamKerja)
}

func getStatusJamKerja(hari string) string {
	switch hari {
	case "senin", "selasa", "rabu", "kamis", "jumat":
		return "Hari Kerja"
	case "sabtu", "minggu":
		return "Weekend gengs"
	default:
		panic("Ga ada hari ini geng, masukin semua huruf kecil dan pastikan sesuai")
	}
}
