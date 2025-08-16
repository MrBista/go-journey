# Panduan Lengkap File I/O di Go untuk Pemula

## 1. Pengenalan File I/O di Go

File I/O (Input/Output) adalah proses membaca dan menulis data dari/ke file di sistem operasi. Di Go, operasi file I/O terutama menggunakan package `os`, `io`, dan `ioutil` (deprecated di Go 1.16+, diganti dengan `io` dan `os`).

### Package Utama untuk File I/O:
- **`os`**: Menyediakan interface untuk berinteraksi dengan sistem operasi
- **`io`**: Menyediakan primitif dasar untuk I/O operations
- **`bufio`**: Menyediakan buffered I/O untuk performa yang lebih baik
- **`filepath`**: Untuk manipulasi path file

## 2. Membaca File di Go

### 2.1 Membaca Seluruh File dengan `os.ReadFile`

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Membaca seluruh isi file sekaligus
    content, err := os.ReadFile("example.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    // content bertipe []byte, konversi ke string untuk display
    fmt.Printf("File content:\n%s", string(content))
}
```

**Penjelasan:**
- `os.ReadFile()` membaca seluruh file ke dalam memory sekaligus
- Return value: `[]byte` dan `error`
- Cocok untuk file kecil (< 100MB)
- Otomatis menutup file setelah selesai membaca

### 2.2 Membaca File dengan `os.Open` dan Manual Close

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // Buka file
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    
    // PENTING: Selalu tutup file untuk mencegah memory leak
    defer file.Close()
    
    // Baca seluruh isi file
    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    
    fmt.Printf("File content:\n%s", string(content))
}
```

**Penjelasan:**
- `os.Open()` membuka file untuk dibaca saja
- `defer file.Close()` memastikan file ditutup ketika function selesai
- `io.ReadAll()` membaca seluruh isi dari reader

### 2.3 Membaca File Baris per Baris dengan `bufio.Scanner`

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Buat scanner untuk membaca baris per baris
    scanner := bufio.NewScanner(file)
    
    lineNumber := 1
    for scanner.Scan() {
        line := scanner.Text() // Mendapat string dari baris current
        fmt.Printf("Line %d: %s\n", lineNumber, line)
        lineNumber++
    }
    
    // Cek error setelah scanning selesai
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}
```

**Penjelasan:**
- `bufio.Scanner` efisien untuk membaca file besar baris per baris
- `scanner.Scan()` return `true` jika masih ada baris untuk dibaca
- `scanner.Text()` return string dari baris current
- Memory efficient karena hanya load satu baris di memory

### 2.4 Membaca File dengan Buffer Size Custom

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Buffer 1024 bytes
    buffer := make([]byte, 1024)
    
    for {
        // Baca ke buffer
        n, err := file.Read(buffer)
        if err != nil {
            if err == io.EOF {
                fmt.Println("Reached end of file")
                break
            }
            fmt.Printf("Error reading: %v\n", err)
            return
        }
        
        // Process data yang dibaca (n bytes)
        fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
    }
}
```

**Penjelasan:**
- Cocok untuk file sangat besar yang tidak bisa dimuat di memory
- `io.EOF` menandakan end of file
- `buffer[:n]` hanya mengambil data yang actual dibaca

## 3. Menulis File di Go

### 3.1 Menulis File dengan `os.WriteFile`

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    content := "Hello, World!\nThis is a new file created with Go."
    
    // Tulis ke file, buat file baru jika tidak ada
    // Permission 0644: owner read+write, group+others read only
    err := os.WriteFile("output.txt", []byte(content), 0644)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }
    
    fmt.Println("File written successfully!")
}
```

**Penjelasan:**
- `os.WriteFile()` menulis seluruh data sekaligus
- Otomatis membuat file jika belum ada
- Permission `0644` adalah octal notation untuk file permissions
- Menimpa file jika sudah ada

### 3.2 Menulis File dengan `os.Create` dan Manual Write

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Buat file baru (atau timpa jika sudah ada)
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Tulis string ke file
    content := "Hello from Go!\n"
    n, err := file.WriteString(content)
    if err != nil {
        fmt.Printf("Error writing to file: %v\n", err)
        return
    }
    
    fmt.Printf("Successfully wrote %d bytes to file\n", n)
}
```

