# Panduan Lengkap MySQL dengan Go untuk Pemula

### PZN MATERI: https://docs.google.com/presentation/d/15pvN3L3HTgA9aIMNkm03PzzIwlff0WDE6hOWWut9pg8/edit?slide=id.gbbaf426c8a_0_361#slide=id.gbbaf426c8a_0_361

## Daftar Isi
1. [Persiapan dan Setup](#persiapan-dan-setup)
2. [Koneksi Database](#koneksi-database)
3. [CRUD Operations](#crud-operations)
4. [Prepared Statements](#prepared-statements)
5. [Transactions](#transactions)
6. [Connection Pooling](#connection-pooling)
7. [Error Handling](#error-handling)
8. [Best Practices](#best-practices)
9. [Studi Kasus: Aplikasi To-Do List](#studi-kasus)

---

## Persiapan dan Setup

### 1. Instalasi Driver MySQL untuk Go

Pertama, kita perlu menginstall driver MySQL. Driver yang paling populer dan direkomendasikan adalah `go-sql-driver/mysql`.

```bash
go mod init mysql-tutorial
go get -u github.com/go-sql-driver/mysql
```

**Penjelasan:**
- `go mod init` membuat module Go baru
- `go get` mengunduh dan menginstall package MySQL driver
- Flag `-u` memastikan kita mendapat versi terbaru

### 2. Setup Database MySQL

Buat database dan tabel untuk tutorial ini:

```sql
CREATE DATABASE tutorial_go;
USE tutorial_go;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    age INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

---

## Koneksi Database

### Koneksi Dasar

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/go-sql-driver/mysql" // Import driver secara anonymous
)

func main() {
    // Data Source Name (DSN) - format koneksi ke database
    dsn := "username:password@tcp(localhost:3306)/tutorial_go"
    
    // Membuka koneksi ke database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error membuka koneksi:", err)
    }
    defer db.Close() // Pastikan koneksi ditutup saat program selesai
    
    // Mengetes koneksi
    err = db.Ping()
    if err != nil {
        log.Fatal("Error koneksi ke database:", err)
    }
    
    fmt.Println("Koneksi ke database berhasil!")
}
```

**Penjelasan Detail:**

1. **Import Anonymous Driver**: `_ "github.com/go-sql-driver/mysql"`
   - Tanda underscore (`_`) berarti kita hanya ingin menjalankan init function dari package
   - Driver akan mendaftarkan dirinya ke `database/sql` package

2. **DSN (Data Source Name)**:
   - Format: `username:password@tcp(host:port)/database_name`
   - Bisa ditambahkan parameter seperti `?parseTime=true&charset=utf8mb4`

3. **sql.Open()**: Tidak langsung membuka koneksi, hanya memvalidasi DSN
4. **db.Ping()**: Benar-benar mengetes koneksi ke database
5. **defer db.Close()**: Memastikan koneksi ditutup saat function selesai

### Koneksi dengan Konfigurasi Lengkap

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    
    _ "github.com/go-sql-driver/mysql"
)

type Database struct {
    DB *sql.DB
}

func NewDatabase() (*Database, error) {
    // DSN dengan parameter tambahan
    dsn := "root:password@tcp(localhost:3306)/tutorial_go?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("error membuka database: %w", err)
    }
    
    // Konfigurasi connection pool
    db.SetMaxOpenConns(25)    // Maximum 25 koneksi terbuka
    db.SetMaxIdleConns(25)    // Maximum 25 koneksi idle
    db.SetConnMaxLifetime(5 * time.Minute) // Koneksi akan di-refresh setiap 5 menit
    
    // Test koneksi
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error ping database: %w", err)
    }
    
    return &Database{DB: db}, nil
}

func main() {
    database, err := NewDatabase()
    if err != nil {
        log.Fatal(err)
    }
    defer database.DB.Close()
    
    fmt.Println("Database terhubung dengan sukses!")
}
```

**Penjelasan Parameter DSN:**
- `parseTime=true`: Mengkonversi DATE dan DATETIME MySQL ke Go's time.Time
- `charset=utf8mb4`: Mendukung emoji dan karakter Unicode penuh
- `collation=utf8mb4_unicode_ci`: Aturan perbandingan string

---

## CRUD Operations

### Create (INSERT)

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

func createUser(db *sql.DB, user User) (int, error) {
    // Query INSERT
    query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
    
    // Eksekusi query dan dapatkan result
    result, err := db.Exec(query, user.Name, user.Email, user.Age)
    if err != nil {
        return 0, fmt.Errorf("error creating user: %w", err)
    }
    
    // Mendapatkan ID yang baru saja di-insert
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("error getting last insert id: %w", err)
    }
    
    return int(id), nil
}

func main() {
    // Koneksi database (menggunakan kode sebelumnya)
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/tutorial_go")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Membuat user baru
    newUser := User{
        Name:  "John Doe",
        Email: "john@example.com",
        Age:   25,
    }
    
    userID, err := createUser(db, newUser)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("User berhasil dibuat dengan ID: %d\n", userID)
}
```

**Penjelasan:**
- `db.Exec()` digunakan untuk query yang tidak mengembalikan rows (INSERT, UPDATE, DELETE)
- `?` adalah placeholder untuk mencegah SQL injection
- `result.LastInsertId()` mendapatkan ID auto-increment yang baru dibuat
- Parameter dikirim secara terpisah untuk keamanan

### Read (SELECT)

#### Select Single Row

```go
func getUserByID(db *sql.DB, userID int) (*User, error) {
    var user User
    
    // Query SELECT untuk satu row
    query := "SELECT id, name, email, age FROM users WHERE id = ?"
    
    // QueryRow mengembalikan maksimal satu row
    row := db.QueryRow(query, userID)
    
    // Scan hasil query ke struct
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user dengan ID %d tidak ditemukan", userID)
        }
        return nil, fmt.Errorf("error scanning user: %w", err)
    }
    
    return &user, nil
}

func main() {
    // ... kode koneksi database
    
    user, err := getUserByID(db, 1)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("User ditemukan: %+v\n", user)
}
```

#### Select Multiple Rows

```go
func getAllUsers(db *sql.DB) ([]User, error) {
    var users []User
    
    // Query untuk mendapatkan semua users
    query := "SELECT id, name, email, age FROM users ORDER BY created_at DESC"
    
    // Query mengembalikan multiple rows
    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error querying users: %w", err)
    }
    defer rows.Close() // Penting! Tutup rows setelah selesai
    
    // Iterasi melalui setiap row
    for rows.Next() {
        var user User
        
        // Scan setiap row ke struct
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
        if err != nil {
            return nil, fmt.Errorf("error scanning user: %w", err)
        }
        
        users = append(users, user)
    }
    
    // Cek apakah ada error selama iterasi
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %w", err)
    }
    
    return users, nil
}

func main() {
    // ... kode koneksi database
    
    users, err := getAllUsers(db)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Ditemukan %d users:\n", len(users))
    for _, user := range users {
        fmt.Printf("- %s (%s), umur %d\n", user.Name, user.Email, user.Age)
    }
}
```

**Penjelasan Penting:**
- `db.QueryRow()` untuk mendapatkan maksimal satu row
- `db.Query()` untuk mendapatkan multiple rows
- **Selalu tutup `rows` dengan `defer rows.Close()`**
- `rows.Next()` return false ketika tidak ada row lagi
- `rows.Err()` mengecek error yang terjadi selama iterasi
- `sql.ErrNoRows` adalah error khusus ketika tidak ada data ditemukan

### Update (UPDATE)

```go
func updateUser(db *sql.DB, userID int, name, email string, age int) error {
    // Query UPDATE
    query := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"
    
    // Eksekusi query
    result, err := db.Exec(query, name, email, age, userID)
    if err != nil {
        return fmt.Errorf("error updating user: %w", err)
    }
    
    // Cek berapa row yang terpengaruh
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    
    // Jika tidak ada row yang terpengaruh, berarti user tidak ditemukan
    if rowsAffected == 0 {
        return fmt.Errorf("user dengan ID %d tidak ditemukan", userID)
    }
    
    fmt.Printf("%d user berhasil diupdate\n", rowsAffected)
    return nil
}

func main() {
    // ... kode koneksi database
    
    err := updateUser(db, 1, "John Smith", "johnsmith@example.com", 26)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("User berhasil diupdate!")
}
```

### Delete (DELETE)

```go
func deleteUser(db *sql.DB, userID int) error {
    // Query DELETE
    query := "DELETE FROM users WHERE id = ?"
    
    // Eksekusi query
    result, err := db.Exec(query, userID)
    if err != nil {
        return fmt.Errorf("error deleting user: %w", err)
    }
    
    // Cek berapa row yang terhapus
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("user dengan ID %d tidak ditemukan", userID)
    }
    
    fmt.Printf("%d user berhasil dihapus\n", rowsAffected)
    return nil
}

// Soft delete alternative
func softDeleteUser(db *sql.DB, userID int) error {
    query := "UPDATE users SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL"
    
    result, err := db.Exec(query, userID)
    if err != nil {
        return fmt.Errorf("error soft deleting user: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error getting rows affected: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("user dengan ID %d tidak ditemukan atau sudah dihapus", userID)
    }
    
    return nil
}
```

---

## Prepared Statements

Prepared statements sangat penting untuk performa dan keamanan aplikasi database.

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/go-sql-driver/mysql"
)

func demonstratePreparedStatements(db *sql.DB) error {
    // Prepare statement sekali
    stmt, err := db.Prepare("SELECT id, name, email FROM users WHERE age > ? AND name LIKE ?")
    if err != nil {
        return fmt.Errorf("error preparing statement: %w", err)
    }
    defer stmt.Close() // Jangan lupa tutup statement
    
    // Gunakan statement berkali-kali dengan parameter berbeda
    testCases := []struct {
        minAge int
        namePattern string
    }{
        {20, "%John%"},
        {25, "%Jane%"},
        {30, "%Bob%"},
    }
    
    for _, tc := range testCases {
        fmt.Printf("\nMencari users dengan umur > %d dan nama mengandung '%s':\n", 
                   tc.minAge, tc.namePattern)
        
        // Eksekusi prepared statement
        rows, err := stmt.Query(tc.minAge, tc.namePattern)
        if err != nil {
            return fmt.Errorf("error executing prepared statement: %w", err)
        }
        
        for rows.Next() {
            var id int
            var name, email string
            
            if err := rows.Scan(&id, &name, &email); err != nil {
                rows.Close()
                return fmt.Errorf("error scanning row: %w", err)
            }
            
            fmt.Printf("  - %s (%s)\n", name, email)
        }
        
        rows.Close()
        
        if err := rows.Err(); err != nil {
            return fmt.Errorf("error iterating rows: %w", err)
        }
    }
    
    return nil
}

// Prepared statement untuk batch insert
func batchInsertUsers(db *sql.DB, users []User) error {
    // Prepare statement untuk insert
    stmt, err := db.Prepare("INSERT INTO users (name, email, age) VALUES (?, ?, ?)")
    if err != nil {
        return fmt.Errorf("error preparing insert statement: %w", err)
    }
    defer stmt.Close()
    
    // Insert banyak users dengan loop
    for _, user := range users {
        result, err := stmt.Exec(user.Name, user.Email, user.Age)
        if err != nil {
            return fmt.Errorf("error inserting user %s: %w", user.Name, err)
        }
        
        id, _ := result.LastInsertId()
        fmt.Printf("User %s berhasil dibuat dengan ID: %d\n", user.Name, id)
    }
    
    return nil
}

func main() {
    // Koneksi database
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/tutorial_go")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Test prepared statements
    err = demonstratePreparedStatements(db)
    if err != nil {
        log.Fatal(err)
    }
    
    // Test batch insert
    newUsers := []User{
        {"Alice Johnson", "alice@example.com", 28},
        {"Bob Wilson", "bob@example.com", 32},
        {"Charlie Brown", "charlie@example.com", 24},
    }
    
    err = batchInsertUsers(db, newUsers)
    if err != nil {
        log.Fatal(err)
    }
}
```

**Keuntungan Prepared Statements:**
1. **Keamanan**: Mencegah SQL injection secara otomatis
2. **Performa**: Query di-compile sekali, digunakan berkali-kali
3. **Efisiensi**: Mengurangi parsing overhead di database
4. **Type Safety**: Parameter otomatis di-escape sesuai tipe data

---

## Transactions

Transaksi memastikan konsistensi data dengan prinsip ACID.

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/go-sql-driver/mysql"
)

func transferMoney(db *sql.DB, fromUserID, toUserID int, amount float64) error {
    // Mulai transaksi
    tx, err := db.Begin()
    if err != nil {
        return fmt.Errorf("error starting transaction: %w", err)
    }
    
    // Defer rollback - akan dipanggil jika terjadi error
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p) // Re-throw panic setelah rollback
        } else if err != nil {
            tx.Rollback()
        }
    }()
    
    // 1. Cek saldo user pengirim
    var senderBalance float64
    err = tx.QueryRow("SELECT balance FROM users WHERE id = ?", fromUserID).Scan(&senderBalance)
    if err != nil {
        return fmt.Errorf("error getting sender balance: %w", err)
    }
    
    // Cek apakah saldo mencukupi
    if senderBalance < amount {
        return fmt.Errorf("saldo tidak mencukupi. Saldo: %.2f, Transfer: %.2f", 
                         senderBalance, amount)
    }
    
    // 2. Kurangi saldo pengirim
    result, err := tx.Exec("UPDATE users SET balance = balance - ? WHERE id = ?", 
                          amount, fromUserID)
    if err != nil {
        return fmt.Errorf("error deducting sender balance: %w", err)
    }
    
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("user pengirim dengan ID %d tidak ditemukan", fromUserID)
    }
    
    // 3. Tambah saldo penerima
    result, err = tx.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", 
                         amount, toUserID)
    if err != nil {
        return fmt.Errorf("error adding to receiver balance: %w", err)
    }
    
    rowsAffected, _ = result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("user penerima dengan ID %d tidak ditemukan", toUserID)
    }
    
    // 4. Catat transaksi ke tabel log
    _, err = tx.Exec(`INSERT INTO transaction_logs (from_user_id, to_user_id, amount, type) 
                     VALUES (?, ?, ?, 'transfer')`, fromUserID, toUserID, amount)
    if err != nil {
        return fmt.Errorf("error logging transaction: %w", err)
    }
    
    // Jika semua berhasil, commit transaksi
    err = tx.Commit()
    if err != nil {
        return fmt.Errorf("error committing transaction: %w", err)
    }
    
    fmt.Printf("Transfer %.2f dari user %d ke user %d berhasil!\n", 
               amount, fromUserID, toUserID)
    return nil
}

// Helper function untuk menjalankan multiple queries dalam satu transaksi
func executeInTransaction(db *sql.DB, fn func(*sql.Tx) error) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()
    
    err = fn(tx)
    return err
}

// Contoh penggunaan helper function
func createUserWithPosts(db *sql.DB, user User, posts []string) error {
    return executeInTransaction(db, func(tx *sql.Tx) error {
        // Insert user
        result, err := tx.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", 
                              user.Name, user.Email, user.Age)
        if err != nil {
            return fmt.Errorf("error creating user: %w", err)
        }
        
        userID, err := result.LastInsertId()
        if err != nil {
            return fmt.Errorf("error getting user ID: %w", err)
        }
        
        // Insert posts untuk user tersebut
        for _, postTitle := range posts {
            _, err = tx.Exec("INSERT INTO posts (title, user_id) VALUES (?, ?)", 
                            postTitle, userID)
            if err != nil {
                return fmt.Errorf("error creating post '%s': %w", postTitle, err)
            }
        }
        
        fmt.Printf("User %s dan %d posts berhasil dibuat!\n", user.Name, len(posts))
        return nil
    })
}

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/tutorial_go")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Test transfer money
    err = transferMoney(db, 1, 2, 100.50)
    if err != nil {
        log.Printf("Transfer gagal: %v", err)
    }
    
    // Test create user with posts
    newUser := User{Name: "David Lee", Email: "david@example.com", Age: 29}
    posts := []string{
        "My First Blog Post",
        "Learning Go and MySQL",
        "Database Best Practices",
    }
    
    err = createUserWithPosts(db, newUser, posts)
    if err != nil {
        log.Printf("Error creating user with posts: %v", err)
    }
}
```

**Penjelasan Transaksi:**
1. `db.Begin()`: Memulai transaksi baru
2. `tx.Commit()`: Menyimpan semua perubahan ke database
3. `tx.Rollback()`: Membatalkan semua perubahan jika terjadi error
4. **Defer Pattern**: Memastikan rollback dipanggil jika terjadi error
5. **ACID Properties** dijamin oleh database engine

---

## Connection Pooling

Go secara otomatis mengelola connection pool, tapi kita bisa mengkonfigurasinya.

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    "sync"
    "context"
    
    _ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
    MaxOpenConns    int
    MaxIdleConns    int
    ConnMaxLifetime time.Duration
    ConnMaxIdleTime time.Duration
}

func setupConnectionPool(dsn string, config DatabaseConfig) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    // Konfigurasi connection pool
    db.SetMaxOpenConns(config.MaxOpenConns)
    db.SetMaxIdleConns(config.MaxIdleConns)
    db.SetConnMaxLifetime(config.ConnMaxLifetime)
    db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
    
    return db, nil
}

func demonstrateConnectionPooling(db *sql.DB) {
    // Simulasi banyak query concurrent
    var wg sync.WaitGroup
    numWorkers := 50
    queriesPerWorker := 10
    
    start := time.Now()
    
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            
            for j := 0; j < queriesPerWorker; j++ {
                // Context dengan timeout untuk setiap query
                ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
                
                var count int
                err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
                if err != nil {
                    log.Printf("Worker %d, Query %d error: %v", workerID, j, err)
                } else {
                    fmt.Printf("Worker %d, Query %d: Found %d users\n", workerID, j, count)
                }
                
                cancel()
                
                // Simulasi delay antar query
                time.Sleep(100 * time.Millisecond)
            }
        }(i)
    }
    
    wg.Wait()
    
    duration := time.Since(start)
    totalQueries := numWorkers * queriesPerWorker
    fmt.Printf("\nSelesai %d queries dalam %v (%.2f queries/second)\n", 
               totalQueries, duration, float64(totalQueries)/duration.Seconds())
    
    // Tampilkan statistik connection pool
    stats := db.Stats()
    fmt.Printf("Connection Pool Stats:\n")
    fmt.Printf("- Max Open Connections: %d\n", stats.MaxOpenConnections)
    fmt.Printf("- Open Connections: %d\n", stats.OpenConnections)
    fmt.Printf("- Idle Connections: %d\n", stats.Idle)
    fmt.Printf("- Connections In Use: %d\n", stats.InUse)
    fmt.Printf("- Wait Count: %d\n", stats.WaitCount)
    fmt.Printf("- Wait Duration: %v\n", stats.WaitDuration)
}

// Monitoring connection pool health
func monitorConnectionPool(db *sql.DB, interval time.Duration) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            stats := db.Stats()
            
            // Log jika ada masalah dengan connection pool
            if stats.WaitCount > 0 {
                log.Printf("WARNING: Connection pool wait count: %d, duration: %v", 
                          stats.WaitCount, stats.WaitDuration)
            }
            
            // Log jika idle connections terlalu banyak
            idleRatio := float64(stats.Idle) / float64(stats.MaxOpenConnections)
            if idleRatio > 0.8 {
                log.Printf("INFO: High idle connection ratio: %.2f%%", idleRatio*100)
            }
            
            fmt.Printf("Pool Stats - Open: %d, Idle: %d, InUse: %d, Wait: %d\n",
                      stats.OpenConnections, stats.Idle, stats.InUse, stats.WaitCount)
        }
    }
}

