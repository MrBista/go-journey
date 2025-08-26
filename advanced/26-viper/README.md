# Panduan Lengkap Viper untuk Pemula Go


### PZN MATERI: https://docs.google.com/presentation/d/10Hff_VEA7NwjI3Qow8vIDE1UqyuFb2JRChAFczkoMdg/edit?slide=id.p#slide=id.p

## 1. Pengenalan Viper

Viper adalah salah satu library Go yang paling populer untuk mengelola konfigurasi aplikasi. Library ini dikembangkan oleh Steve Francia (spf13) dan digunakan oleh banyak project besar seperti Hugo dan Kubernetes.

### Mengapa Viper?
- **Fleksibel**: Mendukung berbagai format file konfigurasi (JSON, TOML, YAML, HCL, etc.)
- **Multi-source**: Bisa membaca dari file, environment variables, command flags, remote config
- **Type-safe**: Konversi otomatis ke berbagai tipe data Go
- **Live reloading**: Bisa reload konfigurasi secara otomatis tanpa restart aplikasi
- **Hierarchical**: Mendukung nested configuration
- **Default values**: Bisa set nilai default untuk setiap konfigurasi

## 2. Instalasi dan Setup

### Instalasi Viper
```bash
go mod init viper-tutorial
go get github.com/spf13/viper
```

### Import yang Diperlukan
```go
import (
    "fmt"
    "log"
    "github.com/spf13/viper"
)
```

## 3. Contoh Dasar Penggunaan Viper

### 3.1 Membaca dari File Konfigurasi

Pertama, kita buat file konfigurasi `config.yaml`:

```yaml
# config.yaml
app:
  name: "MyApp"
  version: "1.0.0"
  debug: true
  port: 8080

database:
  host: "localhost"
  port: 5432
  username: "admin"
  password: "secret"
  name: "mydb"

redis:
  host: "localhost"
  port: 6379
  password: ""

logging:
  level: "info"
  file: "app.log"
```

Sekarang kita buat kode Go untuk membaca konfigurasi ini:

```go
package main

import (
    "fmt"
    "log"
    "github.com/spf13/viper"
)

func main() {
    // Inisialisasi Viper
    initConfig()
    
    // Membaca konfigurasi
    readConfiguration()
}

func initConfig() {
    // Set nama file konfigurasi (tanpa ekstensi)
    viper.SetConfigName("config")
    
    // Set tipe file konfigurasi
    viper.SetConfigType("yaml")
    
    // Tambahkan path dimana file konfigurasi berada
    viper.AddConfigPath(".")      // direktori saat ini
    viper.AddConfigPath("./config") // direktori config
    viper.AddConfigPath("/etc/myapp/") // direktori sistem
    
    // Baca file konfigurasi
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    
    fmt.Println("Config file used:", viper.ConfigFileUsed())
}

func readConfiguration() {
    // Cara 1: Membaca dengan Get (return interface{})
    appName := viper.Get("app.name")
    fmt.Printf("App Name: %v (type: %T)\n", appName, appName)
    
    // Cara 2: Membaca dengan GetString (return string)
    appNameStr := viper.GetString("app.name")
    fmt.Printf("App Name (string): %s\n", appNameStr)
    
    // Membaca berbagai tipe data
    appVersion := viper.GetString("app.version")
    debugMode := viper.GetBool("app.debug")
    appPort := viper.GetInt("app.port")
    
    fmt.Printf("Version: %s\n", appVersion)
    fmt.Printf("Debug Mode: %t\n", debugMode)
    fmt.Printf("Port: %d\n", appPort)
    
    // Membaca nested configuration
    dbHost := viper.GetString("database.host")
    dbPort := viper.GetInt("database.port")
    dbUser := viper.GetString("database.username")
    
    fmt.Printf("Database: %s:%d (user: %s)\n", dbHost, dbPort, dbUser)
    
    // Membaca dengan default value
    redisPassword := viper.GetString("redis.password")
    if redisPassword == "" {
        fmt.Println("Redis password is empty")
    }
    
    // Menggunakan GetStringWithFallback (jika tidak ada, gunakan default)
    logLevel := viper.GetString("logging.level")
    if logLevel == "" {
        logLevel = "warn" // default value
    }
    fmt.Printf("Log Level: %s\n", logLevel)
}
```

**Penjelasan Kode:**

1. **`viper.SetConfigName()`**: Menentukan nama file konfigurasi tanpa ekstensi
2. **`viper.SetConfigType()`**: Menentukan format file (yaml, json, toml, dll)
3. **`viper.AddConfigPath()`**: Menambahkan direktori pencarian file konfigurasi
4. **`viper.ReadInConfig()`**: Membaca file konfigurasi
5. **`viper.Get*()`**: Membaca nilai konfigurasi dengan konversi tipe otomatis

