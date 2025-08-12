# Panduan Package Go untuk Pemula

## Apa itu Package?

Package di Go adalah cara untuk mengorganisir dan mengelompokkan kode. Setiap file Go harus berada dalam sebuah package, dan setiap aplikasi Go dimulai dari package `main`.

## 1. Package Dasar

### Package main
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

- Package `main` adalah entry point dari aplikasi Go
- Harus memiliki fungsi `main()` sebagai titik awal eksekusi

### Package biasa
```go
// File: math/calculator.go
package math

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}
```

## 2. Struktur Direktori Package

```
myproject/
├── main.go
├── go.mod
├── math/
│   ├── calculator.go
│   └── advanced.go
├── utils/
│   ├── string.go
│   └── file.go
└── models/
    └── user.go
```

## 3. Import Package

### Import package standard
```go
import "fmt"
import "os"
import "time"

// Atau dengan cara ini:
import (
    "fmt"
    "os"
    "time"
)
```

### Import package lokal
```go
package main

import (
    "fmt"
    "myproject/math"
    "myproject/utils"
)

func main() {
    result := math.Add(10, 5)
    fmt.Println(result)
}
```

### Import dengan alias
```go
import (
    f "fmt"           // alias
    m "myproject/math"
    _ "database/sql"  // blank import (hanya untuk side effects)
    . "strings"       // dot import (tidak direkomendasikan)
)
```

## 4. Visibility (Public vs Private)

Di Go, visibility ditentukan oleh huruf pertama nama:

### Public (Exported)
```go
package math

// Public - bisa diakses dari package lain
func Add(a, b int) int {
    return a + b
}

// Public variable
var Pi = 3.14159

// Public struct
type Calculator struct {
    Name string
}

// Public method
func (c *Calculator) GetName() string {
    return c.Name
}
```

### Private (Unexported)
```go
package math

// Private - hanya bisa diakses dalam package yang sama
func multiply(a, b int) int {
    return a * b
}

// Private variable
var version = "1.0"

// Private struct
type config struct {
    debug bool
}
```

## 5. Init Function

Fungsi `init()` dijalankan otomatis saat package di-import:

```go
package database

import "fmt"

var connection string

func init() {
    fmt.Println("Initializing database connection...")
    connection = "localhost:5432"
}

func GetConnection() string {
    return connection
}
```

## 6. Package Documentation

Gunakan komentar untuk dokumentasi:

```go
// Package math provides basic mathematical operations.
// This package includes functions for addition, subtraction, and more.
package math

// Add returns the sum of two integers.
// Example:
//   result := Add(5, 3) // returns 8
func Add(a, b int) int {
    return a + b
}
```

## 7. Best Practices

### 1. Naming Conventions
```go
// ✅ Good - singkat dan deskriptif
package user
package http
package json

// ❌ Avoid - terlalu panjang atau tidak jelas
package userManagement
package utilities
package stuff
```

### 2. Package Organization
```go
// ✅ Kelompokkan berdasarkan fungsionalitas
project/
├── user/
│   ├── user.go      // User model
│   ├── service.go   // User business logic
│   └── handler.go   // User HTTP handlers
├── order/
│   ├── order.go
│   ├── service.go
│   └── handler.go
```

### 3. Avoid Circular Dependencies
```go
// ❌ Bad - circular dependency
// Package A imports B, B imports A

// ✅ Good - introduce a third package
project/
├── user/
├── order/
└── common/     // shared types/interfaces
```

### 4. Interface Design
```go
// ✅ Definisikan interface di package yang menggunakannya
package main

import "myproject/database"

type UserRepository interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

func main() {
    var repo UserRepository = database.NewUserRepo()
    // ...
}
```

### 5. Error Handling
```go
package user

import (
    "errors"
    "fmt"
)

var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidEmail = errors.New("invalid email format")
)

func GetUser(id int) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user id: %d", id)
    }
    
    // ... logic here
    
    return nil, ErrUserNotFound
}
```

## 8. Go Modules

### Inisialisasi module
```bash
go mod init myproject
```

### go.mod file
```go
module myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/lib/pq v1.10.9
)
```

### Import external packages
```go
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "myproject/internal/user"
)
```

## 9. Contoh Struktur Project Lengkap

```
ecommerce/
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   ├── server/
│   └── cli/
├── internal/           // private packages
│   ├── user/
│   │   ├── model.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── handler.go
│   ├── order/
│   └── product/
├── pkg/               // public packages
│   ├── logger/
│   ├── config/
│   └── middleware/
├── api/
│   └── routes.go
├── migrations/
├── docs/
└── tests/
```

## 10. Testing Packages

```go
// File: math/calculator_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

## Tips Tambahan

1. **Package names should be lowercase**: `package user`, bukan `package User`
2. **Avoid stuttering**: Jangan `user.UserService`, gunakan `user.Service`
3. **Keep packages focused**: Satu package satu tanggung jawab
4. **Use internal/ for private code**: Code di `internal/` tidak bisa di-import dari luar module
5. **Document your public API**: Semua exported functions/types harus punya dokumentasi

## Latihan

1. Buat package `calculator` dengan operasi matematika dasar
2. Buat package `validator` untuk validasi email dan password
3. Buat aplikasi sederhana yang menggunakan kedua package tersebut
4. Implementasikan error handling yang proper
5. Tulis unit test untuk semua fungsi public

Mulailah dengan contoh sederhana dan secara bertahap tingkatkan kompleksitasnya. Package adalah fondasi penting dalam Go, jadi luangkan waktu untuk memahami konsep ini dengan baik!