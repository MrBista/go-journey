# Panduan Lengkap Reflection di Go untuk Pemula

## Apa itu Reflection?

Reflection adalah kemampuan sebuah program untuk memeriksa struktur dan perilaku dirinya sendiri saat runtime (waktu eksekusi). Dalam Go, reflection memungkinkan kita untuk:

- Memeriksa tipe data variabel
- Mengakses dan memodifikasi nilai variabel
- Memanggil method secara dinamis
- Membuat instance baru dari tipe tertentu
- Memeriksa struct fields, tags, dan metadata lainnya

## Kapan Menggunakan Reflection?

Reflection sangat berguna dalam situasi berikut:
- Membuat library yang bekerja dengan berbagai tipe data (seperti JSON marshaling)
- Dependency injection
- ORM (Object-Relational Mapping)
- Testing framework
- Configuration parsing
- API serialization/deserialization

## Package reflect

Go menyediakan package `reflect` untuk melakukan reflection. Dua konsep utama dalam reflection Go adalah:

1. **Type** - informasi tentang tipe data
2. **Value** - informasi tentang nilai aktual

```go
import "reflect"

// reflect.TypeOf() - mendapatkan informasi tipe
// reflect.ValueOf() - mendapatkan informasi nilai
```

## Konsep Dasar Reflection

### 1. Mendapatkan Type dan Value

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x int = 42
    var s string = "hello"
    
    // Mendapatkan type
    fmt.Println("Type of x:", reflect.TypeOf(x))     // int
    fmt.Println("Type of s:", reflect.TypeOf(s))     // string
    
    // Mendapatkan value
    fmt.Println("Value of x:", reflect.ValueOf(x))   // 42
    fmt.Println("Value of s:", reflect.ValueOf(s))   // hello
    
    // Mendapatkan kind (kategori tipe)
    fmt.Println("Kind of x:", reflect.TypeOf(x).Kind())   // int
    fmt.Println("Kind of s:", reflect.TypeOf(s).Kind())   // string
}
```

### 2. Kind vs Type

**Kind** adalah kategori dasar dari tipe data, sedangkan **Type** adalah tipe spesifik:

```go
type MyInt int
type Person struct {
    Name string
}

func main() {
    var mi MyInt = 10
    var p Person = Person{Name: "John"}
    
    // Type menunjukkan tipe spesifik
    fmt.Println("Type of mi:", reflect.TypeOf(mi))    // main.MyInt
    fmt.Println("Type of p:", reflect.TypeOf(p))      // main.Person
    
    // Kind menunjukkan kategori dasar
    fmt.Println("Kind of mi:", reflect.TypeOf(mi).Kind()) // int
    fmt.Println("Kind of p:", reflect.TypeOf(p).Kind())   // struct
}
```

## Reflection dengan Struct

### 1. Mengakses Field Struct

```go
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    ID       int    `json:"id" validate:"required"`
    Name     string `json:"name" validate:"required,min=2"`
    Email    string `json:"email" validate:"required,email"`
    IsActive bool   `json:"is_active"`
}

func inspectStruct(obj interface{}) {
    // Mendapatkan type dan value
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)
    
    fmt.Printf("Type: %s, Kind: %s\n", t.Name(), t.Kind())
    fmt.Printf("Number of fields: %d\n\n", t.NumField())
    
    // Iterasi semua field
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        fmt.Printf("Field %d:\n", i+1)
        fmt.Printf("  Name: %s\n", field.Name)
        fmt.Printf("  Type: %s\n", field.Type)
        fmt.Printf("  Kind: %s\n", field.Type.Kind())
        fmt.Printf("  Value: %v\n", value.Interface())
        fmt.Printf("  JSON Tag: %s\n", field.Tag.Get("json"))
        fmt.Printf("  Validate Tag: %s\n\n", field.Tag.Get("validate"))
    }
}

