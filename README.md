Ekspor Fungsi
Dalam Go, fungsi yang dimulai dengan huruf kapital (seperti GetGreeting, Sum) dapat diakses dari package lain (exported). Fungsi dengan huruf kecil hanya dapat diakses dalam package yang sama.



Import Path
"myproject/utils" mengacu pada module name (myproject) diikuti dengan path ke package (utils). Ini didefinisikan dalam file go.mod.
Package vs Module

Package: koleksi file Go dalam direktori yang sama
Module: koleksi packages dengan versioning (didefinisikan dalam go.mod)



Tips Tambahan

Naming Convention: Gunakan camelCase untuk nama fungsi dan variabel
Package Name: Sebaiknya menggunakan nama pendek dan deskriptif
Import Grouping: Kelompokkan standard library, third-party, dan local imports
Documentation: Tambahkan komentar sebelum fungsi yang diekspor