## 4. Bekerja dengan Environment Variables

Viper dapat membaca dari environment variables secara otomatis:

```go
package main

import (
    "fmt"
    "os"
    "github.com/spf13/viper"
)

func main() {
    // Set environment variables untuk testing
    os.Setenv("APP_NAME", "MyApp from ENV")
    os.Setenv("APP_PORT", "9000")
    os.Setenv("DATABASE_HOST", "prod-db.example.com")
    
    setupViperWithEnv()
    readFromEnv()
}

func setupViperWithEnv() {
    // Aktifkan pembacaan dari environment variables
    viper.AutomaticEnv()
    
    // Set prefix untuk environment variables
    viper.SetEnvPrefix("APP") // akan membaca APP_* env vars
    
    // Map key konfigurasi ke environment variable
    viper.BindEnv("app.name", "APP_NAME")
    viper.BindEnv("app.port", "APP_PORT")
    viper.BindEnv("database.host", "DATABASE_HOST")
    
    // Alternatif: Set default values
    viper.SetDefault("app.name", "DefaultApp")
    viper.SetDefault("app.port", 8080)
    viper.SetDefault("database.host", "localhost")
}

func readFromEnv() {
    fmt.Println("=== Reading from Environment Variables ===")
    
    // Akan membaca dari env var jika ada, jika tidak akan gunakan default
    appName := viper.GetString("app.name")
    appPort := viper.GetInt("app.port")
    dbHost := viper.GetString("database.host")
    
    fmt.Printf("App Name: %s\n", appName)
    fmt.Printf("App Port: %d\n", appPort)
    fmt.Printf("DB Host: %s\n", dbHost)
    
    // Menampilkan semua environment variables yang dibaca Viper
    fmt.Println("\nAll environment variables with APP prefix:")
    for _, env := range os.Environ() {
        if len(env) >= 4 && env[:4] == "APP_" {
            fmt.Println(env)
        }
    }
}
```

**Key Concepts Environment Variables:**

1. **`viper.AutomaticEnv()`**: Aktifkan pembacaan otomatis dari env vars
2. **`viper.SetEnvPrefix()`**: Set prefix untuk env vars (misal: APP_NAME untuk key "name")
3. **`viper.BindEnv()`**: Map specific key ke environment variable
4. **`viper.SetDefault()`**: Set nilai default jika tidak ada di env var atau file

## 5. Struktur Konfigurasi dengan Struct

Untuk aplikasi yang lebih complex, kita bisa map konfigurasi ke struct:

```go
package main

import (
    "fmt"
    "log"
    "github.com/spf13/viper"
)

// Definisi struct untuk konfigurasi
type Config struct {
    App      AppConfig      `mapstructure:"app"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    Logging  LoggingConfig  `mapstructure:"logging"`
}

type AppConfig struct {
    Name    string `mapstructure:"name"`
    Version string `mapstructure:"version"`
    Debug   bool   `mapstructure:"debug"`
    Port    int    `mapstructure:"port"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    Name     string `mapstructure:"name"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
}

type LoggingConfig struct {
    Level string `mapstructure:"level"`
    File  string `mapstructure:"file"`
}

func main() {
    config, err := loadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    printConfig(config)
    
    // Menggunakan konfigurasi
    useConfig(config)
}

func loadConfig() (*Config, error) {
    // Setup viper
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    
    // Set default values
    setDefaults()
    
    // Baca dari environment variables
    viper.AutomaticEnv()
    
    // Baca file konfigurasi
    if err := viper.ReadInConfig(); err != nil {
        // Jika file tidak ditemukan, tidak masalah, kita pakai default + env vars
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("failed to read config file: %w", err)
        }
        fmt.Println("No config file found, using defaults and environment variables")
    } else {
        fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
    }
    
    // Unmarshal ke struct
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %w", err)
    }
    
    return &config, nil
}

func setDefaults() {
    // App defaults
    viper.SetDefault("app.name", "MyApp")
    viper.SetDefault("app.version", "1.0.0")
    viper.SetDefault("app.debug", false)
    viper.SetDefault("app.port", 8080)
    
    // Database defaults
    viper.SetDefault("database.host", "localhost")
    viper.SetDefault("database.port", 5432)
    viper.SetDefault("database.username", "postgres")
    viper.SetDefault("database.password", "")
    viper.SetDefault("database.name", "mydb")
    
    // Redis defaults
    viper.SetDefault("redis.host", "localhost")
    viper.SetDefault("redis.port", 6379)
    viper.SetDefault("redis.password", "")
    
    // Logging defaults
    viper.SetDefault("logging.level", "info")
    viper.SetDefault("logging.file", "app.log")
}

