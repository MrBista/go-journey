# Panduan Lengkap JSON Handling di Go untuk Pemula

## 1. Pengenalan JSON di Go

JSON (JavaScript Object Notation) adalah format pertukaran data yang sangat populer dalam pengembangan aplikasi modern. Go menyediakan package `encoding/json` yang sangat powerful untuk menangani operasi JSON.

### Mengapa JSON Penting?
- **API Communication**: Hampir semua REST API menggunakan JSON
- **Configuration Files**: Banyak aplikasi menggunakan JSON untuk konfigurasi
- **Data Storage**: NoSQL databases sering menyimpan data dalam format JSON
- **Frontend-Backend Communication**: Web applications menggunakan JSON untuk komunikasi

## 2. Import Package JSON

```go
import (
    "encoding/json"
    "fmt"
    "log"
)
```

## 3. Struktur Data Dasar untuk JSON

### Definisi Struct dengan JSON Tags

```go
type Person struct {
    // JSON tag untuk mapping field
    Name     string `json:"name"`           // field akan bernama "name" di JSON
    Age      int    `json:"age"`            // field akan bernama "age" di JSON
    Email    string `json:"email"`          // field akan bernama "email" di JSON
    IsActive bool   `json:"is_active"`      // field akan bernama "is_active" di JSON
}

// Struct yang lebih kompleks
type Address struct {
    Street   string `json:"street"`
    City     string `json:"city"`
    Country  string `json:"country"`
    ZipCode  string `json:"zip_code"`
}

type Employee struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Position    string  `json:"position"`
    Salary      float64 `json:"salary"`
    Address     Address `json:"address"`        // Nested struct
    Skills      []string `json:"skills"`       // Array/slice
    IsManager   bool    `json:"is_manager"`
    // Field yang tidak akan di-serialize ke JSON (private atau diabaikan)
    internalID  string  `json:"-"`             // Diabaikan dengan "-"
}
```

### JSON Tags yang Penting

```go
type User struct {
    Name     string `json:"name"`                    // Normal mapping
    Email    string `json:"email,omitempty"`         // Hilangkan jika kosong
    Password string `json:"-"`                       // Jangan serialize
    Age      int    `json:"age,string"`              // Konversi ke string
    Bio      string `json:"bio,omitempty"`           // Hilangkan jika kosong
    Active   *bool  `json:"active,omitempty"`        // Pointer untuk nullable
}
```

**Penjelasan JSON Tags:**
- `json:"name"`: Field akan muncul sebagai "name" di JSON
- `json:"email,omitempty"`: Jika field kosong (zero value), tidak akan muncul di JSON
- `json:"-"`: Field tidak akan di-serialize ke JSON sama sekali
- `json:"age,string"`: Angka akan di-convert ke string di JSON
- Menggunakan pointer (`*bool`) memungkinkan nilai null di JSON

## 4. Marshaling (Go Struct → JSON String)

### Marshal Dasar

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    InStock     bool    `json:"in_stock"`
    Tags        []string `json:"tags"`
    Description string  `json:"description,omitempty"`
}

func main() {
    // Membuat instance struct
    product := Product{
        ID:      1,
        Name:    "Laptop Gaming",
        Price:   15000000.50,
        InStock: true,
        Tags:    []string{"electronics", "computer", "gaming"},
        // Description dibiarkan kosong untuk demo omitempty
    }

    // Marshal struct ke JSON
    jsonData, err := json.Marshal(product)
    if err != nil {
        log.Fatal("Error marshaling JSON:", err)
    }

    fmt.Println("JSON String:")
    fmt.Println(string(jsonData))
    
    // Output akan seperti ini:
    // {"id":1,"name":"Laptop Gaming","price":15000000.5,"in_stock":true,"tags":["electronics","computer","gaming"]}
    // Perhatikan bahwa "description" tidak muncul karena kosong dan menggunakan omitempty
}
```

### Marshal dengan Pretty Printing

```go
func prettyPrintJSON(data interface{}) {
    // MarshalIndent untuk formatting JSON yang rapi
    jsonData, err := json.MarshalIndent(data, "", "  ") // prefix: "", indent: "  "
    if err != nil {
        log.Fatal("Error marshaling JSON:", err)
    }
    
    fmt.Println(string(jsonData))
}

