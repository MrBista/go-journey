## 0. Pendahuluan Konsep Penting yang Harus Dipahami:

- Error sebagai Value: Di Go, error bukan exception tapi nilai biasa yang harus di-check secara eksplisit
- Pattern if err != nil: Ini adalah pattern paling umum yang akan kamu lihat di mana-mana dalam kode Go
- Error Wrapping: Menggunakan %w dalam fmt.Errorf() untuk menjaga jejak error asli sambil menambah context
- Early Return: Segera return ketika ada error, jangan lanjutkan eksekusi


# Panduan Lengkap Error Handling di Go

## 1. Konsep Dasar Error di Go

Di Go, error adalah **value** biasa, bukan exception seperti di bahasa lain. Error di Go mengimplementasikan interface `error` yang sederhana:

```go
type error interface {
    Error() string
}
```

### Mengapa Go Menggunakan Error sebagai Value?

- Explicit error handling - kamu harus secara sadar menangani error
- Lebih mudah dipahami alur program
- Tidak ada "hidden control flow" seperti try-catch

## 2. Cara Membuat Error

### a. Menggunakan `errors.New()`

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("tidak bisa membagi dengan nol")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Hasil:", result)
}
```

### b. Menggunakan `fmt.Errorf()` (Recommended)

```go
func validateAge(age int) error {
    if age < 0 {
        return fmt.Errorf("umur tidak valid: %d (harus >= 0)", age)
    }
    if age > 150 {
        return fmt.Errorf("umur tidak realistis: %d", age)
    }
    return nil
}
```

### c. Custom Error Type

```go
type ValidationError struct {
    Field string
    Value interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validasi gagal untuk field '%s' dengan value '%v': %s", 
        e.Field, e.Value, e.Message)
}

func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return ValidationError{
            Field: "email",
            Value: email,
            Message: "format email tidak valid",
        }
    }
    return nil
}
```

## 3. Best Practices Error Handling

### a. Selalu Check Error

```go
// ❌ SALAH - mengabaikan error
result, _ := someFunction()

// ✅ BENAR - selalu check error
result, err := someFunction()
if err != nil {
    // handle error
    return err
}
```

### b. Early Return Pattern

```go
func processData(data string) error {
    // Validate input
    if err := validateInput(data); err != nil {
        return fmt.Errorf("validasi input gagal: %w", err)
    }
    
    // Parse data
    parsed, err := parseData(data)
    if err != nil {
        return fmt.Errorf("parsing data gagal: %w", err)
    }
    
    // Save to database
    if err := saveToDatabase(parsed); err != nil {
        return fmt.Errorf("simpan ke database gagal: %w", err)
    }
    
    return nil
}
```

### c. Error Wrapping dengan `%w`

```go
func readConfigFile(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("gagal membaca file config %s: %w", filename, err)
    }
    
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, fmt.Errorf("gagal parsing config JSON: %w", err)
    }
    
    return &config, nil
}
```

### d. Error Unwrapping

```go
import "errors"

func handleConfigError(err error) {
    // Check apakah error ini adalah file not found
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File config tidak ditemukan")
        return
    }
    
    // Check tipe error custom
    var validationErr ValidationError
    if errors.As(err, &validationErr) {
        fmt.Printf("Error validasi: %s\n", validationErr.Message)
        return
    }
    
    fmt.Printf("Error lain: %v\n", err)
}
```

## 4. Pattern Error Handling yang Umum

### a. Guard Clause Pattern

```go
func processUser(user *User) error {
    if user == nil {
        return errors.New("user tidak boleh nil")
    }
    
    if user.Email == "" {
        return errors.New("email wajib diisi")
    }
    
    if user.Age < 18 {
        return errors.New("umur harus minimal 18 tahun")
    }
    
    // Logic utama di sini
    return nil
}
```

### b. Error Aggregation

```go
func validateUser(user *User) error {
    var errs []string
    
    if user.Name == "" {
        errs = append(errs, "nama wajib diisi")
    }
    
    if user.Email == "" {
        errs = append(errs, "email wajib diisi")
    }
    
    if user.Age < 0 {
        errs = append(errs, "umur tidak valid")
    }
    
    if len(errs) > 0 {
        return fmt.Errorf("validasi gagal: %s", strings.Join(errs, ", "))
    }
    
    return nil
}
```

### c. Retry Pattern dengan Error

```go
func withRetry(operation func() error, maxRetries int) error {
    var lastErr error
    
    for i := 0; i < maxRetries; i++ {
        if err := operation(); err != nil {
            lastErr = err
            fmt.Printf("Attempt %d failed: %v\n", i+1, err)
            time.Sleep(time.Second * time.Duration(i+1)) // exponential backoff
            continue
        }
        return nil // sukses
    }
    
    return fmt.Errorf("operasi gagal setelah %d percobaan, error terakhir: %w", maxRetries, lastErr)
}

