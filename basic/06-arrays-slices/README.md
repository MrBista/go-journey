1. Apa itu Array?
Array adalah koleksi elemen dengan tipe data yang sama dan ukuran tetap. Ukuran array ditentukan saat kompilasi dan tidak bisa diubah.


Karakteristik Array:

- Ukuran tetap (fixed size)
- Semua elemen memiliki tipe data yang sama
- Indeks dimulai dari 0
- Disimpan secara berurutan di memory
- Passed by value (bukan reference)


2. Apa itu Slice?
Slice adalah reference ke bagian dari array yang mendasarinya. Slice memiliki ukuran yang dinamis dan lebih fleksibel daripada array.

Karakteristik Slice:

- Ukuran dinamis (dapat berubah)
- Reference ke underlying array
- Memiliki length dan capacity
- Passed by reference
- Nil slice adalah valid



Struktur Internal Slice:

Slice = {
    ptr:    pointer ke underlying array
    len:    panjang slice saat ini
    cap:    kapasitas maksimum dari posisi ptr
}

namaSlice := []tipeData{elemen1, elemen2, ...}

make([]tipeData, length, capacity)