// Penggunaan:
prettyPrintJSON(product)

// Output:
// {
//   "id": 1,
//   "name": "Laptop Gaming",
//   "price": 15000000.5,
//   "in_stock": true,
//   "tags": [
//     "electronics",
//     "computer",
//     "gaming"
//   ]
// }
```

## 5. Unmarshaling (JSON String → Go Struct)

### Unmarshal Dasar

```go
func unmarshalExample() {
    // JSON string yang akan di-unmarshal
    jsonString := `{
        "id": 2,
        "name": "Smartphone",
        "price": 8000000.0,
        "in_stock": false,
        "tags": ["electronics", "mobile", "communication"]
    }`

    // Variabel untuk menampung hasil unmarshal
    var product Product

    // Unmarshal JSON string ke struct
    err := json.Unmarshal([]byte(jsonString), &product)
    if err != nil {
        log.Fatal("Error unmarshaling JSON:", err)
    }

    fmt.Printf("Parsed Product:\n")
    fmt.Printf("ID: %d\n", product.ID)
    fmt.Printf("Name: %s\n", product.Name)
    fmt.Printf("Price: %.2f\n", product.Price)
    fmt.Printf("In Stock: %t\n", product.InStock)
    fmt.Printf("Tags: %v\n", product.Tags)
}
```

### Unmarshal dengan Error Handling yang Baik

```go
func safeUnmarshal(jsonData []byte, target interface{}) error {
    // Validasi input
    if len(jsonData) == 0 {
        return fmt.Errorf("empty JSON data")
    }

    // Unmarshal dengan error handling
    if err := json.Unmarshal(jsonData, target); err != nil {
        // Memberikan informasi error yang lebih detail
        return fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    return nil
}

// Penggunaan:
var product Product
jsonString := `{"id": 1, "name": "Test Product"}`

if err := safeUnmarshal([]byte(jsonString), &product); err != nil {
    log.Printf("Unmarshal error: %v", err)
} else {
    fmt.Printf("Successfully parsed: %+v\n", product)
}
```

## 6. Bekerja dengan Dynamic JSON (interface{})

### Ketika Struktur JSON Tidak Diketahui

```go
func dynamicJSONExample() {
    jsonString := `{
        "user": {
            "name": "John Doe",
            "age": 30,
            "hobbies": ["reading", "gaming", "cooking"]
        },
        "metadata": {
            "created_at": "2023-01-15T10:30:00Z",
            "version": 1.2
        }
    }`

    // Menggunakan map[string]interface{} untuk struktur yang dinamis
    var data map[string]interface{}

    err := json.Unmarshal([]byte(jsonString), &data)
    if err != nil {
        log.Fatal("Error unmarshaling:", err)
    }

    // Mengakses data nested
    user := data["user"].(map[string]interface{})
    name := user["name"].(string)
    age := user["age"].(float64) // JSON number selalu float64 di Go
    hobbies := user["hobbies"].([]interface{})

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %.0f\n", age)
    fmt.Print("Hobbies: ")
    for _, hobby := range hobbies {
        fmt.Printf("%s ", hobby.(string))
    }
    fmt.Println()
}
```

### Type Assertion yang Aman

```go
func safeTypeAssertion(data interface{}, key string) {
    if dataMap, ok := data.(map[string]interface{}); ok {
        if value, exists := dataMap[key]; exists {
            switch v := value.(type) {
            case string:
                fmt.Printf("%s (string): %s\n", key, v)
            case float64:
                fmt.Printf("%s (number): %.2f\n", key, v)
            case bool:
                fmt.Printf("%s (boolean): %t\n", key, v)
            case []interface{}:
                fmt.Printf("%s (array): %v\n", key, v)
            case map[string]interface{}:
                fmt.Printf("%s (object): %v\n", key, v)
            default:
                fmt.Printf("%s (unknown type): %v\n", key, v)
            }
        }
    }
}
```

## 7. JSON dengan Slice dan Array

### Marshal/Unmarshal Slice

```go
type Book struct {
    Title  string   `json:"title"`
    Author string   `json:"author"`
    ISBN   string   `json:"isbn"`
    Tags   []string `json:"tags"`
}

func sliceExample() {
    books := []Book{
        {
            Title:  "Clean Code",
            Author: "Robert C. Martin",
            ISBN:   "978-0132350884",
            Tags:   []string{"programming", "software-engineering"},
        },
        {
            Title:  "The Go Programming Language",
            Author: "Alan Donovan",
            ISBN:   "978-0134190440",
            Tags:   []string{"go", "programming"},
        },
    }

    // Marshal slice ke JSON
    jsonData, err := json.MarshalIndent(books, "", "  ")
    if err != nil {
        log.Fatal("Error marshaling:", err)
    }

    fmt.Println("Books JSON:")
    fmt.Println(string(jsonData))

    // Unmarshal kembali
    var parsedBooks []Book
    err = json.Unmarshal(jsonData, &parsedBooks)
    if err != nil {
        log.Fatal("Error unmarshaling:", err)
    }

    fmt.Printf("\nParsed %d books\n", len(parsedBooks))
    for i, book := range parsedBooks {
        fmt.Printf("Book %d: %s by %s\n", i+1, book.Title, book.Author)
    }
}
```

## 8. Custom JSON Marshaling/Unmarshaling

### Implementasi MarshalJSON dan UnmarshalJSON

```go
import "time"

type CustomDate struct {
    time.Time
}

// Custom marshaler untuk format tanggal khusus
func (cd CustomDate) MarshalJSON() ([]byte, error) {
    // Format tanggal khusus: "2006-01-02"
    formatted := fmt.Sprintf(`"%s"`, cd.Time.Format("2006-01-02"))
    return []byte(formatted), nil
}

// Custom unmarshaler untuk format tanggal khusus
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
    // Remove quotes dari JSON string
    str := string(data)
    if len(str) < 2 || str[0] != '"' || str[len(str)-1] != '"' {
        return fmt.Errorf("invalid date format")
    }
    str = str[1 : len(str)-1]

    // Parse tanggal
    parsedTime, err := time.Parse("2006-01-02", str)
    if err != nil {
        return err
    }

    cd.Time = parsedTime
    return nil
}

