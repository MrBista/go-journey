# Panduan Lengkap Logger di Go untuk Pemula

### PZN MATERI: https://docs.google.com/presentation/d/1edVvzU_sOExvlN4lzYWOxF38v5GAJHdbbYUqsderNO0/edit?slide=id.p#slide=id.p

## 1. Apa itu Logger?

Logger adalah komponen yang digunakan untuk mencatat pesan atau informasi tentang apa yang terjadi dalam aplikasi Anda. Bayangkan logger seperti buku catatan harian aplikasi yang mencatat semua aktivitas penting, error, warning, dan informasi debug.

### Mengapa Logger Penting?

1. **Debugging**: Membantu menemukan dan memperbaiki bug
2. **Monitoring**: Memantau kesehatan aplikasi
3. **Audit**: Melacak aktivitas pengguna
4. **Performance Analysis**: Menganalisis performa aplikasi
5. **Security**: Mendeteksi aktivitas mencurigakan

## 2. Package Log Bawaan Go

Go menyediakan package `log` bawaan yang sederhana namun fungsional.

### Basic Logging

```go
package main

import (
    "log"
)

func main() {
    // Log pesan sederhana
    log.Println("Aplikasi dimulai")
    
    // Log dengan format
    name := "John"
    age := 25
    log.Printf("User %s berusia %d tahun", name, age)
    
    // Log yang akan menghentikan program
    // log.Fatal("Error fatal! Program akan dihentikan")
    
    // Log yang akan menghentikan program dengan panic
    // log.Panic("Panic! Ada masalah serius")
}
```

**Penjelasan:**
- `log.Println()`: Mencetak pesan dengan newline
- `log.Printf()`: Mencetak pesan dengan format (seperti `fmt.Printf`)
- `log.Fatal()`: Mencetak pesan lalu menghentikan program dengan `os.Exit(1)`
- `log.Panic()`: Mencetak pesan lalu memanggil `panic()`

### Kustomisasi Log Output

```go
package main

import (
    "log"
    "os"
)

func main() {
    // Set output ke file
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Gagal membuka file log:", err)
    }
    defer file.Close()
    
    log.SetOutput(file)
    
    // Set format log
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    
    log.Println("Log ini akan masuk ke file app.log")
}
```

**Flag Options:**
- `log.Ldate`: Menampilkan tanggal (2009/01/23)
- `log.Ltime`: Menampilkan waktu (01:23:23)
- `log.Lmicroseconds`: Menampilkan mikrodetik
- `log.Llongfile`: Menampilkan full path file
- `log.Lshortfile`: Menampilkan nama file saja
- `log.LUTC`: Menggunakan UTC time
- `log.LstdFlags`: Kombinasi Ldate dan Ltime

## 3. Logger dengan Level

Package log bawaan tidak memiliki level logging. Mari buat implementasi sederhana:

```go
package main

import (
    "fmt"
    "log"
    "os"
)

type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARNING
    ERROR
    FATAL
)

type Logger struct {
    debugLog   *log.Logger
    infoLog    *log.Logger
    warningLog *log.Logger
    errorLog   *log.Logger
    fatalLog   *log.Logger
    level      LogLevel
}

func NewLogger(level LogLevel) *Logger {
    return &Logger{
        debugLog:   log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
        infoLog:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
        warningLog: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
        errorLog:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
        fatalLog:   log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
        level:      level,
    }
}

func (l *Logger) Debug(v ...interface{}) {
    if l.level <= DEBUG {
        l.debugLog.Println(v...)
    }
}

func (l *Logger) Info(v ...interface{}) {
    if l.level <= INFO {
        l.infoLog.Println(v...)
    }
}

func (l *Logger) Warning(v ...interface{}) {
    if l.level <= WARNING {
        l.warningLog.Println(v...)
    }
}

func (l *Logger) Error(v ...interface{}) {
    if l.level <= ERROR {
        l.errorLog.Println(v...)
    }
}

func (l *Logger) Fatal(v ...interface{}) {
    if l.level <= FATAL {
        l.fatalLog.Println(v...)
        os.Exit(1)
    }
}

func main() {
    logger := NewLogger(INFO) // Hanya log INFO ke atas yang akan ditampilkan
    
    logger.Debug("Ini adalah pesan debug") // Tidak akan tampil
    logger.Info("Aplikasi dimulai")        // Akan tampil
    logger.Warning("Ini adalah warning")   // Akan tampil
    logger.Error("Terjadi error")         // Akan tampil
}
```

