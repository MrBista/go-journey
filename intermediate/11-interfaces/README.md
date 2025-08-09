# Panduan Lengkap Interface di Go untuk Pemula

## 1. Apa itu Interface?

Interface di Go adalah sebuah kontrak yang mendefinisikan method-method yang harus diimplementasikan oleh sebuah type. Interface memungkinkan kita untuk menulis kode yang lebih fleksibel dan dapat digunakan kembali.

**Karakteristik Interface di Go:**
- Interface adalah implicit (tidak perlu keyword `implements`)
- Jika sebuah type memiliki semua method yang didefinisikan interface, maka type tersebut mengimplementasikan interface
- Interface kosong `interface{}` dapat menerima value apapun

## 2. Syntax Dasar Interface

```go
// Mendefinisikan interface
type NamaInterface interface {
    Method1() return_type
    Method2(parameter_type) return_type
    // ... method lainnya
}
```

## 3. Contoh Sederhana

```go
package main

import "fmt"

// Definisi interface
type Animal interface {
    Sound() string
    Move() string
}

// Type Dog
type Dog struct {
    Name string
}

// Implementasi method Sound untuk Dog
func (d Dog) Sound() string {
    return "Woof!"
}

// Implementasi method Move untuk Dog
func (d Dog) Move() string {
    return "Running"
}

// Type Cat
type Cat struct {
    Name string
}

// Implementasi method untuk Cat
func (c Cat) Sound() string {
    return "Meow!"
}

func (c Cat) Move() string {
    return "Jumping"
}

// Fungsi yang menggunakan interface
func AnimalBehavior(a Animal) {
    fmt.Printf("Sound: %s, Movement: %s\n", a.Sound(), a.Move())
}

func main() {
    dog := Dog{Name: "Buddy"}
    cat := Cat{Name: "Whiskers"}
    
    // Kedua type ini mengimplementasikan interface Animal
    AnimalBehavior(dog) // Output: Sound: Woof!, Movement: Running
    AnimalBehavior(cat) // Output: Sound: Meow!, Movement: Jumping
}
```

## 4. Interface Kosong (Empty Interface)

```go
package main

import "fmt"

func PrintAnything(value interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

func main() {
    PrintAnything(42)           // int
    PrintAnything("Hello")      // string
    PrintAnything([]int{1,2,3}) // slice
    PrintAnything(true)         // bool
}
```

## 5. Type Assertion dan Type Switch

### Type Assertion

Ada dua bentuk:

- Tidak aman: value.(Type) - akan panic jika gagal <br/>
- Aman: value, ok := interface.(Type) - mengembalikan nilai dan boolean

```go
package main

import "fmt"

func CheckType(value interface{}) {
    // Type assertion dengan checking
    if str, ok := value.(string); ok {
        fmt.Printf("String value: %s\n", str)
    } else if num, ok := value.(int); ok {
        fmt.Printf("Integer value: %d\n", num)
    } else {
        fmt.Printf("Unknown type: %T\n", value)
    }
}

func main() {
    CheckType("Hello World")
    CheckType(123)
    CheckType(3.14)
}
```

### Type Switch
```go
package main

import "fmt"

func ProcessValue(value interface{}) {
    switch v := value.(type) {
    case string:
        fmt.Printf("String: %s (length: %d)\n", v, len(v))
    case int:
        fmt.Printf("Integer: %d (squared: %d)\n", v, v*v)
    case float64:
        fmt.Printf("Float: %.2f\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func main() {
    ProcessValue("Go Programming")
    ProcessValue(42)
    ProcessValue(3.14159)
    ProcessValue(true)
    ProcessValue([]int{1, 2, 3})
}
```

## 6. Interface Composition

Interface bisa di-compose dari interface lain:

```go
package main

import "fmt"

// Interface dasar
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

// Interface composition
type ReadWriter interface {
    Reader  // embed Reader interface
    Writer  // embed Writer interface
}

// Implementation
type File struct {
    content string
}

func (f *File) Read() string {
    return f.content
}

func (f *File) Write(data string) {
    f.content = data
}

func ProcessFile(rw ReadWriter) {
    rw.Write("Hello from Go!")
    fmt.Println("Content:", rw.Read())
}

func main() {
    file := &File{}
    ProcessFile(file)
}
```

## 7. Interface dengan Pointer vs Value Receiver

```go
package main

import "fmt"

type Counter interface {
    Increment()
    GetValue() int
}

type IntCounter struct {
    value int
}

// Method dengan pointer receiver
func (c *IntCounter) Increment() {
    c.value++
}

// Method dengan value receiver
func (c IntCounter) GetValue() int {
    return c.value
}

func main() {
    // Harus menggunakan pointer karena Increment() menggunakan pointer receiver
    counter := &IntCounter{value: 0}
    
    var c Counter = counter
    c.Increment()
    c.Increment()
    fmt.Println("Counter value:", c.GetValue()) // Output: 2
}
```

## 8. Best Practices

### 1. Keep Interfaces Small
```go
// GOOD: Interface kecil dan focused
type Writer interface {
    Write([]byte) (int, error)
}

// AVOID: Interface yang terlalu besar
type MegaInterface interface {
    Read() string
    Write(string)
    Connect()
    Disconnect()
    Parse()
    Validate()
    Process()
}
```