type Event struct {
    Name string     `json:"name"`
    Date CustomDate `json:"date"`
}

func customMarshalExample() {
    event := Event{
        Name: "Go Workshop",
        Date: CustomDate{time.Now()},
    }

    jsonData, err := json.MarshalIndent(event, "", "  ")
    if err != nil {
        log.Fatal("Error marshaling:", err)
    }

    fmt.Println("Custom marshaled JSON:")
    fmt.Println(string(jsonData))

    // Unmarshal kembali
    var parsedEvent Event
    err = json.Unmarshal(jsonData, &parsedEvent)
    if err != nil {
        log.Fatal("Error unmarshaling:", err)
    }

    fmt.Printf("Parsed event: %s on %s\n", 
        parsedEvent.Name, 
        parsedEvent.Date.Format("2006-01-02"))
}
```

## 9. Error Handling Best Practices

### Comprehensive Error Handling

```go
type JSONError struct {
    Operation string
    Err       error
    Data      string
}

func (je *JSONError) Error() string {
    return fmt.Sprintf("JSON %s error: %v (data: %s)", je.Operation, je.Err, je.Data)
}

func safeJSONMarshal(v interface{}) ([]byte, error) {
    data, err := json.Marshal(v)
    if err != nil {
        return nil, &JSONError{
            Operation: "marshal",
            Err:       err,
            Data:      fmt.Sprintf("%+v", v),
        }
    }
    return data, nil
}

