# Panduan Lengkap Context di Go untuk Pemula

### PZN MATERI: https://docs.google.com/presentation/d/1WhJvRpKPWq7LY9P6fMN93vkjKa1bJwBQebbieKdefPw/edit?slide=id.p#slide=id.p

## Apa itu Context?

Context di Go adalah package yang sangat penting untuk mengelola pembatalan operasi (cancellation), timeout, deadline, dan berbagi data antar goroutine. Context memungkinkan kita untuk mengontrol siklus hidup operasi yang berjalan secara concurrent.

**Konsep Dasar:**
- Context adalah interface yang menyediakan cara untuk membatalkan operasi
- Context dapat diteruskan antar fungsi dan goroutine
- Context membentuk hierarki parent-child
- Ketika parent context dibatalkan, semua child context juga akan dibatalkan

## Interface Context

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

### Penjelasan Method:
- **Deadline()**: Mengembalikan waktu kapan context akan expired
- **Done()**: Channel yang akan ditutup ketika context dibatalkan
- **Err()**: Mengembalikan error jika context dibatalkan
- **Value()**: Mengambil nilai yang disimpan dalam context

## Jenis-Jenis Context

### 1. Background Context
Context paling dasar yang tidak akan pernah dibatalkan.

```go
ctx := context.Background()
```

### 2. TODO Context  
Digunakan ketika kita belum tahu context apa yang akan digunakan.

```go
ctx := context.TODO()
```

### 3. WithCancel
Context yang bisa dibatalkan secara manual.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // Pastikan untuk memanggil cancel
```

### 4. WithTimeout
Context dengan batas waktu timeout.

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

### 5. WithDeadline
Context dengan deadline spesifik.

```go
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

### 6. WithValue
Context untuk menyimpan key-value pairs.

```go
ctx := context.WithValue(context.Background(), "userID", 123)
```

## Contoh Praktis dan Use Cases

### 1. Pembatalan Manual dengan WithCancel

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func doWork(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            // Context dibatalkan, hentikan pekerjaan
            fmt.Printf("%s: pekerjaan dibatalkan - %v\n", name, ctx.Err())
            return
        default:
            // Lanjutkan bekerja
            fmt.Printf("%s: sedang bekerja...\n", name)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    // Buat context yang bisa dibatalkan
    ctx, cancel := context.WithCancel(context.Background())
    
    // Jalankan goroutine
    go doWork(ctx, "Worker-1")
    go doWork(ctx, "Worker-2")
    
    // Biarkan worker berjalan sebentar
    time.Sleep(2 * time.Second)
    
    // Batalkan semua worker
    fmt.Println("Membatalkan semua worker...")
    cancel()
    
    // Tunggu sebentar untuk melihat pembatalan
    time.Sleep(1 * time.Second)
    fmt.Println("Program selesai")
}
```

**Penjelasan:**
- Kita membuat context yang bisa dibatalkan dengan `WithCancel`
- Setiap worker mengecek `ctx.Done()` dalam loop
- Ketika `cancel()` dipanggil, channel `Done()` akan ditutup
- Semua goroutine yang mendengarkan context akan menerima sinyal pembatalan

### 2. HTTP Request dengan Timeout

```go
package main

import (
    "context"
    "fmt"
    "io"
    "net/http"
    "time"
)

func fetchDataWithTimeout(url string, timeout time.Duration) {
    // Buat context dengan timeout
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    // Buat HTTP request dengan context
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        fmt.Printf("Error membuat request: %v\n", err)
        return
    }
    
    // Kirim request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        // Cek apakah error karena timeout
        if ctx.Err() == context.DeadlineExceeded {
            fmt.Printf("Request timeout setelah %v\n", timeout)
        } else {
            fmt.Printf("Error: %v\n", err)
        }
        return
    }
    defer resp.Body.Close()
    
    // Baca response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error membaca response: %v\n", err)
        return
    }
    
    fmt.Printf("Response berhasil, panjang: %d bytes\n", len(body))
}

func main() {
    fmt.Println("Mencoba request dengan timeout 2 detik...")
    fetchDataWithTimeout("https://httpbin.org/delay/1", 2*time.Second)
    
    fmt.Println("\nMencoba request dengan timeout 500ms...")
    fetchDataWithTimeout("https://httpbin.org/delay/1", 500*time.Millisecond)
}
```

**Penjelasan:**
- `WithTimeout` membuat context yang akan expired setelah durasi tertentu
- `http.NewRequestWithContext` mengikat request dengan context
- Jika request memakan waktu lebih lama dari timeout, akan dibatalkan otomatis
- Kita bisa mengecek `ctx.Err()` untuk mengetahui penyebab error

### 3. Database Query dengan Context

```go
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"
    
    // Contoh menggunakan sqlite driver
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    ID   int
    Name string
    Email string
}