func main() {
    // Konfigurasi connection pool
    config := DatabaseConfig{
        MaxOpenConns:    25,                // Maksimal 25 koneksi terbuka
        MaxIdleConns:    10,                // Maksimal 10 koneksi idle
        ConnMaxLifetime: 1 * time.Hour,     // Koneksi di-refresh setiap 1 jam
        ConnMaxIdleTime: 10 * time.Minute,  // Koneksi idle ditutup setelah 10 menit
    }
    
    dsn := "root:password@tcp(localhost:3306)/tutorial_go?parseTime=true"
    
    db, err := setupConnectionPool(dsn, config)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Database connection pool siap!")
    
    // Jalankan monitoring di background
    go monitorConnectionPool(db, 5*time.Second)
    
    // Demonstrasi concurrent queries
    demonstrateConnectionPooling(db)
    
    // Tunggu sebentar untuk melihat monitoring
    time.Sleep(30 * time.Second)
}
```

**Penjelasan Connection Pool:**

1. **MaxOpenConns**: Jumlah maksimal koneksi yang bisa dibuka
2. **MaxIdleConns**: Jumlah maksimal koneksi idle yang disimpan
3. **ConnMaxLifetime**: Berapa lama koneksi bisa hidup sebelum di-refresh
4. **ConnMaxIdleTime**: Berapa lama koneksi idle sebelum ditutup

**Tips Konfigurasi:**
- **MaxOpenConns**: Sesuaikan dengan kapasitas database (biasanya 25-100)
- **MaxIdleConns**: Sekitar 25