### 3.3 Append ke File yang Sudah Ada

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Buka file dengan flag O_APPEND untuk menambah di akhir
    file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Tambah content di akhir file
    newContent := "This line is appended!\n"
    n, err := file.WriteString(newContent)
    if err != nil {
        fmt.Printf("Error appending to file: %v\n", err)
        return
    }
    
    fmt.Printf("Successfully appended %d bytes to file\n", n)
}
```

**Penjelasan Flag os.OpenFile:**
- `os.O_APPEND`: Tulis di akhir file
- `os.O_WRONLY`: Buka untuk write saja
- `os.O_CREATE`: Buat file jika tidak ada
- `os.O_RDWR`: Read dan write
- `os.O_TRUNC`: Truncate (kosongkan) file jika ada

### 3.4 Buffered Writing untuk Performa

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("buffered_output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Buat buffered writer
    writer := bufio.NewWriter(file)
    defer writer.Flush() // PENTING: Flush buffer ke file
    
    // Tulis multiple lines
    lines := []string{
        "Line 1: Buffered writing is efficient",
        "Line 2: For multiple write operations",
        "Line 3: Buffer collects data before writing to disk",
    }
    
    for i, line := range lines {
        n, err := writer.WriteString(fmt.Sprintf("%s\n", line))
        if err != nil {
            fmt.Printf("Error writing line %d: %v\n", i+1, err)
            return
        }
        fmt.Printf("Buffered %d bytes for line %d\n", n, i+1)
    }
    
    fmt.Println("All data written to buffer and flushed to file!")
}
```

**Penjelasan:**
- `bufio.NewWriter()` membuat buffer untuk write operations
- Data dikumpulkan di buffer sebelum ditulis ke disk
- `writer.Flush()` memaksa menulis buffer ke disk
- Lebih efisien untuk multiple write operations

## 4. Working dengan CSV Files

### 4.1 Membaca CSV File

```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    // Buka CSV file
    file, err := os.Open("data.csv")
    if err != nil {
        fmt.Printf("Error opening CSV file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Buat CSV reader
    reader := csv.NewReader(file)
    
    // Baca semua records
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Printf("Error reading CSV: %v\n", err)
        return
    }
    
    // Process each record
    for i, record := range records {
        fmt.Printf("Row %d: %v\n", i+1, record)
    }
}
```

### 4.2 Menulis CSV File

```go
package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    // Buat CSV file
    file, err := os.Create("output.csv")
    if err != nil {
        fmt.Printf("Error creating CSV file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Buat CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // Data untuk ditulis
    data := [][]string{
        {"Name", "Age", "City"},
        {"John Doe", "30", "New York"},
        {"Jane Smith", "25", "Los Angeles"},
        {"Bob Johnson", "35", "Chicago"},
    }
    
    // Tulis semua records
    err = writer.WriteAll(data)
    if err != nil {
        fmt.Printf("Error writing CSV: %v\n", err)
        return
    }
    
    fmt.Println("CSV file written successfully!")
}
```

## 5. Working dengan JSON Files

### 5.1 Membaca JSON File

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// Struct untuk data JSON
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    // Baca JSON file
    content, err := os.ReadFile("person.json")
    if err != nil {
        fmt.Printf("Error reading JSON file: %v\n", err)
        return
    }
    
    // Parse JSON ke struct
    var person Person
    err = json.Unmarshal(content, &person)
    if err != nil {
        fmt.Printf("Error parsing JSON: %v\n", err)
        return
    }
    
    fmt.Printf("Parsed Person: %+v\n", person)
}
```

### 5.2 Menulis JSON File

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    // Data untuk ditulis
    person := Person{
        Name:  "John Doe",
        Age:   30,
        Email: "john@example.com",
    }
    
    // Convert struct ke JSON dengan indentasi
    jsonData, err := json.MarshalIndent(person, "", "  ")
    if err != nil {
        fmt.Printf("Error marshaling JSON: %v\n", err)
        return
    }
    
    // Tulis ke file
    err = os.WriteFile("output.json", jsonData, 0644)
    if err != nil {
        fmt.Printf("Error writing JSON file: %v\n", err)
        return
    }
    
    fmt.Println("JSON file written successfully!")
}
```