### 2. Accept Interfaces, Return Concrete Types
```go
package main

import "fmt"

type Printer interface {
    Print() string
}

type Document struct {
    content string
}

func (d Document) Print() string {
    return d.content
}

// GOOD: Accept interface sebagai parameter
func ProcessDocument(p Printer) string {
    return "Processed: " + p.Print()
}

// GOOD: Return concrete type
func NewDocument(content string) Document {
    return Document{content: content}
}

func main() {
    doc := NewDocument("Hello World")
    result := ProcessDocument(doc)
    fmt.Println(result)
}
```

### 3. Interface Segregation
```go
// Pisahkan interface berdasarkan tanggung jawab
type EmailSender interface {
    SendEmail(to, subject, body string) error
}

type SMSSender interface {
    SendSMS(to, message string) error
}

type NotificationSender interface {
    EmailSender
    SMSSender
}

// Implementasi bisa memilih interface mana yang akan diimplementasikan
type EmailService struct{}

func (e EmailService) SendEmail(to, subject, body string) error {
    fmt.Printf("Email sent to %s: %s\n", to, subject)
    return nil
}

type SMSService struct{}

func (s SMSService) SendSMS(to, message string) error {
    fmt.Printf("SMS sent to %s: %s\n", to, message)
    return nil
}
```

### 4. Error Interface yang Sering Digunakan
```go
package main

import (
    "fmt"
    "errors"
)

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

// Implementasi error interface
func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}

func ValidateEmail(email string) error {
    if email == "" {
        return ValidationError{
            Field:   "email",
            Message: "email cannot be empty",
        }
    }
    if !strings.Contains(email, "@") {
        return ValidationError{
            Field:   "email",
            Message: "invalid email format",
        }
    }
    return nil
}

func main() {
    err := ValidateEmail("")
    if err != nil {
        fmt.Println("Error:", err)
        
        // Type assertion untuk custom error
        if valErr, ok := err.(ValidationError); ok {
            fmt.Printf("Field: %s, Message: %s\n", valErr.Field, valErr.Message)
        }
    }
}
```

## 9. Interface yang Sering Digunakan di Standard Library

### io.Reader dan io.Writer
```go
package main

import (
    "fmt"
    "strings"
    "io"
)

func ReadData(r io.Reader) string {
    data := make([]byte, 100)
    n, _ := r.Read(data)
    return string(data[:n])
}

func main() {
    // strings.Reader mengimplementasikan io.Reader
    reader := strings.NewReader("Hello from io.Reader!")
    content := ReadData(reader)
    fmt.Println(content)
}
```

### fmt.Stringer
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Implementasi fmt.Stringer interface
func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    fmt.Println(person) // Otomatis memanggil String() method
}
```

## 10. Tips dan Trik

### 1. Gunakan Interface untuk Testing
```go
// Definisi interface untuk database
type UserRepository interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

// Mock implementation untuk testing
type MockUserRepository struct {
    users map[int]*User
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    user := m.users[id]
    if user == nil {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (m *MockUserRepository) SaveUser(user *User) error {
    m.users[user.ID] = user
    return nil
}
```

### 2. Nil Interface vs Nil Value
```go
package main

import "fmt"

type Animal interface {
    Sound() string
}

type Dog struct{}
func (d *Dog) Sound() string { return "Woof!" }

func main() {
    var animal Animal
    fmt.Println("animal == nil:", animal == nil) // true
    
    var dog *Dog = nil
    animal = dog
    fmt.Println("animal == nil:", animal == nil) // false! (nil value, not nil interface)
    
    // Check for nil value in interface
    if animal == nil {
        fmt.Println("Interface is nil")
    } else {
        fmt.Printf("Interface is not nil, type: %T, value: %v\n", animal, animal)
    }
}
```

## 11. Kesalahan Umum yang Harus Dihindari

### 1. Interface Terlalu Besar
```go
// BURUK: Interface terlalu besar
type UserService interface {
    CreateUser(user User) error
    UpdateUser(user User) error
    DeleteUser(id int) error
    GetUser(id int) (User, error)
    GetAllUsers() ([]User, error)
    ValidateUser(user User) error
    SendWelcomeEmail(user User) error
    GenerateReport() (Report, error)
}

// BAIK: Pisahkan berdasarkan tanggung jawab
type UserRepository interface {
    Create(user User) error
    Update(user User) error
    Delete(id int) error
    Get(id int) (User, error)
    GetAll() ([]User, error)
}

type UserValidator interface {
    Validate(user User) error
}

type EmailSender interface {
    SendWelcomeEmail(user User) error
}
```

### 2. Premature Interface Abstraction
```go
// BURUK: Membuat interface sebelum ada kebutuhan
type Calculator interface {
    Add(a, b int) int
}

type SimpleCalculator struct{}
func (c SimpleCalculator) Add(a, b int) int { return a + b }

// BAIK: Buat interface ketika ada multiple implementations atau untuk testing
```

## Kesimpulan

Interface di Go adalah tools yang powerful untuk:
- Membuat kode yang fleksibel dan dapat di-test
- Implementasi polymorphism
- Decoupling antar components
- Membuat API yang clean dan extensible

**Ingat:** Start simple, keep interfaces small, dan gunakan interface ketika benar-benar dibutuhkan!