func safeJSONUnmarshal(data []byte, v interface{}) error {
    if err := json.Unmarshal(data, v); err != nil {
        return &JSONError{
            Operation: "unmarshal",
            Err:       err,
            Data:      string(data),
        }
    }
    return nil
}
```

## 10. Use Cases Praktis

### 1. HTTP API Response Handler

```go
type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func createAPIResponse(status, message string, data interface{}, err error) []byte {
    response := APIResponse{
        Status:  status,
        Message: message,
        Data:    data,
    }

    if err != nil {
        response.Error = err.Error()
    }

    jsonData, _ := json.MarshalIndent(response, "", "  ")
    return jsonData
}

// Contoh penggunaan:
func handleUserRequest() {
    user := User{Name: "John", Email: "john@example.com"}
    
    // Success response
    successResponse := createAPIResponse("success", "User retrieved", user, nil)
    fmt.Println("Success Response:")
    fmt.Println(string(successResponse))

    // Error response
    errorResponse := createAPIResponse("error", "User not found", nil, 
        fmt.Errorf("user with ID 123 not found"))
    fmt.Println("\nError Response:")
    fmt.Println(string(errorResponse))
}
```

### 2. Configuration File Handler

```go
type Config struct {
    Server   ServerConfig   `json:"server"`
    Database DatabaseConfig `json:"database"`
    Logging  LoggingConfig  `json:"logging"`
}

type ServerConfig struct {
    Host string `json:"host"`
    Port int    `json:"port"`
    SSL  bool   `json:"ssl"`
}

