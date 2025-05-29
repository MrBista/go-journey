Perbedaan var vs :=:

var bisa digunakan di level package dan function
:= hanya bisa digunakan di dalam function
:= lebih singkat dan umum digunakan

1. Deklarasi Variabel:

var nama string = "value" - deklarasi eksplisit
nama := "value" - short declaration (hanya di dalam fungsi)
Go menggunakan type inference untuk menentukan tipe otomatis


2. Zero Values:

Setiap tipe punya nilai default: 0 untuk angka, false untuk bool, "" untuk string