func printConfig(config *Config) {
    fmt.Println("=== Loaded Configuration ===")
    fmt.Printf("App: %s v%s (Debug: %t, Port: %d)\n", 
        config.App.Name, config.App.Version, config.App.Debug, config.App.Port)
    
    fmt.Printf("Database: %s:%d (User: %s, DB: %s)\n", 
        config.Database.Host, config.Database.Port, 
        config.Database.Username, config.Database.Name)
    
    fmt.Printf("Redis: %s:%d\n", config.Redis.Host, config.Redis.Port)
    
    fmt.Printf("Logging: Level=%s, File=%s\n", 
        config.Logging.Level, config.Logging.File)
}

func useConfig(config *Config) {
    fmt.Println("\n=== Using Configuration ===")
    
    // Simulasi penggunaan konfigurasi
    if config.App.Debug {
        fmt.Println("Debug mode is enabled")
    }
    
    // Database connection string
    dbConnStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s", 
        config.Database.Host, config.Database.Port, 
        config.Database.Username, config.Database.Name)
    fmt.Printf("Database connection: %s\n", dbConnStr)
    
    // Redis connection
    redisAddr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
    fmt.Printf("Redis address: %s\n", redisAddr)
    
    // Server address
    serverAddr := fmt.Sprintf(":%d", config.App.Port)
    fmt.Printf("Server will listen on: %s\n", serverAddr)
}
```

**Key Points Struct Mapping:**

1. **`mapstructure` tag**: Digunakan untuk mapping field struct ke key konfigurasi
2. **`viper.Unmarshal()`**: Convert konfigurasi Viper ke struct Go
3. **Nested struct**: Mendukung konfigurasi bertingkat
4. **Type safety**: Struct memberikan type safety dan IDE completion

## 6. Advanced Features

### 6.1 Watch Configuration File (Live Reload)

```go
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/spf13/viper"
    "github.com/fsnotify/fsnotify"
)

func main() {
    // Setup viper
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    
    // Baca initial config
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config: %v", err)
    }
    
    // Setup watch untuk live reload
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Printf("Config file changed: %s\n", e.Name)
        printCurrentConfig()
    })
    
    fmt.Println("Watching for config changes...")
    printCurrentConfig()
    
    // Keep program running
    for {
        time.Sleep(1 * time.Second)
    }
}

func printCurrentConfig() {
    fmt.Printf("Current app name: %s\n", viper.GetString("app.name"))
    fmt.Printf("Current port: %d\n", viper.GetInt("app.port"))
    fmt.Printf("Current debug: %t\n", viper.GetBool("app.debug"))
    fmt.Println("---")
}
```

### 6.2 Multiple Configuration Files

```go
package main

import (
    "fmt"
    "log"
    "github.com/spf13/viper"
)

func main() {
    // Load base configuration
    baseConfig := viper.New()
    baseConfig.SetConfigName("config")
    baseConfig.SetConfigType("yaml")
    baseConfig.AddConfigPath(".")
    
    if err := baseConfig.ReadInConfig(); err != nil {
        log.Printf("Base config not found: %v", err)
    }
    
    // Load environment-specific configuration
    envConfig := viper.New()
    envConfig.SetConfigName("config.production") // atau config.development
    envConfig.SetConfigType("yaml")
    envConfig.AddConfigPath(".")
    
    if err := envConfig.ReadInConfig(); err != nil {
        log.Printf("Environment config not found: %v", err)
    }
    
    // Merge configurations (env config overrides base config)
    mergedConfig := viper.New()
    
    // Copy base config
    for key, value := range baseConfig.AllSettings() {
        mergedConfig.Set(key, value)
    }
    
    // Override with env config
    for key, value := range envConfig.AllSettings() {
        mergedConfig.Set(key, value)
    }
    
    // Print merged configuration
    fmt.Println("Merged Configuration:")
    for key, value := range mergedConfig.AllSettings() {
        fmt.Printf("%s: %v\n", key, value)
    }
}
```

### 6.3 Configuration with Command Line Flags

```go
package main

import (
    "flag"
    "fmt"
    "github.com/spf13/viper"
    "github.com/spf13/pflag"
)

func main() {
    // Define command line flags
    pflag.String("app-name", "DefaultApp", "Application name")
    pflag.Int("port", 8080, "Server port")
    pflag.Bool("debug", false, "Enable debug mode")
    pflag.String("config", "config.yaml", "Config file path")
    
    // Parse flags
    pflag.Parse()
    
    // Bind flags to viper
    viper.BindPFlags(pflag.CommandLine)
    
    // Set config file from flag
    configFile := viper.GetString("config")
    viper.SetConfigFile(configFile)
    
    // Try to read config file
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            fmt.Println("Config file not found, using defaults and flags")
        } else {
            fmt.Printf("Error reading config file: %v\n", err)
        }
    }
    
    // Environment variables override config file
    viper.AutomaticEnv()
    
    // Print final configuration
    fmt.Println("Final Configuration:")
    fmt.Printf("App Name: %s\n", viper.GetString("app-name"))
    fmt.Printf("Port: %d\n", viper.GetInt("port"))
    fmt.Printf("Debug: %t\n", viper.GetBool("debug"))
    
    // Priority order: 
    // 1. Command line flags (highest)
    // 2. Environment variables
    // 3. Config file
    // 4. Default values (lowest)
}
```

## 7. Best Practices

### 7.1 Configuration Validation

```go
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/spf13/viper"
)