// Penggunaan:
err := withRetry(func() error {
    return callExternalAPI()
}, 3)
```

## 5. Logging Error dengan Proper Context

```go
import "log/slog"

func processOrder(orderID string) error {
    logger := slog.With("order_id", orderID)
    
    order, err := getOrder(orderID)
    if err != nil {
        logger.Error("gagal mengambil order", "error", err)
        return fmt.Errorf("gagal mengambil order %s: %w", orderID, err)
    }
    
    if err := validateOrder(order); err != nil {
        logger.Warn("validasi order gagal", "error", err)
        return fmt.Errorf("validasi order gagal: %w", err)
    }
    
    logger.Info("order berhasil diproses")
    return nil
}
```

## 6. Testing Error Handling

```go
func TestDivideByZero(t *testing.T) {
    _, err := divide(10, 0)
    if err == nil {
        t.Error("Expected error when dividing by zero")
    }
    
    expectedMsg := "tidak bisa membagi dengan nol"
    if err.Error() != expectedMsg {
        t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
    }
}

func TestErrorWrapping(t *testing.T) {
    err := processData("invalid data")
    if err == nil {
        t.Fatal("Expected error")
    }
    
    // Test error wrapping
    if !errors.Is(err, SomeSpecificError) {
        t.Error("Error should wrap SomeSpecificError")
    }
}
```

## 7. Anti-Patterns yang Harus Dihindari

### ❌ Mengabaikan Error

```go
// JANGAN seperti ini
result, _ := riskyOperation()
```

### ❌ Panic untuk Error Biasa

```go
// JANGAN seperti ini
if err != nil {
    panic(err) // hanya untuk programmer error
}
```

### ❌ Error Message yang Tidak Informatif

```go
// JANGAN seperti ini
return errors.New("error")

// Lebih baik:
return fmt.Errorf("gagal memvalidasi email %s: format tidak valid", email)
```

## 8. Error Handling di HTTP Handler

```go
func userHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")
    if userID == "" {
        http.Error(w, "user ID wajib diisi", http.StatusBadRequest)
        return
    }
    
    user, err := getUserFromDB(userID)
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            http.Error(w, "user tidak ditemukan", http.StatusNotFound)
            return
        }
        // Log internal error
        log.Printf("Internal error: %v", err)
        http.Error(w, "internal server error", http.StatusInternalServerError)
        return
    }
    
    // Return user data
    json.NewEncoder(w).Encode(user)
}
```

## 9. Error Sentinel Values

```go
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrUnauthorized = errors.New("unauthorized")
)

// Penggunaan:
func getUser(id string) (*User, error) {
    // ... logic ...
    if userNotExists {
        return nil, ErrUserNotFound
    }
    return user, nil
}

// Check error:
user, err := getUser("123")
if errors.Is(err, ErrUserNotFound) {
    // handle user not found
}
```

## Key Takeaways

1. **Selalu handle error** - jangan pernah abaikan
2. **Gunakan error wrapping** dengan `%w` untuk context
3. **Buat error message yang informatif**
4. **Gunakan early return pattern**
5. **Test error handling** sama pentingnya dengan happy path
6. **Log error dengan context yang cukup**
7. **Gunakan custom error types** untuk error yang kompleks
8. **Konsisten dalam error handling** di seluruh codebase