**Penjelasan Konsep Level:**
- **DEBUG**: Informasi detail untuk debugging (level terendah)
- **INFO**: Informasi umum tentang jalannya aplikasi
- **WARNING**: Peringatan tentang situasi yang tidak normal tapi tidak fatal
- **ERROR**: Error yang terjadi tapi aplikasi masih bisa berjalan
- **FATAL**: Error serius yang menyebabkan aplikasi harus dihentikan

## 4. Menggunakan Library Logger Populer

### A. Logrus - Logger yang Powerful

```go
package main

import (
    "github.com/sirupsen/logrus"
)

func main() {
    // Inisialisasi logrus
    logger := logrus.New()
    
    // Set level
    logger.SetLevel(logrus.InfoLevel)
    
    // Set format JSON
    logger.SetFormatter(&logrus.JSONFormatter{})
    
    // Contoh penggunaan
    logger.WithFields(logrus.Fields{
        "user_id": 12345,
        "action":  "login",
    }).Info("User berhasil login")
    
    logger.WithFields(logrus.Fields{
        "error": "database connection failed",
        "retry": 3,
    }).Error("Gagal koneksi ke database")
}
```

### B. Zap - High Performance Logger

```go
package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {
    // Development config (untuk debugging)
    devLogger, _ := zap.NewDevelopment()
    defer devLogger.Sync()
    
    devLogger.Info("Ini adalah development logger",
        zap.String("user", "john"),
        zap.Int("age", 25),
    )
    
    // Production config (untuk production)
    prodLogger, _ := zap.NewProduction()
    defer prodLogger.Sync()
    
    prodLogger.Info("Ini adalah production logger",
        zap.String("service", "user-api"),
        zap.String("version", "1.0.0"),
    )
}
```

## 5. Best Practices untuk Logging

### A. Structured Logging

Gunakan structured logging untuk memudahkan parsing dan analisis:

```go
package main

import (
    "context"
    "github.com/sirupsen/logrus"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func loginUser(ctx context.Context, user User, logger *logrus.Logger) error {
    // Log dengan context dan structured fields
    logger.WithFields(logrus.Fields{
        "user_id":    user.ID,
        "user_name":  user.Name,
        "operation":  "login",
        "request_id": ctx.Value("request_id"),
    }).Info("Attempting user login")
    
    // Simulasi proses login
    success := true // Anggap berhasil
    
    if success {
        logger.WithFields(logrus.Fields{
            "user_id":   user.ID,
            "operation": "login",
            "result":    "success",
        }).Info("User login successful")
        return nil
    } else {
        logger.WithFields(logrus.Fields{
            "user_id":   user.ID,
            "operation": "login",
            "result":    "failed",
            "reason":    "invalid_credentials",
        }).Error("User login failed")
        return fmt.Errorf("login failed")
    }
}
```

### B. Logger Middleware untuk HTTP Server

```go
package main

import (
    "net/http"
    "time"
    
    "github.com/sirupsen/logrus"
)

func LoggerMiddleware(logger *logrus.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            
            // Buat wrapper untuk response writer agar bisa capture status code
            wrapped := &responseWriter{
                ResponseWriter: w,
                statusCode:     http.StatusOK,
            }
            
            // Jalankan handler berikutnya
            next.ServeHTTP(wrapped, r)
            
            // Log request details
            logger.WithFields(logrus.Fields{
                "method":      r.Method,
                "url":         r.URL.String(),
                "status_code": wrapped.statusCode,
                "duration":    time.Since(start),
                "user_agent":  r.UserAgent(),
                "remote_addr": r.RemoteAddr,
            }).Info("HTTP Request")
        })
    }
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func main() {
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})
    
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
    
    // Gunakan logger middleware
    handler := LoggerMiddleware(logger)(mux)
    
    http.ListenAndServe(":8080", handler)
}
```

## 6. Configuration Logger

### Logger Configuration dengan File

```go
package main

import (
    "encoding/json"
    "os"
    
    "github.com/sirupsen/logrus"
)

type LogConfig struct {
    Level      string `json:"level"`
    Format     string `json:"format"`
    OutputFile string `json:"output_file"`
}

func setupLogger(configFile string) (*logrus.Logger, error) {
    logger := logrus.New()
    
    // Baca konfigurasi dari file
    file, err := os.Open(configFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var config LogConfig
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }
    
    // Set level
    level, err := logrus.ParseLevel(config.Level)
    if err != nil {
        return nil, err
    }
    logger.SetLevel(level)
    
    // Set format
    if config.Format == "json" {
        logger.SetFormatter(&logrus.JSONFormatter{})
    } else {
        logger.SetFormatter(&logrus.TextFormatter{
            FullTimestamp: true,
        })
    }
    
    // Set output file
    if config.OutputFile != "" {
        file, err := os.OpenFile(config.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
        if err != nil {
            return nil, err
        }
        logger.SetOutput(file)
    }
    
    return logger, nil
}

func main() {
    // File config.json:
    // {
    //   "level": "info",
    //   "format": "json",
    //   "output_file": "app.log"
    // }
    
    logger, err := setupLogger("config.json")
    if err != nil {
        log.Fatal("Error setting up logger:", err)
    }
    
    logger.Info("Logger berhasil dikonfigurasi")
}
```