type Config struct {
    App      AppConfig      `mapstructure:"app" validate:"required"`
    Database DatabaseConfig `mapstructure:"database" validate:"required"`
}

type AppConfig struct {
    Name string `mapstructure:"name" validate:"required,min=1"`
    Port int    `mapstructure:"port" validate:"required,min=1,max=65535"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host" validate:"required"`
    Port     int    `mapstructure:"port" validate:"required,min=1,max=65535"`
    Username string `mapstructure:"username" validate:"required"`
    Password string `mapstructure:"password" validate:"required"`
}

func loadAndValidateConfig() (*Config, error) {
    // Setup viper
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()
    
    // Set required defaults
    setConfigDefaults()
    
    // Try to read config
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("config file error: %w", err)
        }
    }
    
    // Unmarshal to struct
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("unmarshal error: %w", err)
    }
    
    // Validate configuration
    if err := validateConfig(&config); err != nil {
        return nil, fmt.Errorf("config validation error: %w", err)
    }
    
    return &config, nil
}

func setConfigDefaults() {
    viper.SetDefault("app.name", "MyApp")
    viper.SetDefault("app.port", 8080)
    viper.SetDefault("database.host", "localhost")
    viper.SetDefault("database.port", 5432)
}

func validateConfig(config *Config) error {
    // Validate app name
    if config.App.Name == "" {
        return fmt.Errorf("app name is required")
    }
    
    // Validate port range
    if config.App.Port < 1 || config.App.Port > 65535 {
        return fmt.Errorf("app port must be between 1 and 65535")
    }
    
    // Validate database config
    if config.Database.Host == "" {
        return fmt.Errorf("database host is required")
    }
    
    if config.Database.Username == "" {
        return fmt.Errorf("database username is required")
    }
    
    if config.Database.Password == "" {
        return fmt.Errorf("database password is required")
    }
    
    return nil
}

func main() {
    config, err := loadAndValidateConfig()
    if err != nil {
        log.Fatalf("Configuration error: %v", err)
        os.Exit(1)
    }
    
    fmt.Println("Configuration loaded and validated successfully!")
    fmt.Printf("App: %s on port %d\n", config.App.Name, config.App.Port)
    fmt.Printf("Database: %s@%s:%d\n", 
        config.Database.Username, config.Database.Host, config.Database.Port)
}
```

### 7.2 Configuration Manager Pattern

```go
package main

import (
    "fmt"
    "log"
    "sync"
    "github.com/spf13/viper"
)

// ConfigManager mengelola konfigurasi aplikasi
type ConfigManager struct {
    viper  *viper.Viper
    config *Config
    mutex  sync.RWMutex
}

// NewConfigManager membuat instance baru ConfigManager
func NewConfigManager() *ConfigManager {
    v := viper.New()
    
    // Setup viper
    v.SetConfigName("config")
    v.SetConfigType("yaml")
    v.AddConfigPath(".")
    v.AddConfigPath("./config")
    v.AutomaticEnv()
    
    return &ConfigManager{
        viper: v,
    }
}

// Load memuat konfigurasi
func (cm *ConfigManager) Load() error {
    cm.mutex.Lock()
    defer cm.mutex.Unlock()
    
    // Set defaults
    cm.setDefaults()
    
    // Read config file
    if err := cm.viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return fmt.Errorf("failed to read config: %w", err)
        }
        log.Println("No config file found, using defaults and env vars")
    }
    
    // Unmarshal to struct
    var config Config
    if err := cm.viper.Unmarshal(&config); err != nil {
        return fmt.Errorf("failed to unmarshal config: %w", err)
    }
    
    cm.config = &config
    return nil
}

// Get mengembalikan konfigurasi saat ini
func (cm *ConfigManager) Get() *Config {
    cm.mutex.RLock()
    defer cm.mutex.RUnlock()
    return cm.config
}

// GetString mendapatkan string value
func (cm *ConfigManager) GetString(key string) string {
    cm.mutex.RLock()
    defer cm.mutex.RUnlock()
    return cm.viper.GetString(key)
}