func main() {
    user := User{
        ID:       1,
        Name:     "John Doe",
        Email:    "john@example.com",
        IsActive: true,
    }
    
    inspectStruct(user)
}
```

### 2. Mengubah Nilai Field (Menggunakan Pointer)

```go
func modifyStruct(obj interface{}) {
    // Harus menggunakan pointer untuk bisa mengubah nilai
    v := reflect.ValueOf(obj)
    
    // Pastikan ini adalah pointer
    if v.Kind() != reflect.Ptr {
        fmt.Println("Harus menggunakan pointer untuk mengubah nilai")
        return
    }
    
    // Dapatkan element dari pointer
    v = v.Elem()
    
    // Pastikan ini adalah struct
    if v.Kind() != reflect.Struct {
        fmt.Println("Object harus berupa struct")
        return
    }
    
    // Cari field berdasarkan nama
    nameField := v.FieldByName("Name")
    if nameField.IsValid() && nameField.CanSet() {
        if nameField.Kind() == reflect.String {
            nameField.SetString("Jane Doe")
            fmt.Println("Name berhasil diubah")
        }
    }
    
    // Mengubah field ID
    idField := v.FieldByName("ID")
    if idField.IsValid() && idField.CanSet() {
        if idField.Kind() == reflect.Int {
            idField.SetInt(999)
            fmt.Println("ID berhasil diubah")
        }
    }
}

func main() {
    user := User{
        ID:   1,
        Name: "John Doe",
    }
    
    fmt.Println("Before:", user)
    
    // Harus pass sebagai pointer
    modifyStruct(&user)
    
    fmt.Println("After:", user)
}
```

## Reflection dengan Method

### 1. Memanggil Method Secara Dinamis

```go
package main

import (
    "fmt"
    "reflect"
)

type Calculator struct {
    Result float64
}

func (c *Calculator) Add(a, b float64) float64 {
    result := a + b
    c.Result = result
    return result
}

func (c *Calculator) Multiply(a, b float64) float64 {
    result := a * b
    c.Result = result
    return result
}

func (c Calculator) GetResult() float64 {
    return c.Result
}

func callMethod(obj interface{}, methodName string, args ...interface{}) {
    v := reflect.ValueOf(obj)
    
    // Cari method berdasarkan nama
    method := v.MethodByName(methodName)
    if !method.IsValid() {
        fmt.Printf("Method %s tidak ditemukan\n", methodName)
        return
    }
    
    // Konversi arguments ke reflect.Value
    var reflectArgs []reflect.Value
    for _, arg := range args {
        reflectArgs = append(reflectArgs, reflect.ValueOf(arg))
    }
    
    // Panggil method
    results := method.Call(reflectArgs)
    
    // Tampilkan hasil
    fmt.Printf("Method %s dipanggil\n", methodName)
    if len(results) > 0 {
        for i, result := range results {
            fmt.Printf("  Return value %d: %v\n", i+1, result.Interface())
        }
    }
}

func main() {
    calc := &Calculator{}
    
    // Panggil method Add secara dinamis
    callMethod(calc, "Add", 10.0, 5.0)
    
    // Panggil method GetResult
    callMethod(calc, "GetResult")
    
    // Panggil method Multiply
    callMethod(calc, "Multiply", 3.0, 4.0)
    
    callMethod(calc, "GetResult")
}
```

## Use Cases Praktis

### 1. JSON-like Serialization

```go
package main

import (
    "fmt"
    "reflect"
    "strings"
)

func structToMap(obj interface{}) map[string]interface{} {
    result := make(map[string]interface{})
    
    v := reflect.ValueOf(obj)
    t := reflect.TypeOf(obj)
    
    // Jika pointer, ambil element-nya
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    // Pastikan ini struct
    if v.Kind() != reflect.Struct {
        return nil
    }
    
    // Iterasi semua field
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        // Skip unexported fields
        if !field.IsExported() {
            continue
        }
        
        // Gunakan json tag jika ada, atau nama field
        key := field.Tag.Get("json")
        if key == "" {
            key = strings.ToLower(field.Name)
        }
        
        // Skip field dengan tag "-"
        if key == "-" {
            continue
        }
        
        result[key] = value.Interface()
    }
    
    return result
}

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    Description string  `json:"description"`
    InternalID  string  `json:"-"` // Field ini akan di-skip
}

func main() {
    product := Product{
        ID:          1,
        Name:        "Laptop",
        Price:       15000000,
        Description: "Gaming Laptop",
        InternalID:  "INTERNAL_123",
    }
    
    data := structToMap(product)
    fmt.Printf("Serialized data: %+v\n", data)
}
```

### 2. Simple Validation Framework

```go
package main

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

