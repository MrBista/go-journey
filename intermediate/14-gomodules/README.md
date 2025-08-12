# Panduan Lengkap Go Modules untuk Pemula

## Apa itu Go Modules?

Go Modules adalah sistem manajemen dependensi resmi untuk Go yang menggantikan GOPATH. Dengan Go Modules, Anda dapat:
- Mengelola versi dependensi dengan tepat
- Bekerja di luar GOPATH
- Memastikan reproducible builds
- Mengelola dependensi secara otomatis

## Konsep Dasar

### Module
Module adalah kumpulan Go packages yang disimpan dalam satu repository dan di-release bersama. Module didefinisikan oleh file `go.mod` di root directory.

### go.mod File
File ini berisi:
- Module path (biasanya URL repository)
- Go version requirement
- Daftar dependensi dengan versinya

### go.sum File
File ini berisi checksums untuk memverifikasi integritas dependensi.

## Memulai dengan Go Modules

### 1. Inisialisasi Module Baru

```bash
# Membuat direktori project baru
mkdir my-project
cd my-project

# Inisialisasi module
go mod init github.com/username/my-project
```

Ini akan membuat file `go.mod`:
```go
module github.com/username/my-project

go 1.21
```

### 2. Menambahkan Dependensi

#### Cara Manual
Edit `go.mod` dan tambahkan require:
```go
module github.com/username/my-project

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
)
```

#### Cara Otomatis (Recommended)
```bash
# Download dependensi
go get github.com/gin-gonic/gin

# Download versi spesifik
go get github.com/gin-gonic/gin@v1.9.1

# Download versi latest
go get github.com/gin-gonic/gin@latest
```

### 3. Menggunakan Dependensi dalam Code

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })
    r.Run()
}
```

## Perintah Go Modules yang Penting

### go get
```bash
# Download package terbaru
go get github.com/gorilla/mux

# Download versi spesifik
go get github.com/gorilla/mux@v1.8.0

# Upgrade ke versi patch terbaru
go get -u=patch github.com/gorilla/mux

# Upgrade ke versi minor terbaru
go get -u github.com/gorilla/mux

# Download untuk development saja
go get -d github.com/gorilla/mux
```

### go mod tidy
```bash
# Membersihkan dependensi yang tidak digunakan
# dan menambahkan yang missing
go mod tidy
```

### go mod download
```bash
# Download semua dependensi ke module cache
go mod download
```

### go mod verify
```bash
# Verifikasi integritas dependensi
go mod verify
```

### go list
```bash
# List semua dependensi
go list -m all

# List dependensi yang bisa di-upgrade
go list -u -m all
```

## Semantic Versioning

Go Modules menggunakan semantic versioning (semver):
- `v1.2.3` - Major.Minor.Patch
- `v1.2.3-alpha.1` - Pre-release
- `v1.2.3+build.1` - Build metadata

### Version Queries
```bash
go get github.com/gin-gonic/gin@latest    # Versi stabil terbaru
go get github.com/gin-gonic/gin@upgrade   # Upgrade ke versi terbaru
go get github.com/gin-gonic/gin@patch     # Upgrade ke patch terbaru
go get github.com/gin-gonic/gin@v1.9.0    # Versi spesifik
go get github.com/gin-gonic/gin@master    # Branch spesifik
```

## Working dengan Vendor

### Vendor Directory
```bash
# Menyimpan dependensi ke vendor/ directory
go mod vendor

# Build menggunakan vendor
go build -mod=vendor
```

## Contoh go.mod File Lengkap

```go
module github.com/username/my-project

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/gorilla/mux v1.8.0
    gorm.io/gorm v1.25.4
)

require (
    // Indirect dependencies (otomatis ditambahkan)
    github.com/bytedance/sonic v1.9.1 // indirect
    github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
    // ... dependencies lainnya
)

// Replace directive untuk development lokal
replace github.com/my-fork/package => ../local-package

// Exclude directive untuk mengabaikan versi tertentu
exclude github.com/problematic/package v1.0.0
```

## Best Practices

### 1. Gunakan Semantic Import Versioning
Untuk major version v2+, sertakan versi dalam import path:
```go
import "github.com/my/module/v2"
```

### 2. Minimal Version Selection
Go memilih versi minimal yang memenuhi semua requirement. Ini memastikan reproducible builds.

### 3. Always Run go mod tidy
```bash
# Setelah menambah/hapus import
go mod tidy
```

### 4. Commit go.sum File
File `go.sum` harus di-commit ke version control untuk memastikan reproducible builds.

### 5. Gunakan go mod vendor untuk Production
```bash
go mod vendor
```
Kemudian build dengan: `go build -mod=vendor`

### 6. Naming Convention
- Module path menggunakan lowercase dan hyphens
- Hindari underscore dalam module path
- Gunakan domain yang Anda kontrol

### 7. Version Tagging
```bash
# Tag versi baru
git tag v1.0.0
git push origin v1.0.0
```

## Troubleshooting Common Issues

### 1. Module Sum Mismatch
```bash
# Hapus go.sum dan download ulang
rm go.sum
go mod download
```

### 2. Proxy Issues
```bash
# Disable module proxy
export GOPROXY=direct

# Atau set ke proxy lain
export GOPROXY=https://proxy.golang.org,direct
```

### 3. Private Repositories
```bash
# Set GOPRIVATE untuk private repos
export GOPRIVATE=github.com/mycompany/*
```

## Contoh Project Structure

```
my-project/
├── go.mod
├── go.sum
├── main.go
├── internal/
│   ├── handler/
│   │   └── user.go
│   └── service/
│       └── user.go
├── pkg/
│   └── utils/
│       └── helper.go
└── vendor/ (optional)
```

## Environment Variables

```bash
# Module proxy (default: https://proxy.golang.org,direct)
export GOPROXY=https://proxy.golang.org,direct

# Private modules
export GOPRIVATE=github.com/mycompany/*

# Module checksum database
export GOSUMDB=sum.golang.org

# Disable checksum verification
export GOSUMDB=off
```

## Tips untuk Pemula

1. **Selalu mulai dengan `go mod init`** sebelum menulis code
2. **Gunakan `go mod tidy`** secara rutin
3. **Commit `go.sum`** ke repository
4. **Jangan edit `go.mod`** secara manual kecuali perlu
5. **Gunakan semantic versioning** untuk module Anda sendiri
6. **Test dengan berbagai versi Go** jika membuat library public
7. **Dokumentasikan breaking changes** dengan jelas

Dengan mengikuti panduan ini, Anda akan dapat mengelola dependensi Go dengan efektif menggunakan Go Modules!