// GetInt mendapatkan int value
func (cm *ConfigManager) GetInt(key string) int {
    cm.mutex.RLock()
    defer cm.mutex.RUnlock()
    return cm.viper.GetInt(key)
}

// GetBool mendapatkan bool value
func (cm *ConfigManager) GetBool(key string) bool {
    cm.mutex.RLock()
    defer cm.mutex.RUnlock()
    return cm.viper.GetBool(key)
}

// Reload memuat ulang konfigurasi
func (cm *ConfigManager) Reload() error {
    return cm.Load()
}

func (cm *ConfigManager) setDefaults() {
    cm.viper.SetDefault("app.name", "MyApp")
    cm.viper.SetDefault("app.port", 8080)
    cm.viper.SetDefault("app.debug", false)
    cm.viper.SetDefault("database.host", "localhost")
    cm.viper.SetDefault("database.port", 5432)
}

// Global config manager instance
var configManager *ConfigManager

// GetConfig mengembalikan config manager global
func GetConfig() *ConfigManager {
    return configManager
}

func init() {
    configManager = NewConfigManager()
    if err := configManager.Load(); err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
}

func main() {
    // Menggunakan config manager
    config := GetConfig().Get()
    
    fmt.Printf("App: %s\n", config.App.Name)
    fmt.Printf("Port: %d\n", config.App.Port)
    
    // Atau akses langsung dengan key
    debugMode := GetConfig().GetBool("app.debug")
    fmt.Printf("Debug: %t\n", debugMode)
}
```

## 8. Use Cases Praktis

### 8.1 Web Server Configuration

```go
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
    "github.com/spf13/viper"
)

type ServerConfig struct {
    Host         string        `mapstructure:"host"`
    Port         int           `mapstructure:"port"`
    ReadTimeout  time.Duration `mapstructure:"read_timeout"`
    WriteTimeout time.Duration `mapstructure:"write_timeout"`
    IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

type Config struct {
    Server ServerConfig `mapstructure:"server"`
    Debug  bool         `mapstructure:"debug"`
}

func main() {
    config := loadServerConfig()
    server := createServer(config)
    
    // Graceful shutdown
    go func() {
        fmt.Printf("Server starting on %s:%d\n", config.Server.Host, config.Server.Port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()
    
    // Wait for interrupt signal
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    
    fmt.Println("Shutting down server...")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Printf("Server shutdown error: %v", err)
    }
}

func loadServerConfig() *Config {
    viper.SetConfigName("server")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    
    // Set defaults
    viper.SetDefault("server.host", "0.0.0.0")
    viper.SetDefault("server.port", 8080)
    viper.SetDefault("server.read_timeout", "30s")
    viper.SetDefault("server.write_timeout", "30s")
    viper.SetDefault("server.idle_timeout", "120s")
    viper.SetDefault("debug", false)
    
    viper.AutomaticEnv()
    
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            log.Fatalf("Config error: %v", err)
        }
        log.Println("Using default configuration")
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatalf("Unmarshal error: %v", err)
    }
    
    return &config
}

func createServer(config *Config) *http.Server {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from %s!", viper.GetString("server.host"))
    })
    
    mux.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintf(w, `{
            "host": "%s",
            "port": %d,
            "debug": %t
        }`, config.Server.Host, config.Server.Port, config.Debug)
    })
    
    addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
    
    return &http.Server{
        Addr:         addr,
        Handler:      mux,
        ReadTimeout:  config.Server.ReadTimeout,
        WriteTimeout: config.Server.WriteTimeout,
        IdleTimeout:  config.Server.IdleTimeout,
    }
}
```

### 8.2 Database Connection Pool Configuration

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"
    "github.com/spf13/viper"
    _ "github.com/lib/pq" // PostgreSQL driver
)

type DatabaseConfig struct {
    Host            string        `mapstructure:"host"`
    Port            int           `mapstructure:"port"`
    Username        string        `mapstructure:"username"`
    Password        string        `mapstructure:"password"`
    Database        string        `mapstructure:"database"`
    SSLMode         string        `mapstructure:"ssl_mode"`
    MaxOpenConns    int           `mapstructure:"max_open_conns"`
    MaxIdleConns    int           `mapstructure:"max_idle_conns"`
    ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
    ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

func main() {
    config := loadDBConfig()
    db := connectDatabase(config)
    defer db.Close()
    
    // Test connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Database ping failed: %v", err)
    }
    
    fmt.Println("Database connected successfully!")
    printConnectionStats(db)
}

func loadDBConfig() *DatabaseConfig {
    viper.SetConfigName("database")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    
    // Set database defaults
    viper.SetDefault("host", "localhost")
    viper.SetDefault("port", 5432)
    viper.SetDefault("username", "postgres")
    viper.SetDefault("password", "")
    viper.SetDefault("database", "myapp")
    viper.SetDefault("ssl_mode", "disable")
    viper.SetDefault("max_open_conns", 25)
    viper.SetDefault("max_idle_conns", 25)
    viper.SetDefault("conn_max_lifetime", "5m")
    viper.SetDefault("conn_max_idle_time", "5m")
    
    // Environment variables override
    viper.AutomaticEnv()
    viper.SetEnvPrefix("DB")
    
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            log.Fatalf("Database config error: %v", err)
        }
        log.Println("Using default database configuration")
    }
    
    var config DatabaseConfig
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatalf("Database unmarshal error: %v", err)
    }
    
    return &config
}

func connectDatabase(config *DatabaseConfig) *sql.DB {
    // Build connection string
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        config.Host, config.Port, config.Username, 
        config.Password, config.Database, config.SSLMode)
    
    // Open connection
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    
    // Configure connection pool
    db.SetMaxOpenConns(config.MaxOpenConns)
    db.SetMaxIdleConns(config.MaxIdleConns)
    db.SetConnMaxLifetime(config.ConnMaxLifetime)
    db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
    
    return db
}

func printConnectionStats(db *sql.DB) {
    stats := db.Stats()
    fmt.Printf("Database Stats:\n")
    fmt.Printf("  Max Open Connections: %d\n", stats.MaxOpenConnections)
    fmt.Printf("  Open Connections: %d\n", stats.OpenConnections)
    fmt.Printf("  In Use: %d\n", stats.InUse)
    fmt.Printf("  Idle: %d\n", stats.Idle)
}
```