func validateStruct(obj interface{}) []string {
    var errors []string
    
    v := reflect.ValueOf(obj)
    t := reflect.TypeOf(obj)
    
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }
    
    if v.Kind() != reflect.Struct {
        return []string{"Object harus berupa struct"}
    }
    
    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        
        validateTag := field.Tag.Get("validate")
        if validateTag == "" {
            continue
        }
        
        fieldName := field.Name
        rules := strings.Split(validateTag, ",")
        
        for _, rule := range rules {
            rule = strings.TrimSpace(rule)
            
            switch {
            case rule == "required":
                if isZeroValue(value) {
                    errors = append(errors, fmt.Sprintf("%s is required", fieldName))
                }
                
            case strings.HasPrefix(rule, "min="):
                minStr := strings.TrimPrefix(rule, "min=")
                min, err := strconv.Atoi(minStr)
                if err != nil {
                    continue
                }
                
                if value.Kind() == reflect.String {
                    if len(value.String()) < min {
                        errors = append(errors, fmt.Sprintf("%s minimum length is %d", fieldName, min))
                    }
                } else if value.Kind() == reflect.Int {
                    if int(value.Int()) < min {
                        errors = append(errors, fmt.Sprintf("%s minimum value is %d", fieldName, min))
                    }
                }
                
            case rule == "email":
                if value.Kind() == reflect.String {
                    email := value.String()
                    if !strings.Contains(email, "@") {
                        errors = append(errors, fmt.Sprintf("%s must be a valid email", fieldName))
                    }
                }
            }
        }
    }
    
    return errors
}

func isZeroValue(v reflect.Value) bool {
    switch v.Kind() {
    case reflect.String:
        return v.String() == ""
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int() == 0
    case reflect.Bool:
        return !v.Bool()
    case reflect.Float32, reflect.Float64:
        return v.Float() == 0
    default:
        return false
    }
}

type RegisterForm struct {
    Username string `validate:"required,min=3"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"required,min=18"`
}

func main() {
    // Valid form
    validForm := RegisterForm{
        Username: "johndoe",
        Email:    "john@example.com",
        Age:      25,
    }
    
    // Invalid form
    invalidForm := RegisterForm{
        Username: "jo", // too short
        Email:    "invalid-email", // no @
        Age:      16, // too young
    }
    
    fmt.Println("Valid form errors:", validateStruct(validForm))
    fmt.Println("Invalid form errors:", validateStruct(invalidForm))
}
```

### 3. Simple Dependency Injection

```go
package main

import (
    "fmt"
    "reflect"
)

type Container struct {
    services map[reflect.Type]reflect.Value
}

func NewContainer() *Container {
    return &Container{
        services: make(map[reflect.Type]reflect.Value),
    }
}

func (c *Container) Register(service interface{}) {
    t := reflect.TypeOf(service)
    v := reflect.ValueOf(service)
    c.services[t] = v
}

func (c *Container) Get(serviceType interface{}) interface{} {
    t := reflect.TypeOf(serviceType).Elem() // Elem() karena kita pass pointer
    if service, ok := c.services[t]; ok {
        return service.Interface()
    }
    return nil
}

func (c *Container) Inject(target interface{}) {
    v := reflect.ValueOf(target)
    if v.Kind() != reflect.Ptr {
        fmt.Println("Target harus berupa pointer")
        return
    }
    
    v = v.Elem()
    t := v.Type()
    
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldType := t.Field(i)
        
        // Cek apakah field memiliki tag "inject"
        if fieldType.Tag.Get("inject") != "true" {
            continue
        }
        
        if !field.CanSet() {
            continue
        }
        
        // Cari service yang cocok
        if service, ok := c.services[field.Type()]; ok {
            field.Set(service)
        }
    }
}

// Services
type DatabaseService struct {
    ConnectionString string
}

func (db *DatabaseService) Connect() {
    fmt.Println("Connected to database:", db.ConnectionString)
}

type EmailService struct {
    SMTPServer string
}

func (email *EmailService) SendEmail(to, subject, body string) {
    fmt.Printf("Sending email to %s via %s\n", to, email.SMTPServer)
}

// Controller yang membutuhkan services
type UserController struct {
    DB    *DatabaseService `inject:"true"`
    Email *EmailService    `inject:"true"`
}

