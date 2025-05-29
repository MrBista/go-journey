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


3. Tipe Data:

Numerik: int, int8/16/32/64, uint, float32/64
Boolean: true/false
String: teks dalam tanda petik




7. Tips dan Best Practices

Gunakan := untuk deklarasi variabel lokal karena lebih ringkas
Pilih tipe data yang tepat untuk menghemat memori:

Gunakan int untuk bilangan bulat umum
Gunakan float64 untuk bilangan desimal
Gunakan int8, int16 hanya jika benar-benar perlu menghemat memori


Nama variabel harus deskriptif: umurPengguna lebih baik dari u
Konstanta gunakan UPPER_CASE atau PascalCase
Selalu handle error saat konversi string ke number