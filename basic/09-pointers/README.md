**Apa Itu Pointer?**
Pointer adalah variabel yang menyimpan alamat memori dari variabel lain. Ibaratnya, pointer menyimpan "lokasi" atau "alamat rumah" dari sebuah data, bukan datanya itu sendiri.

**Kenapa Pakai Pointer?**

1. Mengubah Variabel Asli: Go meneruskan salinan variabel ke fungsi. Pointer memungkinkan kamu mengubah nilai variabel asli di luar lingkup fungsi yang memanggilnya.
2. Efisiensi Memori: Untuk tipe data besar (seperti struct yang kompleks), meneruskan pointer jauh lebih efisien daripada menyalin seluruh data.
3. Membangun Struktur Data: Penting untuk membuat struktur data dinamis seperti linked list dan tree.
4. Menunjukkan Ketiadaan Nilai: Pointer bisa bernilai nil, artinya ia tidak menunjuk ke alamat memori manapun. Berguna untuk mengindikasikan bahwa suatu variabel belum diinisialisasi atau tidak memiliki nilai.




**Dasar Penggunaan Pointer**
1. Deklarasi Pointer: Tambahkan * di depan tipe data. Contoh: var namaPointer *tipeData.
2. Mengambil Alamat Variabel: Gunakan operator & (address-of). Contoh: ptrAngka = &angka. Ini membuat ptrAngka menyimpan alamat memori dari angka.
3. Mengakses Nilai (Dereferencing): Gunakan operator * di depan pointer. Contoh: fmt.Println(*ptrAngka). Ini akan menampilkan nilai yang ditunjuk oleh ptrAngka. Untuk mengubah nilai, gunakan *ptrAngka = nilaiBaru.



**Hal Penting yang Perlu Diingat**
1. Nilai nil: Pointer yang belum diinisialisasi memiliki nilai nil. Mencoba mengakses nilai dari pointer nil akan menyebabkan panic (kesalahan program).
2. Tidak Ada Aritmatika Pointer: Berbeda dengan C/C++, Go tidak mengizinkan penambahan atau pengurangan pada pointer. Ini adalah fitur keamanan Go untuk mencegah akses memori yang tidak valid.
3. Pointer ke Struct: Saat menggunakan pointer ke struct, kamu bisa langsung mengakses field menggunakan operator . (titik). Go akan secara otomatis melakukan dereference untukmu. Contoh: ptrP.Name (otomatis jadi (*ptrP).Name).
4. Fungsi new(): Mengalokasikan memori untuk tipe data dan mengembalikan pointer ke nilai zero value (nilai default) dari tipe data tersebut. Jarang digunakan untuk variabel yang sudah ada, lebih sering untuk alokasi baru.
5. make() vs new(): new() mengembalikan pointer ke zero value dari tipe data apa pun. make() hanya untuk slice, map, dan channel, menginisialisasi struktur data internalnya, dan mengembalikan tipe datanya sendiri (bukan pointer).



**Kapan Menggunakan Pointer?**
1. Saat kamu harus mengubah nilai variabel asli di dalam sebuah fungsi.
2. Untuk efisiensi memori saat menangani tipe data atau struct yang sangat besar.
3. Untuk membangun struktur data dinamis.
4. Untuk mengindikasikan bahwa suatu nilai opsional atau belum ada (nil).


**Kapan Tidak Perlu Menggunakan Pointer?**
1. Untuk tipe data dasar yang kecil (misalnya int, string, bool) jika kamu tidak perlu mengubah nilai aslinya di dalam fungsi.
2. Jika kamu hanya ingin membaca nilai variabel.