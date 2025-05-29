1. Main Function (Entry Point)
Setiap program Go yang dapat dijalankan harus memiliki fungsi main() dalam package main. Ini adalah titik masuk program Anda.


2. Packages
Package adalah cara Go untuk mengorganisir dan mengelompokkan kode. Setiap file Go harus dimulai dengan deklarasi package.

package main - untuk program yang dapat dijalankan
package namapackage - untuk library/package lain


3. Import Statement
Import digunakan untuk menggunakan package lain dalam kode Anda.


Menjalankan Kode Go:

1. go run
Kompilasi dan jalankan kode secara langsung:

```
go run main.go
```

2. go build
Kompilasi kode menjadi executable:

```
go build main.go      # menghasilkan executable
./main               # menjalankan executable (Linux/Mac)
main.exe    
```

3. go mod init
Membuat Go module (untuk dependency management):

```
go mod init nama-project
```