## 6. File Operations Utility Functions

### 6.1 Cek Keberadaan File

```go
package main

import (
    "fmt"
    "os"
)

func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return !os.IsNotExist(err)
}

func main() {
    filename := "example.txt"
    
    if fileExists(filename) {
        fmt.Printf("File %s exists\n", filename)
    } else {
        fmt.Printf("File %s does not exist\n", filename)
    }
}
```

### 6.2 Copy File

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func copyFile(src, dst string) error {
    // Buka source file
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()
    
    // Buat destination file
    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()
    
    // Copy content
    _, err = io.Copy(destFile, sourceFile)
    return err
}

func main() {
    err := copyFile("source.txt", "destination.txt")
    if err != nil {
        fmt.Printf("Error copying file: %v\n", err)
        return
    }
    
    fmt.Println("File copied successfully!")
}
```

### 6.3 Delete File

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    filename := "file_to_delete.txt"
    
    err := os.Remove(filename)
    if err != nil {
        fmt.Printf("Error deleting file: %v\n", err)
        return
    }
    
    fmt.Printf("File %s deleted successfully!\n", filename)
}
```

## 7. Best Practices untuk File I/O di Go

### 7.1 Selalu Handle Error

```go
// ❌ Bad: Ignore error
content, _ := os.ReadFile("file.txt")

// ✅ Good: Handle error properly
content, err := os.ReadFile("file.txt")
if err != nil {
    log.Printf("Error reading file: %v", err)
    return
}
```

### 7.2 Gunakan defer untuk Close Resource

```go
// ✅ Good: Gunakan defer
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close() // Akan dipanggil ketika function selesai

// ... use file
```

### 7.3 Pilih Method yang Tepat berdasarkan Use Case

```go
// Untuk file kecil (< 100MB)
content, err := os.ReadFile("small_file.txt")

// Untuk file besar, baca baris per baris
file, err := os.Open("large_file.txt")
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    processLine(scanner.Text())
}

// Untuk multiple write operations, gunakan buffered writer
writer := bufio.NewWriter(file)
// ... multiple writes
writer.Flush()
```

### 7.4 Set Permission yang Tepat

```go
// 0644: owner read+write, group+others read only (recommended for data files)
os.WriteFile("data.txt", data, 0644)

// 0755: owner read+write+execute, group+others read+execute (untuk executable)
os.WriteFile("script.sh", script, 0755)

// 0600: owner read+write only (untuk sensitive data)
os.WriteFile("secrets.txt", secrets, 0600)
```

### 7.5 Validate Input Path

```go
package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func validatePath(path string) error {
    // Clean path untuk menghindari path traversal
    cleanPath := filepath.Clean(path)
    
    // Cek path traversal attempt
    if strings.Contains(cleanPath, "..") {
        return fmt.Errorf("invalid path: path traversal detected")
    }
    
    return nil
}
```

## 8. Common Pitfalls dan Solusinya

### 8.1 Memory Leak karena tidak Close File

```go
// ❌ Bad: File descriptor leak
func badReadFile() {
    file, _ := os.Open("file.txt")
    // Missing defer file.Close()
    content, _ := io.ReadAll(file)
    return content
}

// ✅ Good: Always close
func goodReadFile() ([]byte, error) {
    file, err := os.Open("file.txt")
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    return io.ReadAll(file)
}
```

### 8.2 Buffer tidak di-Flush

```go
// ❌ Bad: Buffer tidak di-flush
writer := bufio.NewWriter(file)
writer.WriteString("data")
// Data mungkin masih di buffer, tidak tertulis ke file

// ✅ Good: Flush buffer
writer := bufio.NewWriter(file)
defer writer.Flush() // Atau panggil manual sebelum return
writer.WriteString("data")
```

### 8.3 Race Condition saat Concurrent File Access

```go
package main

import (
    "sync"
)

// ✅ Good: Gunakan mutex untuk concurrent access
type SafeFileWriter struct {
    mu   sync.Mutex
    file *os.File
}

func (w *SafeFileWriter) Write(data []byte) error {
    w.mu.Lock()
    defer w.mu.Unlock()
    
    _, err := w.file.Write(data)
    return err
}
```

## 9. Contoh Aplikasi Praktis: Log File Rotator

