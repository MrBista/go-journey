# Dasar-Dasar Goroutines Go untuk Pemula

#### Gorutines materi: https://docs.google.com/presentation/d/1A78dn_g6HfxfRor9XBUAGPQM9vT6_SnrQGrQ2z0myOo/edit?slide=id.p#slide=id.p


## 1. Apa itu Goroutines?

Bayangkan Anda sedang memasak di dapur:
- **Tanpa goroutines**: Anda masak nasi dulu sampai selesai, baru masak sayur, baru masak lauk
- **Dengan goroutines**: Anda masak nasi, sambil menunggu nasi matang, Anda masak sayur dan lauk secara bersamaan

Goroutines memungkinkan program Go menjalankan beberapa tugas secara **bersamaan** (concurrent).

### Karakteristik Goroutines:
- Sangat ringan (hanya butuh ~2KB memory)
- Bisa membuat ribuan goroutines tanpa masalah
- Dikelola otomatis oleh Go runtime

## 2. Syntax Dasar

```go
// Cara biasa menjalankan function
functionName()

// Cara menjalankan function sebagai goroutine
go functionName()
```

Cukup tambahkan kata `go` di depan pemanggilan function!

## 3. Contoh Pertama - Tanpa Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func cetakNomor(nama string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", nama, i)
        time.Sleep(500 * time.Millisecond) // Tidur 0.5 detik
    }
}

func main() {
    fmt.Println("Program dimulai")
    
    // Tanpa goroutine - berjalan berurutan
    cetakNomor("Pertama")
    cetakNomor("Kedua")
    
    fmt.Println("Program selesai")
}
```

**Output:**
```
Program dimulai
Pertama: 1
Pertama: 2
Pertama: 3
Pertama: 4
Pertama: 5
Kedua: 1
Kedua: 2
Kedua: 3
Kedua: 4
Kedua: 5
Program selesai
```

Program ini butuh waktu sekitar 5 detik karena berjalan **berurutan**.

## 4. Contoh Kedua - Dengan Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func cetakNomor(nama string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", nama, i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    fmt.Println("Program dimulai")
    
    // Dengan goroutine - berjalan bersamaan
    go cetakNomor("Goroutine-1")
    go cetakNomor("Goroutine-2")
    
    // MASALAH: Program langsung selesai!
    fmt.Println("Program selesai")
}
```

**Output (kemungkinan):**
```
Program dimulai
Program selesai
```

**Masalah**: Program utama (main function) selesai lebih dulu sebelum goroutines sempat jalan!

## 5. Solusi: Menunggu Goroutines

### Solusi 1: time.Sleep (Cara sederhana tapi tidak ideal)

```go
package main

import (
    "fmt"
    "time"
)

func cetakNomor(nama string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", nama, i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    fmt.Println("Program dimulai")
    
    go cetakNomor("Goroutine-1")
    go cetakNomor("Goroutine-2")
    
    // Tunggu 3 detik agar goroutines sempat selesai
    time.Sleep(3 * time.Second)
    
    fmt.Println("Program selesai")
}
```

**Output:**
```
Program dimulai
Goroutine-1: 1
Goroutine-2: 1
Goroutine-1: 2
Goroutine-2: 2
Goroutine-1: 3
Goroutine-2: 3
Goroutine-1: 4
Goroutine-2: 4
Goroutine-1: 5
Goroutine-2: 5
Program selesai
```

Sekarang program butuh waktu sekitar 3 detik karena berjalan **bersamaan**!

## 6. Solusi yang Lebih Baik: sync.WaitGroup

`time.Sleep` tidak ideal karena:
- Kita harus menebak berapa lama goroutines butuh waktu
- Kalau terlalu pendek, goroutines belum selesai
- Kalau terlalu panjang, program menunggu tidak perlu

Solusi yang lebih baik adalah `sync.WaitGroup`:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func cetakNomor(nama string, wg *sync.WaitGroup) {
    // Beritahu WaitGroup bahwa goroutine ini sudah selesai
    defer wg.Done()
    
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s: %d\n", nama, i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    fmt.Println("Program dimulai")
    
    var wg sync.WaitGroup
    
    // Beritahu WaitGroup bahwa akan ada 2 goroutines
    wg.Add(2)
    
    go cetakNomor("Goroutine-1", &wg)
    go cetakNomor("Goroutine-2", &wg)
    
    // Tunggu sampai semua goroutines selesai
    wg.Wait()
    
    fmt.Println("Program selesai")
}
```

### Penjelasan WaitGroup:
1. `wg.Add(2)` - Beritahu bahwa akan ada 2 goroutines
2. `defer wg.Done()` - Ketika function selesai, beritahu WaitGroup
3. `wg.Wait()` - Tunggu sampai semua goroutines memanggil `Done()`

## 7. Contoh dengan Anonymous Function

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    
    // Goroutine dengan anonymous function
    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 1; i <= 3; i++ {
            fmt.Printf("Anonymous goroutine: %d\n", i)
            time.Sleep(300 * time.Millisecond)
        }
    }()
    
    // Goroutine dengan anonymous function dan parameter
    wg.Add(1)
    go func(nama string) {
        defer wg.Done()
        for i := 1; i <= 3; i++ {
            fmt.Printf("%s: %d\n", nama, i)
            time.Sleep(300 * time.Millisecond)
        }
    }("Parameter goroutine")
    
    wg.Wait()
    fmt.Println("Selesai!")
}
```