## 7. Use Cases Logging dalam Aplikasi Nyata

### A. Database Operations

```go
func getUserByID(db *sql.DB, userID int, logger *logrus.Logger) (*User, error) {
    logger.WithField("user_id", userID).Debug("Fetching user from database")
    
    start := time.Now()
    
    query := "SELECT id, name, email FROM users WHERE id = ?"
    row := db.QueryRow(query, userID)
    
    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    
    duration := time.Since(start)
    
    if err != nil {
        if err == sql.ErrNoRows {
            logger.WithFields(logrus.Fields{
                "user_id":  userID,
                "duration": duration,
            }).Warn("User not found")
            return nil, nil
        }
        
        logger.WithFields(logrus.Fields{
            "user_id":  userID,
            "duration": duration,
            "error":    err.Error(),
        }).Error("Database query failed")
        return nil, err
    }
    
    logger.WithFields(logrus.Fields{
        "user_id":  userID,
        "duration": duration,
    }).Debug("User fetched successfully")
    
    return &user, nil
}
```

### B. External API Calls

```go
func callExternalAPI(url string, logger *logrus.Logger) (*http.Response, error) {
    logger.WithField("url", url).Debug("Making external API call")
    
    start := time.Now()
    client := &http.Client{Timeout: 30 * time.Second}
    
    resp, err := client.Get(url)
    duration := time.Since(start)
    
    if err != nil {
        logger.WithFields(logrus.Fields{
            "url":      url,
            "duration": duration,
            "error":    err.Error(),
        }).Error("External API call failed")
        return nil, err
    }
    
    logger.WithFields(logrus.Fields{
        "url":         url,
        "status_code": resp.StatusCode,
        "duration":    duration,
    }).Info("External API call completed")
    
    return resp, nil
}
```

## 8. Tips untuk Logging yang Efektif

### A. Jangan Berlebihan dalam Logging
```go
// BURUK - Terlalu verbose
logger.Debug("Memulai fungsi processUser")
logger.Debug("Validasi input")
logger.Debug("Input valid")
logger.Debug("Memulai proses")
logger.Debug("Proses selesai")

// BAIK - Hanya log yang penting
logger.WithField("user_id", userID).Info("Processing user started")
// ... proses ...
logger.WithField("user_id", userID).Info("User processed successfully")
```

### B. Gunakan Context untuk Request Tracking
```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    requestID := generateRequestID()
    ctx := context.WithValue(r.Context(), "request_id", requestID)
    
    logger := logrus.WithField("request_id", requestID)
    logger.Info("Request started")
    
    // Pass context dan logger ke function lain
    result, err := processRequest(ctx, logger)
    if err != nil {
        logger.WithError(err).Error("Request processing failed")
        http.Error(w, "Internal Server Error", 500)
        return
    }
    
    logger.Info("Request completed successfully")
    json.NewEncoder(w).Encode(result)
}
```

### C. Log Errors dengan Stack Trace
```go
import "github.com/pkg/errors"

func processData() error {
    err := someFunction()
    if err != nil {
        // Wrap error dengan stack trace
        return errors.Wrap(err, "failed to process data")
    }
    return nil
}

func main() {
    err := processData()
    if err != nil {
        logger.WithFields(logrus.Fields{
            "error": err.Error(),
            "stack": fmt.Sprintf("%+v", err), // Stack trace
        }).Error("Processing failed")
    }
}
```

## 9. Kesimpulan

Logger adalah komponen penting dalam aplikasi Go yang membantu dalam:
- **Development**: Debugging dan troubleshooting
- **Operations**: Monitoring dan alerting  
- **Business**: Analytics dan audit trails

**Key Takeaways:**
1. Mulai dengan package `log` bawaan untuk aplikasi sederhana
2. Gunakan library seperti Logrus atau Zap untuk aplikasi yang lebih kompleks
3. Selalu gunakan structured logging dengan fields yang konsisten
4. Implementasikan level logging yang sesuai
5. Jangan lupa konfigurasi logger untuk berbagai environment (dev, staging, prod)
6. Gunakan context untuk request tracking
7. Log error dengan informasi yang cukup untuk debugging

Dengan mengikuti panduan ini, Anda akan dapat mengimplementasikan logging yang efektif dalam aplikasi Go Anda!