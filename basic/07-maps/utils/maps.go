package utils

import (
	"fmt"
	"go-journey/basic/helper"
	"sort"
)

func CallMapLearnSection() {
	helper.CalledFunction("CallMapLearnSection")
	comperaseionWithOtherLanguage()

	fmt.Println()
	makeMap()
}

func comperaseionWithOtherLanguage() {
	// ===== BAGIAN 1: PERBANDINGAN DENGAN JAVA/JAVASCRIPT =====
	fmt.Println("1. PERBANDINGAN DENGAN JAVA/JAVASCRIPT")
	fmt.Println("Java HashMap<K,V>  | JavaScript Object/Map | Go map[K]V")
	fmt.Println("map.put(k, v)      | obj[k] = v / map.set() | map[k] = v")
	fmt.Println("map.get(k)         | obj[k] / map.get()     | v := map[k]")
	fmt.Println("map.containsKey(k) | k in obj / map.has()   | _, ok := map[k]")
	fmt.Println("map.remove(k)      | delete obj[k]          | delete(map, k)")
	fmt.Println()
}

func makeMap() {

	fmt.Println("1. CARA MEMBUAT MAP")
	// cara 1. mengunakan Map

	var siswa map[string]int
	siswa = make(map[string]int)
	fmt.Println("Map dengan zero value", siswa)

	// cara 2. Deklari langung dengan make

	nilai := make(map[string]int)
	fmt.Println("Map kosong langsung tanpa var", nilai)

	// cara 3. mengunakan Map Literal

	umur := map[string]int{
		"Ali":   20,
		"Baba":  10,
		"Citra": 15,
	}
	// people := map[string]interface{}{
	// 	"1": siswa,
	// }

	// fmt.Println("People isinya", people)

	fmt.Println("key value store umur ", umur)

	fmt.Println()

	fmt.Println("2. MENAMBAH DAN MENGUBAH DATA MAP")

	siswa["Olivia"] = 85
	siswa["Rodri"] = 80
	siswa["Golang"] = 100

	fmt.Println("Setelah menambah data siswa:", siswa)

	siswa["Rodri"] = 90
	fmt.Println("Setelah Nilai Rodri berubah", siswa)

	fmt.Println()

	fmt.Println("4. MENGAKSES DATA")
	nilaiOlivia := siswa["Olivia"]

	fmt.Println("Nilai Olivia", nilaiOlivia)

	nilaiBisboy, exist := siswa["Bisboy"]

	if exist {
		fmt.Println("Nilai Si Jenius Bisboy", nilaiBisboy)
	} else {
		fmt.Println("Bisboy tidak ditemukan dalam Map")
	}

	nilaiKosong := siswa["siswaKosong"]

	fmt.Println("Nilai siswaKosong (tidak ada)", nilaiKosong) // kalau ga ada dia nilainya adalah 0
	fmt.Println()

	fmt.Println("5. MENGHAPUS DATA")

	fmt.Println("Sebelum hapus Golang", siswa)
	delete(siswa, "Golang")

	fmt.Println("Setelah Hapus Golang", siswa)
	fmt.Println()

	fmt.Println("6. ITERASI MELALUI MAP")

	for key, value := range siswa {
		fmt.Printf("Siswa %v memiliki nilai %v\n", key, value)
	}

	// Hanya key saja
	fmt.Println("\nHanya nama-nama siswa:")
	for nama := range siswa {
		fmt.Printf("- %s\n", nama)
	}

	// Hanya value saja
	fmt.Println("\nHanya nilai-nilai:")
	for _, nilai := range siswa {
		fmt.Printf("- %d\n", nilai)
	}
	fmt.Println()

	fmt.Println("7. PANJANG MAP")
	fmt.Printf("Jumlah siswa ada %v \n", len(siswa))
	fmt.Println()

	fmt.Println("8. MAP DENGAN TIPE DATA BERBEDA")

	// map dengan value berupa slice
	hobi := map[string][]string{
		"Ali":  {"Membaca", "Menulis", "Tidur"},
		"Bobi": {"Olaharga", "Musik"},
		"Budi": {"Melukis", "nonton anime"},
	}

	fmt.Println("Map dengan slice sebagai value: ")
	for key, value := range hobi {
		fmt.Printf("%s: %v\n", key, value)
	}
	fmt.Println()

	// structcs -> mirip object kalau di js / java

	fmt.Println("9. MAP DENGAN STRUCT")

	type Mahasiswa struct {
		Nama string
		Umur int
		IPK  float64
	}

	mahasiswa := map[string]Mahasiswa{
		"230801": {
			Nama: "Gusti Bisman Taka",
			Umur: 23,
			IPK:  4.0,
		},
		"180103": {
			Nama: "Dyah Mekar Pangestu",
			Umur: 22,
			IPK:  4.0,
		},
	}

	fmt.Println("Nilai Map dengan struct/object", mahasiswa)

	for key, value := range mahasiswa {
		fmt.Printf("Mahasiswa dengan nim %s adalah %s dengan umur %v dan IPK %v\n", key, value.Nama, value.Umur, value.IPK)
	}
	fmt.Println()

	fmt.Println("10. NESTED MAP")

	sekolah := map[string]map[string]int{
		"Kelas 10A": {
			"Matematika": 85,
			"Fisika":     78,
			"Kimia":      82,
		},
		"Kelas 10B": {
			"Matematika": 90,
			"Fisika":     88,
			"Kimia":      85,
		},
	}

	fmt.Println("Nested map (map dalam map):")
	for kelas, mataPelajaran := range sekolah {
		fmt.Printf("%s:\n", kelas)
		for mapel, nilai := range mataPelajaran {
			fmt.Printf("  %s: %d\n", mapel, nilai)
		}
	}
	fmt.Println()

	// 11. TIPS DAN TRIK
	fmt.Println("11. TIPS DAN TRIK")

	// Mengecek apakah map kosong
	mapKosong := make(map[string]int)
	if len(mapKosong) == 0 {
		fmt.Println("Map kosong!")
	}

	fmt.Println("\nMengurutkan key map:")
	var keys []string
	for k := range siswa {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Siswa terurut berdasarkan nama:")
	for _, k := range keys {
		fmt.Printf("- %s: %d\n", k, siswa[k])
	}
	fmt.Println()

	fmt.Println("12. CONTOH PENGGUNAAN PRAKTIS")
	kalimat := "go adalah bahasa pemrograman go yang mudah dipelajari go"

	var kata []string

	perKata := ""

	// manual tanpa string.Split
	for _, value := range kalimat {

		if value != ' ' {
			perKata += string(value)
		} else {
			if perKata != "" {
				kata = append(kata, perKata)
				perKata = ""

			}
		}
	}

	frekuensiWordUsed := make(map[string]int)

	for _, val := range kata {
		frekuensiWordUsed[val]++
	}

	fmt.Println("Frekuensi kata digunakan", frekuensiWordUsed)

	// kalau mencari kata yang sering digunakan
	amountOfUsed := 0
	wordMostUsed := ""

	for key, value := range frekuensiWordUsed {
		if value > amountOfUsed {
			amountOfUsed = value
			wordMostUsed = key
		}
	}

	fmt.Printf("Word of most used is '%s' with used for '%v; times", wordMostUsed, amountOfUsed)

	fmt.Println()

	// 14. COPY MAP
	fmt.Println("14. COPY MAP")

	// Map adalah reference type, jadi hati-hati saat mengcopy
	mapAsli := map[string]int{"a": 1, "b": 2}
	mapRef := mapAsli // ini bukan copy, tapi reference yang sama!

	mapRef["c"] = 3
	fmt.Println("Map asli setelah mengubah mapRef:", mapAsli) // akan berubah juga!

	// Cara yang benar untuk copy map
	mapCopy := make(map[string]int)
	for k, v := range mapAsli {
		mapCopy[k] = v
	}

	mapCopy["d"] = 4
	fmt.Println("Map asli setelah mengubah mapCopy:", mapAsli) // tidak berubah
	fmt.Println("Map copy:", mapCopy)
	fmt.Println()

	fmt.Println("=== SELESAI ===")
}