## 8. Contoh Praktis: Download Simulator

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func downloadFile(namaFile string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("Mulai download %s...\n", namaFile)
    
    // Simulasi waktu download (1-3 detik)
    waktuDownload := time.Duration(rand.Intn(3)+1) * time.Second
    time.Sleep(waktuDownload)
    
    fmt.Printf("✅ %s selesai didownload (%v)\n", namaFile, waktuDownload)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    
    files := []string{
        "dokumen.pdf",
        "gambar.jpg", 
        "video.mp4",
        "musik.mp3",
    }
    
    fmt.Println("🚀 Mulai download semua file...")
    startTime := time.Now()
    
    var wg sync.WaitGroup
    
    // Download semua file secara bersamaan
    for _, file := range files {
        wg.Add(1)
        go downloadFile(file, &wg)
    }
    
    wg.Wait()
    
    totalTime := time.Since(startTime)
    fmt.Printf("🎉 Semua file selesai didownload dalam %v\n", totalTime)
}
```

## 9. Loop dengan Goroutines - HATI-HATI!

### ❌ SALAH - Variable Capture Problem

```go
// INI SALAH! Jangan lakukan ini
func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Printf("Goroutine nomor: %d\n", i) // MASALAH DI SINI!
        }()
    }
    
    wg.Wait()
}
// Output mungkin: 4, 4, 4 (karena loop sudah selesai, i = 4)
```

### ✅ BENAR - Cara yang Tepat

```go
func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(nomor int) { // Pass sebagai parameter
            defer wg.Done()
            fmt.Printf("Goroutine nomor: %d\n", nomor)
        }(i) // Pass nilai i
    }
    
    wg.Wait()
}
// Output: 1, 2, 3 (dalam urutan acak)
```

## 10. Kapan Menggunakan Goroutines?

### ✅ Cocok untuk:
- **I/O Operations**: File reading, HTTP requests, database queries
- **Independent Tasks**: Tugas-tugas yang tidak saling bergantung
- **Parallel Processing**: Mengolah data besar secara parallel

### ❌ Tidak cocok untuk:
- **CPU-Intensive Sequential Tasks**: Perhitungan matematika yang harus berurutan
- **Shared Resources**: Ketika banyak goroutines mengakses data yang sama

## 11. Contoh Sederhana: Hitung Faktorial Parallel

```go
package main

import (
    "fmt"
    "sync"
)

func hitungFaktorial(n int, hasil *int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    faktorial := 1
    for i := 1; i <= n; i++ {
        faktorial *= i
    }
    
    *hasil = faktorial
    fmt.Printf("Faktorial %d = %d\n", n, faktorial)
}

func main() {
    var wg sync.WaitGroup
    
    // Hitung faktorial dari 1 sampai 5 secara bersamaan
    results := make([]int, 5)
    
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go hitungFaktorial(i, &results[i-1], &wg)
    }
    
    wg.Wait()
    
    fmt.Println("Semua perhitungan selesai!")
    fmt.Println("Hasil:", results)
}
```

## 12. Tips untuk Pemula

1. **Selalu gunakan WaitGroup** - Jangan andalkan `time.Sleep`
2. **Hati-hati dengan loop** - Selalu pass variable sebagai parameter
3. **Jangan buat terlalu banyak goroutines** - Ribuan oke, jutaan mungkin bermasalah
4. **Mulai sederhana** - Pahami konsep dulu sebelum ke yang kompleks

## 13. Debugging Goroutines

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("Worker %d mulai bekerja\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Worker %d selesai bekerja\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    fmt.Printf("Jumlah goroutines sebelum: %d\n", runtime.NumGoroutine())
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    
    fmt.Printf("Jumlah goroutines setelah spawn: %d\n", runtime.NumGoroutine())
    
    wg.Wait()
    
    fmt.Printf("Jumlah goroutines setelah selesai: %d\n", runtime.NumGoroutine())
}
```

## Kesimpulan

Goroutines adalah cara Go untuk menjalankan tugas secara bersamaan. Konsep penting:

1. Tambahkan `go` di depan function call
2. Gunakan `sync.WaitGroup` untuk menunggu goroutines selesai
3. Hati-hati dengan variable capture di loop
4. Mulai dengan contoh sederhana

Setelah paham konsep dasar ini, baru nanti kita bisa lanjut ke channels untuk komunikasi antar goroutines!