type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*User, error) {
    // Query dengan context untuk timeout
    query := "SELECT id, name, email FROM users WHERE id = ?"
    
    var user User
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.Name, &user.Email,
    )
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            return nil, fmt.Errorf("query timeout: %w", ctx.Err())
        }
        return nil, err
    }
    
    return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) error {
    query := "INSERT INTO users (name, email) VALUES (?, ?)"
    
    result, err := r.db.ExecContext(ctx, query, user.Name, user.Email)
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            return fmt.Errorf("insert timeout: %w", ctx.Err())
        }
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    user.ID = int(id)
    return nil
}

func main() {
    // Setup database (contoh dengan SQLite)
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Buat tabel
    _, err = db.Exec(`
        CREATE TABLE users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL
        )
    `)
    if err != nil {
        log.Fatal(err)
    }
    
    repo := &UserRepository{db: db}
    
    // Context dengan timeout 5 detik
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // Buat user baru
    newUser := &User{
        Name:  "John Doe",
        Email: "john@example.com",
    }
    
    err = repo.CreateUser(ctx, newUser)
    if err != nil {
        log.Printf("Error creating user: %v", err)
        return
    }
    
    fmt.Printf("User berhasil dibuat dengan ID: %d\n", newUser.ID)
    
    // Ambil user berdasarkan ID
    user, err := repo.GetUserByID(ctx, newUser.ID)
    if err != nil {
        log.Printf("Error getting user: %v", err)
        return
    }
    
    fmt.Printf("User ditemukan: %+v\n", user)
}
```

**Penjelasan:**
- Database operations menggunakan methods yang berakhiran `Context`
- `QueryRowContext`, `ExecContext` menerima context sebagai parameter pertama
- Jika query memakan waktu terlalu lama, akan dibatalkan sesuai timeout context

### 4. Context dengan Values

```go
package main

import (
    "context"
    "fmt"
    "log"
)

// Definisikan custom type untuk key agar type-safe
type contextKey string

const (
    UserIDKey    contextKey = "userID"
    RequestIDKey contextKey = "requestID"
)

func processRequest(ctx context.Context) {
    // Ambil nilai dari context
    userID, ok := ctx.Value(UserIDKey).(int)
    if !ok {
        log.Println("UserID tidak ditemukan dalam context")
        return
    }
    
    requestID, ok := ctx.Value(RequestIDKey).(string)
    if !ok {
        log.Println("RequestID tidak ditemukan dalam context")
        return
    }
    
    fmt.Printf("Memproses request %s untuk user %d\n", requestID, userID)
    
    // Lanjutkan ke fungsi lain dengan context yang sama
    authenticateUser(ctx)
    fetchUserData(ctx)
}

func authenticateUser(ctx context.Context) {
    userID := ctx.Value(UserIDKey).(int)
    requestID := ctx.Value(RequestIDKey).(string)
    
    fmt.Printf("[%s] Authenticating user %d\n", requestID, userID)
    // Simulasi autentikasi
}

func fetchUserData(ctx context.Context) {
    userID := ctx.Value(UserIDKey).(int)
    requestID := ctx.Value(RequestIDKey).(string)
    
    fmt.Printf("[%s] Fetching data for user %d\n", requestID, userID)
    // Simulasi pengambilan data
}

