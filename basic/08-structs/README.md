# 📘 Catatan Belajar: Struct di Go

## 🧱 Apa Itu Struct?
- `struct` (singkatan dari *structure*) adalah tipe data komposit di Go.
- Digunakan untuk mengelompokkan beberapa field (properti) dengan tipe data berbeda menjadi satu kesatuan.
- Mirip dengan *class* di bahasa lain, tetapi tanpa pewarisan (*inheritance*).

## 🧠 Konsep Kunci
- **Field**: Anggota dari sebuah struct, bisa memiliki nama dan tipe data berbeda.
- **Instance/Object**: Ketika kamu membuat variabel dari struct, itu disebut instance.
- **Literal**: Cara langsung membuat struct beserta nilainya.

## 📌 Hal-Hal Penting Tentang Struct
- Field bisa diakses menggunakan dot (`.`).
- Struct bisa disisipkan ke struct lain (*embedded struct*) untuk komposisi.
- Bisa digunakan sebagai parameter atau return value pada fungsi.
- Bisa di-tag dengan metadata menggunakan *struct tags* (umumnya untuk JSON, DB, dll).
- Struct adalah *value type*, artinya ketika diberikan ke variabel/fungsi lain, akan dibuat salinan.

## 🛠️ Fungsi Pendukung Struct
- Bisa digunakan bersama dengan method untuk menambahkan *behavior*.
- Tidak ada *constructor* resmi, tetapi bisa menggunakan fungsi biasa sebagai pembuat instance.
- Dapat digunakan pointer untuk menghindari copy data saat mengoper ke fungsi.
- Bisa dibuat anonymous struct jika hanya dibutuhkan sekali.

## 🔁 Perilaku Khusus
- Jika ada field dengan huruf kapital (**exported**), bisa diakses dari luar package.
- Jika huruf kecil (**unexported**), hanya bisa diakses dari dalam package yang sama.

## ⚠️ Best Practice
- Gunakan struct untuk representasi entitas data (misal: `User`, `Product`, `Book`).
- Gunakan **pointer receiver** jika struct memiliki banyak field dan akan dimodifikasi.
- Pisahkan struct ke dalam folder seperti `dto`, `entity`, atau `model` sesuai arsitektur proyek.
- Gunakan naming yang jelas dan konsisten sesuai konvensi Go (camelCase atau PascalCase sesuai kebutuhan).

## 📚 Istilah-Istilah Penting
| Istilah            | Penjelasan                                                                 |
|--------------------|-----------------------------------------------------------------------------|
| **Field**          | Anggota struct, seperti atribut dalam OOP                                  |
| **Receiver**       | Parameter tambahan pada method untuk menunjukkan struct yang dimaksud      |
| **Method**         | Fungsi yang terasosiasi dengan struct                                       |
| **Embedded Struct**| Teknik menyisipkan satu struct ke dalam struct lain                        |
| **Struct Tag**     | Metadata tambahan dalam tanda backtick (contoh: `json:"name"`)             |
| **Value Type**     | Tipe data yang disalin saat diberikan ke fungsi atau variabel lain         |
| **Pointer Receiver**| Receiver yang menggunakan pointer agar method bisa mengubah struct asli  |

---