## 9. Tips dan Best Practices

### 9.1 Struktur Project yang Baik

```
project/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   ├── config.go          # Configuration struct dan loader
│   ├── config.yaml        # Development config
│   ├── config.prod.yaml   # Production config
│   └── config.test.yaml   # Test config
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
└── go.mod
```

### 9.2 Configuration Package

```go
// config/config.go
package config

import (
    "fmt"
    "log"
    "strings"
    "github.com/spf13/viper"
)

type Config struct {
    Environment string        `mapstructure:"environment"`
    App         AppConfig     `mapstructure:"app"`
    Database    DatabaseConfig `mapstructure:"database"`
    Redis       RedisConfig   `mapstructure:"redis"`
    Logger      LoggerConfig  `mapstructure:"logger"`
}

type AppConfig struct {
    Name        string `mapstructure:"name"`
    Version     string `mapstructure:"version"`
    Port        int    `mapstructure:"port"`
    Host        string `mapstructure:"host"`
    Debug       bool   `mapstructure:"debug"`
    APIPrefix   string `mapstructure:"api_prefix"`
}

type DatabaseConfig struct {
    Host            string `mapstructure:"host"`
    Port            int    `mapstructure:"port"`
    Username        string `mapstructure:"username"`
    Password        string `mapstructure:"password"`
    Database        string `mapstructure:"database"`
    SSLMode         string `mapstructure:"ssl_mode"`
    MaxOpenConns    int    `mapstructure:"max_open_conns"`
    MaxIdleConns    int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type LoggerConfig struct {
    Level      string `mapstructure:"level"`
    Format     string `mapstructure:"format"`
    Output     string `mapstructure:"output"`
    Filename   string `mapstructure:"filename"`
    MaxSize    int    `mapstructure:"max_size"`
    MaxBackups int    `mapstructure:"max_backups"`
    MaxAge     int    `mapstructure:"max_age"`
}

var cfg *Config

// Load memuat konfigurasi berdasarkan environment
func Load(env string) (*Config, error) {
    v := viper.New()
    
    // Set environment
    if env == "" {
        env = "development"
    }
    
    // Configuration file name
    configName := "config"
    if env != "development" {
        configName = fmt.Sprintf("config.%s", env)
    }
    
    v.SetConfigName(configName)
    v.SetConfigType("yaml")
    v.AddConfigPath("./config")
    v.AddConfigPath(".")
    
    // Environment variables
    v.AutomaticEnv()
    v.SetEnvPrefix("MYAPP")
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    
    // Set defaults
    setDefaults(v, env)
    
    // Read config file
    if err := v.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
            return nil, fmt.Errorf("failed to read config file: %w", err)
        }
        log.Printf("Config file not found, using defaults for %s", env)
    } else {
        log.Printf("Using config file: %s", v.ConfigFileUsed())
    }
    
    // Unmarshal
    var config Config
    if err := v.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %w", err)
    }
    
    config.Environment = env
    
    // Validate
    if err := validateConfig(&config); err != nil {
        return nil, fmt.Errorf("config validation failed: %w", err)
    }
    
    cfg = &config
    return cfg, nil
}

// Get mengembalikan konfigurasi global
func Get() *Config {
    if cfg == nil {
        log.Fatal("Configuration not loaded. Call Load() first.")
    }
    return cfg
}

func setDefaults(v *viper.Viper, env string) {
    // App defaults
    v.SetDefault("app.name", "MyApp")
    v.SetDefault("app.version", "1.0.0")
    v.SetDefault("app.host", "0.0.0.0")
    v.SetDefault("app.port", 8080)
    v.SetDefault("app.debug", env == "development")
    v.SetDefault("app.api_prefix", "/api/v1")
    
    // Database defaults
    v.SetDefault("database.host", "localhost")
    v.SetDefault("database.port", 5432)
    v.SetDefault("database.username", "postgres")
    v.SetDefault("database.password", "")
    v.SetDefault("database.database", "myapp")
    v.SetDefault("database.ssl_mode", "disable")
    v.SetDefault("database.max_open_conns", 25)
    v.SetDefault("database.max_idle_conns", 25)
    
    // Redis defaults
    v.SetDefault("redis.host", "localhost")
    v.SetDefault("redis.port", 6379)
    v.SetDefault("redis.password", "")
    v.SetDefault("redis.db", 0)
    
    // Logger defaults
    logLevel := "info"
    if env == "development" {
        logLevel = "debug"
    }
    v.SetDefault("logger.level", logLevel)
    v.SetDefault("logger.format", "json")
    v.SetDefault("logger.output", "stdout")
}

func validateConfig(config *Config) error {
    if config.App.Name == "" {
        return fmt.Errorf("app name is required")
    }
    
    if config.App.Port < 1 || config.App.Port > 65535 {
        return fmt.Errorf("app port must be between 1 and 65535")
    }
    
    if config.Database.Host == "" {
        return fmt.Errorf("database host is required")
    }
    
    // Add more validations as needed
    return nil
}

// DatabaseURL mengembalikan database connection URL
func (c *Config) DatabaseURL() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
        c.Database.Username,
        c.Database.Password,
        c.Database.Host,
        c.Database.Port,
        c.Database.Database,
        c.Database.SSLMode,
    )
}

// ServerAddr mengembalikan server address
func (c *Config) ServerAddr() string {
    return fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)
}

// RedisAddr mengembalikan Redis address
func (c *Config) RedisAddr() string {
    return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}

// IsProduction mengecek apakah environment production
func (c *Config) IsProduction() bool {
    return c.Environment == "production"
}

// IsDevelopment mengecek apakah environment development
func (c *Config) IsDevelopment() bool {
    return c.Environment == "development"
}
```