func main() {
    // Buat base context
    ctx := context.Background()
    
    // Tambahkan values ke context
    ctx = context.WithValue(ctx, UserIDKey, 123)
    ctx = context.WithValue(ctx, RequestIDKey, "req-456")
    
    // Proses request dengan context yang berisi data
    processRequest(ctx)
}
```

**Penjelasan:**
- `WithValue` menyimpan key-value pairs dalam context
- Gunakan custom type untuk key agar type-safe dan menghindari collision
- Values dalam context bersifat immutable dan dapat diakses di seluruh call stack

### 5. Real-world Example: Web Server dengan Context

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "time"
)

type contextKey string

const RequestIDKey contextKey = "requestID"

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Middleware untuk menambahkan request ID
func requestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        requestID := fmt.Sprintf("req-%d", time.Now().UnixNano())
        
        // Tambahkan request ID ke context
        ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
        r = r.WithContext(ctx)
        
        // Tambahkan request ID ke response header
        w.Header().Set("X-Request-ID", requestID)
        
        next(w, r)
    }
}

// Middleware untuk timeout
func timeoutMiddleware(timeout time.Duration) func(http.HandlerFunc) http.HandlerFunc {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            ctx, cancel := context.WithTimeout(r.Context(), timeout)
            defer cancel()
            
            r = r.WithContext(ctx)
            next(w, r)
        }
    }
}

// Simulasi database service
func getUserFromDB(ctx context.Context, userID int) (*User, error) {
    requestID := ctx.Value(RequestIDKey).(string)
    log.Printf("[%s] Fetching user %d from database", requestID, userID)
    
    // Simulasi query database yang lambat
    select {
    case <-time.After(2 * time.Second):
        return &User{ID: userID, Name: fmt.Sprintf("User%d", userID)}, nil
    case <-ctx.Done():
        return nil, ctx.Err()
    }
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    requestID := r.Context().Value(RequestIDKey).(string)
    log.Printf("[%s] Handling GET /user request", requestID)
    
    // Parse user ID dari URL
    userIDStr := r.URL.Query().Get("id")
    if userIDStr == "" {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }
    
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    // Fetch user dengan context
    user, err := getUserFromDB(r.Context(), userID)
    if err != nil {
        if err == context.DeadlineExceeded {
            log.Printf("[%s] Request timeout", requestID)
            http.Error(w, "Request timeout", http.StatusRequestTimeout)
        } else {
            log.Printf("[%s] Database error: %v", requestID, err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    
    // Response dengan JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
    
    log.Printf("[%s] Request completed successfully", requestID)
}

func main() {
    // Setup routes dengan middleware
    http.HandleFunc("/user", 
        requestIDMiddleware(
            timeoutMiddleware(5*time.Second)(getUserHandler),
        ),
    )
    
    fmt.Println("Server starting on :8080")
    fmt.Println("Try: curl 'http://localhost:8080/user?id=123'")
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Penjelasan:**
- Middleware menambahkan request ID dan timeout ke context
- Context diteruskan dari HTTP handler ke database service
- Jika database query terlalu lama, request akan di-timeout
- Request ID membantu tracking request di log

## Best Practices

### 1. Selalu Pass Context sebagai Parameter Pertama
```go
// ✅ Benar
func processData(ctx context.Context, data string) error {
    // ...
}

// ❌ Salah - context bukan parameter pertama
func processData(data string, ctx context.Context) error {
    // ...
}
```

### 2. Jangan Simpan Context di Struct
```go
// ❌ Salah - jangan simpan context di struct
type Service struct {
    ctx context.Context
    db  *sql.DB
}

// ✅ Benar - pass context ke method
type Service struct {
    db *sql.DB
}

func (s *Service) GetData(ctx context.Context, id int) error {
    // ...
}
```

### 3. Selalu Panggil Cancel Function
```go
// ✅ Benar - selalu defer cancel
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Gunakan ctx...
```

### 4. Gunakan Custom Type untuk Context Keys
```go
// ✅ Benar - custom type mencegah collision
type contextKey string
const UserIDKey contextKey = "userID"

// ❌ Salah - string literal rentan collision
ctx = context.WithValue(ctx, "userID", 123)
```

### 5. Jangan Gunakan WithValue untuk Data yang Sering Diakses
```go
// ❌ Salah - untuk data yang sering dibutuhkan
func processUser(ctx context.Context) {
    userID := ctx.Value("userID").(int) // Setiap saat perlu extract
    // ...
}

// ✅ Benar - pass sebagai parameter eksplisit
func processUser(ctx context.Context, userID int) {
    // ...
}
```

## Kesalahan Umum yang Harus Dihindari

### 1. Menggunakan nil Context
```go
// ❌ Salah
doSomething(nil)

// ✅ Benar
doSomething(context.Background())
```

### 2. Tidak Menangani Context Cancellation
```go
// ❌ Salah - tidak mengecek cancellation
func longRunningTask(ctx context.Context) {
    for i := 0; i < 1000000; i++ {
        // Pekerjaan berat tanpa cek cancellation
        heavyWork(i)
    }
}

// ✅ Benar - selalu cek cancellation
func longRunningTask(ctx context.Context) {
    for i := 0; i < 1000000; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            heavyWork(i)
        }
    }
}
```

### 3. Context Leak
```go
// ❌ Salah - tidak memanggil cancel
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// cancel tidak dipanggil -> memory leak

// ✅ Benar
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel() // Pastikan cancel dipanggil
```

## Kapan Menggunakan Context

**Gunakan Context untuk:**
- HTTP requests dengan timeout
- Database operations yang bisa di-cancel
- Long-running operations yang perlu dibatalkan
- Passing request-scoped values (user ID, request ID, trace ID)
- Koordinasi antar goroutine

**Jangan Gunakan Context untuk:**
- Optional parameters ke fungsi
- Data yang seharusnya jadi struct fields
- Data yang tidak terkait dengan request lifecycle

## Penutup

Context adalah tool yang sangat powerful di Go untuk managing concurrent operations. Dengan memahami konsep dan best practices di atas, Anda dapat menulis aplikasi Go yang lebih robust dan responsive. Ingat untuk selalu:

1. Pass context sebagai parameter pertama
2. Selalu panggil cancel function
3. Handle context cancellation dengan proper
4. Gunakan context values dengan bijak
5. Jangan simpan context di struct

Praktikkan contoh-contoh di atas dan Anda akan menguasai context di Go dengan baik!