type DatabaseConfig struct {
    Driver   string `json:"driver"`
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Database string `json:"database"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoggingConfig struct {
    Level    string `json:"level"`
    FilePath string `json:"file_path"`
    MaxSize  int    `json:"max_size"`
}

func loadConfig(filename string) (*Config, error) {
    // Dalam implementasi nyata, baca dari file
    configJSON := `{
        "server": {
            "host": "localhost",
            "port": 8080,
            "ssl": false
        },
        "database": {
            "driver": "postgres",
            "host": "localhost",
            "port": 5432,
            "database": "myapp",
            "username": "user",
            "password": "password"
        },
        "logging": {
            "level": "info",
            "file_path": "/var/log/app.log",
            "max_size": 100
        }
    }`

    var config Config
    err := json.Unmarshal([]byte(configJSON), &config)
    if err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }

    return &config, nil
}
```

## 11. Best Practices

### 1. Naming Conventions
```go
// ✅ Good: Konsisten dengan JSON naming conventions
type User struct {
    UserID    int    `json:"user_id"`    // snake_case untuk JSON
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
}

// ❌ Avoid: Inconsistent naming
type BadUser struct {
    UserID    int    `json:"userId"`      // camelCase mixed with snake_case
    FirstName string `json:"firstName"`
    LastName  string `json:"last_name"`  // Inconsistent
}
```

### 2. Error Handling
```go
// ✅ Good: Proper error handling
func processJSON(data []byte) error {
    var result MyStruct
    
    if err := json.Unmarshal(data, &result); err != nil {
        return fmt.Errorf("failed to unmarshal JSON: %w", err)
    }
    
    // Process result...
    return nil
}

// ❌ Bad: Ignoring errors
func badProcessJSON(data []byte) {
    var result MyStruct
    json.Unmarshal(data, &result) // Ignoring error
    // Process result... (might be invalid!)
}
```

### 3. Validation
```go
type ValidatedUser struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

func (u *ValidatedUser) Validate() error {
    if u.Name == "" {
        return fmt.Errorf("name is required")
    }
    if u.Email == "" {
        return fmt.Errorf("email is required")
    }
    if u.Age < 0 || u.Age > 150 {
        return fmt.Errorf("age must be between 0 and 150")
    }
    return nil
}

func processValidatedJSON(data []byte) error {
    var user ValidatedUser
    
    if err := json.Unmarshal(data, &user); err != nil {
        return fmt.Errorf("invalid JSON: %w", err)
    }
    
    if err := user.Validate(); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    // Process valid user...
    return nil
}
```

### 4. Performance Tips
```go
// ✅ Good: Reuse structs when possible
var userPool = sync.Pool{
    New: func() interface{} {
        return &User{}
    },
}

func efficientProcessing(data []byte) error {
    user := userPool.Get().(*User)
    defer userPool.Put(user)
    
    // Reset struct
    *user = User{}
    
    return json.Unmarshal(data, user)
}

// ✅ Good: Use streaming for large data
func streamProcessLargeJSON(reader io.Reader) error {
    decoder := json.NewDecoder(reader)
    
    for decoder.More() {
        var item MyStruct
        if err := decoder.Decode(&item); err != nil {
            return err
        }
        // Process item...
    }
    
    return nil
}
```

## 12. Common Pitfalls dan Solusinya

### 1. Zero Values dan omitempty
```go
type Settings struct {
    Notifications bool `json:"notifications,omitempty"` // ❌ Problem: false akan dihilangkan
    Volume       int  `json:"volume,omitempty"`        // ❌ Problem: 0 akan dihilangkan
}

// ✅ Solution: Gunakan pointer untuk distinguish antara zero value dan unset
type BetterSettings struct {
    Notifications *bool `json:"notifications,omitempty"` // nil = unset, false = explicitly false
    Volume       *int  `json:"volume,omitempty"`        // nil = unset, 0 = explicitly 0
}
```

### 2. Integer/Float Precision
```go
// ❌ Problem: JSON numbers adalah float64
func badNumberHandling() {
    jsonStr := `{"large_id": 9007199254740992}` // Larger than int53
    
    var data map[string]interface{}
    json.Unmarshal([]byte(jsonStr), &data)
    
    id := data["large_id"].(float64) // Precision loss!
    fmt.Printf("ID: %.0f\n", id)
}

// ✅ Solution: Use string for large integers
type SafeID struct {
    LargeID string `json:"large_id"` // Keep as string to preserve precision
}
```

## 13. Testing JSON Code

```go
import "testing"

func TestUserMarshalUnmarshal(t *testing.T) {
    original := User{
        Name:  "Test User",
        Email: "test@example.com",
        Age:   25,
    }

    // Test Marshal
    data, err := json.Marshal(original)
    if err != nil {
        t.Fatalf("Marshal failed: %v", err)
    }

    // Test Unmarshal
    var parsed User
    err = json.Unmarshal(data, &parsed)
    if err != nil {
        t.Fatalf("Unmarshal failed: %v", err)
    }

    // Verify data integrity
    if parsed.Name != original.Name {
        t.Errorf("Name mismatch: got %s, want %s", parsed.Name, original.Name)
    }
    if parsed.Email != original.Email {
        t.Errorf("Email mismatch: got %s, want %s", parsed.Email, original.Email)
    }
    if parsed.Age != original.Age {
        t.Errorf("Age mismatch: got %d, want %d", parsed.Age, original.Age)
    }
}

func TestInvalidJSON(t *testing.T) {
    invalidJSON := `{"name": "test", "age": "invalid"}`
    
    var user User
    err := json.Unmarshal([]byte(invalidJSON), &user)
    
    if err == nil {
        t.Error("Expected error for invalid JSON, but got none")
    }
}
```

Dengan panduan lengkap ini, Anda sekarang memiliki pemahaman yang solid tentang JSON handling di Go. Mulai dengan contoh-contoh sederhana, praktikkan, dan secara bertahap tingkatkan ke use case yang lebih kompleks!