### 9.3 Main Application

```go
// cmd/server/main.go
package main

import (
    "flag"
    "log"
    "os"
    "myapp/config"
    "myapp/internal/server"
)

func main() {
    // Parse command line flags
    var env = flag.String("env", "development", "Environment (development, staging, production)")
    flag.Parse()
    
    // Override with environment variable if set
    if envVar := os.Getenv("APP_ENV"); envVar != "" {
        *env = envVar
    }
    
    // Load configuration
    cfg, err := config.Load(*env)
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    log.Printf("Starting %s v%s in %s mode", 
        cfg.App.Name, cfg.App.Version, cfg.Environment)
    
    // Start server
    srv := server.New(cfg)
    if err := srv.Start(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
```

## 10. Common Pitfalls dan Solusinya

### 10.1 Problem: Konfigurasi Tidak Ter-load

```go
// SALAH - tidak set config path
viper.SetConfigName("config")
viper.ReadInConfig() // Error: config file not found

// BENAR - set config path terlebih dahulu
viper.SetConfigName("config")
viper.AddConfigPath(".")
viper.AddConfigPath("./config")
if err := viper.ReadInConfig(); err != nil {
    // Handle error properly
    log.Printf("Config file not found: %v", err)
}
```

### 10.2 Problem: Environment Variables Tidak Override

```go
// SALAH - tidak aktifkan AutomaticEnv
viper.SetConfigName("config")
viper.ReadInConfig()
// Environment variables tidak akan dibaca

// BENAR - aktifkan AutomaticEnv
viper.SetConfigName("config")
viper.AutomaticEnv() // Penting!
viper.ReadInConfig()
```

### 10.3 Problem: Type Conversion Error

```go
// SALAH - tidak handle type conversion
port := viper.Get("port") // return interface{}
server := fmt.Sprintf(":%d", port) // panic: interface conversion

// BENAR - gunakan type-specific getter
port := viper.GetInt("port") // return int
server := fmt.Sprintf(":%d", port) // OK

// ATAU handle dengan type assertion
portInterface := viper.Get("port")
if port, ok := portInterface.(int); ok {
    server := fmt.Sprintf(":%d", port)
} else {
    log.Fatal("Port is not an integer")
}
```

### 10.4 Problem: Nested Configuration

