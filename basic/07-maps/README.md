RINGKASAN PENTING TENTANG MAP DI GO:

1. DEFINISI:
   - Map adalah struktur data key-value
   - Mirip dengan HashMap, yang pada pada dasarnya memang hasmap yg menyimpan key value store

2. KARAKTERISTIK:
   - Map adalah reference type (seperti slice dan pointer) / passed by reference
   - Map tidak berurutan (unordered)
   - Key harus bisa dibandingkan (comparable)
   - Zero value dari map adalah nil

3. OPERASI DASAR:
   - Buat: make(map[KeyType]ValueType) atau map[KeyType]ValueType{}
   - Tambah/Ubah: map[key] = value
   - Akses: value := map[key] atau value, ok := map[key]
   - Hapus: delete(map, key)
   - Panjang: len(map)

4. BEST PRACTICES:
   - Selalu cek apakah key ada sebelum menggunakan nilai
   - Gunakan make() untuk membuat map kosong
   - Hati-hati dengan reference saat passing ke function
   - Map tidak thread-safe, gunakan mutex jika perlu

5. KEGUNAAN UMUM:
   - Menyimpan data dengan key unik
   - Menghitung frekuensi
   - Caching/memoization
   - Lookup table
   - Grouping data