func (uc *UserController) CreateUser(username, email string) {
    uc.DB.Connect()
    fmt.Printf("Creating user: %s\n", username)
    uc.Email.SendEmail(email, "Welcome", "Welcome to our platform!")
}

func main() {
    // Setup container
    container := NewContainer()
    
    // Register services
    container.Register(&DatabaseService{ConnectionString: "localhost:5432"})
    container.Register(&EmailService{SMTPServer: "smtp.gmail.com"})
    
    // Create controller dan inject dependencies
    controller := &UserController{}
    container.Inject(controller)
    
    // Gunakan controller
    controller.CreateUser("johndoe", "john@example.com")
}
```

## Best Practices

### 1. Performance Considerations

```go
// ❌ Buruk - melakukan reflection berulang kali
func slowFunction(objects []interface{}) {
    for _, obj := range objects {
        t := reflect.TypeOf(obj) // Reflection di dalam loop
        // ... process
    }
}

// ✅ Baik - cache hasil reflection
var typeCache = make(map[reflect.Type]bool)

func fastFunction(objects []interface{}) {
    for _, obj := range objects {
        t := reflect.TypeOf(obj)
        if cached, ok := typeCache[t]; ok {
            // Gunakan hasil yang sudah di-cache
            _ = cached
        } else {
            // Lakukan reflection dan cache hasilnya
            result := expensiveReflectionOperation(t)
            typeCache[t] = result
        }
    }
}

func expensiveReflectionOperation(t reflect.Type) bool {
    // Operasi reflection yang mahal
    return t.Kind() == reflect.Struct
}
```

### 2. Error Handling

```go
func safeReflectionOperation(obj interface{}) error {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
        }
    }()
    
    v := reflect.ValueOf(obj)
    
    // Validasi sebelum operasi
    if !v.IsValid() {
        return fmt.Errorf("invalid value")
    }
    
    if v.Kind() != reflect.Struct {
        return fmt.Errorf("expected struct, got %s", v.Kind())
    }
    
    // Lakukan operasi reflection dengan aman
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        if field.CanInterface() {
            fmt.Printf("Field %d: %v\n", i, field.Interface())
        }
    }
    
    return nil
}
```

### 3. Type Safety

```go
// Gunakan type assertion untuk memastikan type safety
func processUser(obj interface{}) {
    // Method 1: Direct type assertion
    if user, ok := obj.(User); ok {
        fmt.Printf("User: %+v\n", user)
        return
    }
    
    // Method 2: Reflection dengan type checking
    v := reflect.ValueOf(obj)
    t := reflect.TypeOf(obj)
    
    // Pastikan ini adalah User type
    var userType reflect.Type = reflect.TypeOf(User{})
    if t == userType {
        // Safe untuk di-cast
        user := obj.(User)
        fmt.Printf("User via reflection: %+v\n", user)
    }
}
```

## Kapan TIDAK Menggunakan Reflection

1. **Performance-critical code** - Reflection lambat dibanding operasi normal
2. **Simple operations** - Jika bisa dilakukan tanpa reflection, lebih baik tidak menggunakan
3. **Compile-time safety** - Reflection menghilangkan type checking saat compile time

## Alternatif Reflection

```go
// ❌ Menggunakan reflection untuk hal sederhana
func printValueBad(obj interface{}) {
    v := reflect.ValueOf(obj)
    fmt.Println(v.Interface())
}

// ✅ Gunakan type assertion atau interface
func printValueGood(obj interface{}) {
    fmt.Println(obj) // fmt.Println sudah handle interface{}
}

// ✅ Atau gunakan generics (Go 1.18+)
func printValueGeneric[T any](value T) {
    fmt.Println(value)
}
```

## Kesimpulan

Reflection adalah tool yang powerful di Go, tetapi harus digunakan dengan bijak. Gunakan reflection ketika:

- Membuat library generic
- Butuh runtime type inspection
- Membuat framework atau tool development
- Tidak ada cara lain yang lebih sederhana

Selalu pertimbangkan performance dan type safety ketika menggunakan reflection. Dalam kebanyakan kasus, interface{} dengan type assertion atau generics (Go 1.18+) bisa menjadi alternatif yang lebih baik.