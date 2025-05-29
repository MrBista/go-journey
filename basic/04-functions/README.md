💡 Tips Penting

Error Handling: Gunakan multiple return values untuk error handling
Resource Management: Selalu gunakan defer untuk cleanup
Closures: Hati-hati dengan variable capture dalam loop
Performance: Variadic functions membuat slice baru, pertimbangkan performance untuk data besar
Readability: Named return values bagus untuk dokumentasi, tapi jangan overuse

🚨 Hal yang Perlu Diperhatikan

Defer dengan loops: Jika menggunakan defer dalam loop, semua defer akan menumpuk
Closure dalam loop: Variable loop akan di-capture by reference, bukan by value
Panic recovery: Gunakan defer dengan recover() untuk menangani panic
Memory leaks: Closure bisa menyebabkan memory leak jika menangkap variabel besar







Closures
Closure adalah fungsi yang "menangkap" variabel dari scope yang melingkupinya. Variabel tersebut tetap hidup selama closure masih digunakan.Artinya, fungsi bisa "mengingat" variabel luar meskipun eksekusi fungsi asalnya sudah selesai.


Currying
Currying adalah proses mengubah sebuah fungsi yang menerima banyak argumen menjadi rangkaian fungsi yang masing-masing hanya menerima satu argumen.

Perbedaan Konseptual
Closure: Fokus pada menyimpan konteks dari variabel luar.

Currying: Fokus pada transformasi fungsi dari multi-argumen jadi nested single-argumen.

Kalau closure itu kayak:

"Saya simpan nilai x dan akan ingat selamanya."

Sedangkan currying itu kayak:

"Saya belum punya semua argumen, kasih saya satu per satu, nanti saya kerjakan kalau lengkap."


Defer
defer menunda eksekusi statement hingga fungsi berakhir
Dieksekusi dalam urutan LIFO (Last In, First Out) / stack
Sangat berguna untuk cleanup resources (close files, unlock mutex, etc.)
Parameter untuk defer dievaluasi saat defer dipanggil, bukan saat dieksekusi
hindari call deffer function di for loop karena bisa menyebabkan memory leek


Panic
panic adalah mekanisme di GoLang yang digunakan untuk menangani kesalahan yang sifatnya tidak dapat dipulihkan (unrecoverable errors). Ketika sebuah panic terjadi, eksekusi program normal akan berhenti, dan fungsi-fungsi yang ditunda (defer) akan dieksekusi. Setelah semua defer dieksekusi, program akan berhenti (crash) dan mencetak stack trace (urutan panggilan fungsi yang mengarah ke panic).


Recover
Apa itu Recover?

recover adalah fungsi built-in di GoLang yang digunakan untuk menangkap atau mengatasi panic yang sedang berlangsung. recover hanya efektif jika dipanggil di dalam sebuah fungsi defer. Jika recover dipanggil di luar defer, ia tidak akan melakukan apa-apa dan mengembalikan nil.

Mengapa Menggunakan Recover?

Anda mungkin ingin menggunakan recover untuk:

1. gracefully shut down (menutup secara anggun):** Membersihkan sumber daya atau mencatat kesalahan sebelum program berhenti.

2. Melanjutkan eksekusi (jarang dan hati-hati): Dalam kasus-kasus tertentu, Anda mungkin ingin mencoba melanjutkan program setelah panic, tetapi ini sangat jarang dan harus dilakukan dengan sangat hati-hati karena bisa menyembunyikan bug yang lebih besar.