```go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

type LogRotator struct {
    basePath  string
    maxSize   int64 // bytes
    currentFile *os.File
}

func NewLogRotator(basePath string, maxSize int64) *LogRotator {
    return &LogRotator{
        basePath: basePath,
        maxSize:  maxSize,
    }
}

func (lr *LogRotator) Write(data []byte) error {
    // Cek apakah perlu rotate file
    if lr.needsRotation() {
        if err := lr.rotate(); err != nil {
            return err
        }
    }
    
    // Pastikan file terbuka
    if lr.currentFile == nil {
        if err := lr.openCurrentFile(); err != nil {
            return err
        }
    }
    
    // Tulis data
    _, err := lr.currentFile.Write(data)
    return err
}

func (lr *LogRotator) needsRotation() bool {
    if lr.currentFile == nil {
        return true
    }
    
    // Cek ukuran file
    info, err := lr.currentFile.Stat()
    if err != nil {
        return true
    }
    
    return info.Size() >= lr.maxSize
}

func (lr *LogRotator) rotate() error {
    // Tutup file current
    if lr.currentFile != nil {
        lr.currentFile.Close()
        lr.currentFile = nil
    }
    
    // Rename current log file dengan timestamp
    currentPath := lr.basePath
    if _, err := os.Stat(currentPath); err == nil {
        timestamp := time.Now().Format("2006-01-02_15-04-05")
        ext := filepath.Ext(currentPath)
        name := currentPath[:len(currentPath)-len(ext)]
        rotatedPath := fmt.Sprintf("%s_%s%s", name, timestamp, ext)
        
        if err := os.Rename(currentPath, rotatedPath); err != nil {
            return err
        }
    }
    
    return lr.openCurrentFile()
}

func (lr *LogRotator) openCurrentFile() error {
    file, err := os.OpenFile(lr.basePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    
    lr.currentFile = file
    return nil
}

func (lr *LogRotator) Close() error {
    if lr.currentFile != nil {
        return lr.currentFile.Close()
    }
    return nil
}

func main() {
    // Contoh penggunaan
    rotator := NewLogRotator("app.log", 1024) // Rotate setiap 1KB
    defer rotator.Close()
    
    // Tulis log entries
    for i := 0; i < 100; i++ {
        logEntry := fmt.Sprintf("[%s] Log entry %d\n", 
            time.Now().Format("2006-01-02 15:04:05"), i)
        
        if err := rotator.Write([]byte(logEntry)); err != nil {
            fmt.Printf("Error writing log: %v\n", err)
            break
        }
    }
    
    fmt.Println("Log rotation example completed!")
}
```

## 10. Testing File I/O Operations

```go
package main

import (
    "os"
    "path/filepath"
    "testing"
)

func TestFileOperations(t *testing.T) {
    // Buat temporary directory untuk testing
    tmpDir := t.TempDir()
    testFile := filepath.Join(tmpDir, "test.txt")
    
    // Test write
    content := "test content"
    err := os.WriteFile(testFile, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to write test file: %v", err)
    }
    
    // Test read
    readContent, err := os.ReadFile(testFile)
    if err != nil {
        t.Fatalf("Failed to read test file: %v", err)
    }
    
    if string(readContent) != content {
        t.Errorf("Expected %q, got %q", content, string(readContent))
    }
    
    // Test file exists
    if _, err := os.Stat(testFile); os.IsNotExist(err) {
        t.Error("File should exist after writing")
    }
    
    // File akan otomatis dihapus karena menggunakan t.TempDir()
}
```

---

## Kesimpulan

File I/O di Go menyediakan berbagai cara untuk berinteraksi dengan file system:

1. **Untuk file kecil**: Gunakan `os.ReadFile()` dan `os.WriteFile()`
2. **Untuk file besar**: Gunakan `bufio.Scanner` untuk membaca dan `bufio.Writer` untuk menulis
3. **Selalu handle error** dan gunakan `defer` untuk close resources
4. **Pilih method yang tepat** berdasarkan use case Anda
5. **Validate input path** untuk keamanan
6. **Gunakan buffered I/O** untuk performa yang lebih baik pada multiple operations

Dengan memahami konsep-konsep ini, Anda akan dapat menangani berbagai skenario file I/O dengan efisien dan aman di Go!