```yaml
# config.yaml
database:
  primary:
    host: "localhost"
    port: 5432
  replica:
    host: "replica.localhost"
    port: 5432
```

```go
// SALAH - akses nested config
host := viper.GetString("database.primary.host") // OK
port := viper.GetString("database.primary.port") // Return string "5432"

// BENAR - akses dengan tipe yang tepat
host := viper.GetString("database.primary.host")
port := viper.GetInt("database.primary.port") // Return int 5432

// ATAU gunakan sub-config
dbPrimary := viper.Sub("database.primary")
if dbPrimary != nil {
    host := dbPrimary.GetString("host")
    port := dbPrimary.GetInt("port")
}
```

## 11. Testing Configuration

```go
package config

import (
    "os"
    "testing"
    "github.com/spf13/viper"
    "github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
    tests := []struct {
        name        string
        env         string
        envVars     map[string]string
        expectError bool
        expectedApp AppConfig
    }{
        {
            name: "development config",
            env:  "development",
            expectedApp: AppConfig{
                Name:  "MyApp",
                Port:  8080,
                Debug: true,
            },
        },
        {
            name: "production config",
            env:  "production",
            expectedApp: AppConfig{
                Name:  "MyApp",
                Port:  8080,
                Debug: false,
            },
        },
        {
            name: "environment variable override",
            env:  "development",
            envVars: map[string]string{
                "MYAPP_APP_PORT": "9000",
                "MYAPP_APP_NAME": "TestApp",
            },
            expectedApp: AppConfig{
                Name:  "TestApp",
                Port:  9000,
                Debug: true,
            },
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Set environment variables
            for key, value := range tt.envVars {
                os.Setenv(key, value)
                defer os.Unsetenv(key)
            }
            
            // Load config
            cfg, err := Load(tt.env)
            
            if tt.expectError {
                assert.Error(t, err)
                return
            }
            
            assert.NoError(t, err)
            assert.Equal(t, tt.expectedApp.Name, cfg.App.Name)
            assert.Equal(t, tt.expectedApp.Port, cfg.App.Port)
            assert.Equal(t, tt.expectedApp.Debug, cfg.App.Debug)
        })
    }
}

func TestConfigValidation(t *testing.T) {
    tests := []struct {
        name        string
        config      Config
        expectError bool
        errorMsg    string
    }{
        {
            name: "valid config",
            config: Config{
                App: AppConfig{
                    Name: "TestApp",
                    Port: 8080,
                },
                Database: DatabaseConfig{
                    Host: "localhost",
                },
            },
            expectError: false,
        },
        {
            name: "empty app name",
            config: Config{
                App: AppConfig{
                    Name: "",
                    Port: 8080,
                },
            },
            expectError: true,
            errorMsg:    "app name is required",
        },
        {
            name: "invalid port",
            config: Config{
                App: AppConfig{
                    Name: "TestApp",
                    Port: 70000,
                },
            },
            expectError: true,
            errorMsg:    "app port must be between 1 and 65535",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := validateConfig(&tt.config)
            
            if tt.expectError {
                assert.Error(t, err)
                assert.Contains(t, err.Error(), tt.errorMsg)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}

// TestViperWithMockConfig untuk testing tanpa file
func TestViperWithMockConfig(t *testing.T) {
    v := viper.New()
    
    // Set values directly instead of reading from file
    v.Set("app.name", "TestApp")
    v.Set("app.port", 8080)
    v.Set("app.debug", true)
    
    assert.Equal(t, "TestApp", v.GetString("app.name"))
    assert.Equal(t, 8080, v.GetInt("app.port"))
    assert.Equal(t, true, v.GetBool("app.debug"))
}
```

## 12. Kesimpulan

Viper adalah tool yang sangat powerful untuk mengelola konfigurasi dalam aplikasi Go. Keunggulan utamanya:

1. **Fleksibilitas**: Mendukung berbagai format dan sumber konfigurasi
2. **Hierarki**: Mendukung prioritas konfigurasi (flags > env vars > config file > defaults)
3. **Type Safety**: Konversi tipe otomatis dengan getter methods
4. **Live Reload**: Watch file configuration untuk perubahan real-time
5. **Environment Aware**: Easy switching antara development, staging, production

### Best Practices yang Harus Diingat:

- Selalu set default values untuk semua konfigurasi
- Gunakan struct mapping untuk type safety
- Validate konfigurasi setelah load
- Buat configuration package terpisah untuk reusability
- Handle error dengan proper error handling
- Test konfigurasi dengan berbagai skenario
- Dokumentasikan semua konfigurasi yang available

Dengan memahami konsep-konsep ini, Anda sudah siap menggunakan Viper untuk mengelola konfigurasi aplikasi Go dengan baik dan benar!