# Panduan Lengkap Methods di Go untuk Pemula

## Apa itu Methods?

Methods di Go adalah fungsi yang memiliki **receiver** khusus. Receiver adalah tipe data yang "memiliki" method tersebut. Berbeda dengan function biasa, method terikat pada tipe data tertentu.

## Anatomi Method

```go
func (receiver ReceiverType) MethodName(parameters) ReturnType {
    // body method
}
```

- `receiver`: variabel yang menerima method
- `ReceiverType`: tipe data yang memiliki method
- `MethodName`: nama method
- `parameters`: parameter input (opsional)
- `ReturnType`: tipe data return (opsional)

## 1. Method Dasar dengan Value Receiver

### Contoh Sederhana

```go
package main

import "fmt"

// Definisi struct
type Person struct {
    Name string
    Age  int
}

// Method dengan value receiver
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

func (p Person) GetAge() int {
    return p.Age
}

func main() {
    person := Person{Name: "John", Age: 25}
    
    fmt.Println(person.Greet())  // Output: Hello, my name is John
    fmt.Println(person.GetAge()) // Output: 25
}
```

## 2. Pointer Receiver vs Value Receiver

### Value Receiver
- Menerima **copy** dari value
- Tidak bisa mengubah nilai asli
- Cocok untuk operasi read-only

```go
type Counter struct {
    count int
}

// Value receiver - TIDAK bisa mengubah count asli
func (c Counter) IncrementWrong() {
    c.count++ // Hanya mengubah copy
}

func (c Counter) GetCount() int {
    return c.count
}
```

### Pointer Receiver
- Menerima **pointer** ke value asli
- Bisa mengubah nilai asli
- Cocok untuk operasi yang mengubah state

```go
// Pointer receiver - BISA mengubah count asli
func (c *Counter) Increment() {
    c.count++
}

func (c *Counter) Decrement() {
    c.count--
}

func (c *Counter) Reset() {
    c.count = 0
}

// Contoh penggunaan
func main() {
    counter := Counter{count: 0}
    
    counter.IncrementWrong()
    fmt.Println(counter.GetCount()) // Output: 0 (tidak berubah)
    
    counter.Increment()
    fmt.Println(counter.GetCount()) // Output: 1 (berubah)
}
```

## 3. Methods pada Built-in Types

Anda tidak bisa langsung menambah method pada built-in types, tapi bisa membuat type alias:

```go
type MyString string
type MyInt int

func (s MyString) IsPalindrome() bool {
    str := string(s)
    runes := []rune(str)
    n := len(runes)
    
    for i := 0; i < n/2; i++ {
        if runes[i] != runes[n-1-i] {
            return false
        }
    }
    return true
}

func (num MyInt) IsEven() bool {
    return int(num)%2 == 0
}

// Penggunaan
func main() {
    word := MyString("radar")
    fmt.Println(word.IsPalindrome()) // true
    
    number := MyInt(10)
    fmt.Println(number.IsEven()) // true
}
```

## 4. Method Chaining

Method chaining memungkinkan pemanggilan method secara beruntun:

```go
type Calculator struct {
    result float64
}

func (c *Calculator) Add(n float64) *Calculator {
    c.result += n
    return c
}

func (c *Calculator) Subtract(n float64) *Calculator {
    c.result -= n
    return c
}

func (c *Calculator) Multiply(n float64) *Calculator {
    c.result *= n
    return c
}

func (c *Calculator) GetResult() float64 {
    return c.result
}

// Penggunaan
func main() {
    calc := &Calculator{}
    result := calc.Add(10).Subtract(5).Multiply(2).GetResult()
    fmt.Println(result) // Output: 10
}
```

## 5. Methods dengan Interface

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

// Methods untuk Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Methods untuk Circle
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14159 * c.Radius
}

// Function yang menerima interface
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 4}
    
    PrintShapeInfo(rect)
    PrintShapeInfo(circle)
}
```

## 6. Contoh Real-World: User Management

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

type User struct {
    ID        int
    Username  string
    Email     string
    CreatedAt time.Time
    IsActive  bool
}

// Constructor-like function
func NewUser(username, email string) *User {
    return &User{
        Username:  username,
        Email:     email,
        CreatedAt: time.Now(),
        IsActive:  true,
    }
}

// Validation methods
func (u *User) ValidateEmail() bool {
    return strings.Contains(u.Email, "@")
}

func (u *User) ValidateUsername() bool {
    return len(u.Username) >= 3
}

// State management methods
func (u *User) Activate() {
    u.IsActive = true
}

func (u *User) Deactivate() {
    u.IsActive = false
}

// Information methods
func (u User) GetFullInfo() string {
    status := "Active"
    if !u.IsActive {
        status = "Inactive"
    }
    
    return fmt.Sprintf("User: %s (%s) - Status: %s - Created: %s",
        u.Username, u.Email, status, u.CreatedAt.Format("2006-01-02"))
}

func (u User) IsValid() bool {
    return u.ValidateEmail() && u.ValidateUsername()
}

// Usage example
func main() {
    user := NewUser("john_doe", "john@example.com")
    
    if user.IsValid() {
        fmt.Println("User is valid")
        fmt.Println(user.GetFullInfo())
    }
    
    user.Deactivate()
    fmt.Println(user.GetFullInfo())
}
```

## Best Practices

### 1. Kapan Menggunakan Pointer Receiver
```go
// GUNAKAN pointer receiver jika:
// - Method mengubah state receiver
// - Receiver adalah struct besar (untuk efisiensi)
// - Konsistensi (jika ada method lain yang butuh pointer)

type LargeStruct struct {
    data [1000]int
}

// Good: menggunakan pointer untuk struct besar
func (ls *LargeStruct) Process() {
    // processing...
}

// Good: mengubah state
func (ls *LargeStruct) UpdateData(index, value int) {
    ls.data[index] = value
}
```

### 2. Naming Conventions
```go
// GOOD: verb untuk action methods
func (u *User) Activate() {}
func (u *User) Save() {}
func (u *User) Delete() {}

// GOOD: Get/Is/Has untuk query methods
func (u User) GetName() string { return u.name }
func (u User) IsActive() bool { return u.active }
func (u User) HasPermission() bool { return u.hasPermission }

// AVOID: noun untuk method names
func (u *User) Status() {} // kurang jelas
func (u *User) GetStatus() {} // lebih jelas
```

### 3. Method Organization
```go
type BankAccount struct {
    balance float64
    owner   string
}

// Group related methods together
// 1. Constructor
func NewBankAccount(owner string) *BankAccount {
    return &BankAccount{owner: owner, balance: 0}
}

// 2. Query methods (value receiver)
func (ba BankAccount) GetBalance() float64 {
    return ba.balance
}

func (ba BankAccount) GetOwner() string {
    return ba.owner
}

// 3. Action methods (pointer receiver)
func (ba *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("amount must be positive")
    }
    ba.balance += amount
    return nil
}

func (ba *BankAccount) Withdraw(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("amount must be positive")
    }
    if amount > ba.balance {
        return fmt.Errorf("insufficient funds")
    }
    ba.balance -= amount
    return nil
}
```

### 4. Error Handling dalam Methods
```go
type FileProcessor struct {
    filename string
}

func (fp *FileProcessor) Process() error {
    if fp.filename == "" {
        return fmt.Errorf("filename cannot be empty")
    }
    
    // processing logic...
    return nil
}

// Usage with error handling
func main() {
    processor := &FileProcessor{}
    
    if err := processor.Process(); err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Println("Processing successful")
}
```

### 5. Method Sets dan Interface Implementation
```go
// Interface dengan method yang sering digunakan
type Validator interface {
    Validate() error
}

type Saver interface {
    Save() error
}

type Entity interface {
    Validator
    Saver
}

// Implementasi
type Product struct {
    Name  string
    Price float64
}

func (p Product) Validate() error {
    if p.Name == "" {
        return fmt.Errorf("name is required")
    }
    if p.Price < 0 {
        return fmt.Errorf("price must be non-negative")
    }
    return nil
}

func (p *Product) Save() error {
    if err := p.Validate(); err != nil {
        return err
    }
    // save logic...
    return nil
}
```

## Tips Penting

1. **Konsistensi Receiver**: Jika satu method menggunakan pointer receiver, gunakan pointer receiver untuk semua methods pada type yang sama

2. **Performance**: Gunakan pointer receiver untuk struct besar untuk menghindari copying

3. **Mutability**: Gunakan pointer receiver jika method perlu mengubah state

4. **Interface Satisfaction**: Pastikan method signature sesuai dengan interface yang ingin diimplementasikan

5. **Testing**: Selalu test methods dengan berbagai skenario input

Method adalah salah satu fitur paling powerful di Go yang memungkinkan Anda menulis code yang lebih terorganisir dan mudah dipahami. Praktikkan contoh-contoh di atas dan eksperimen dengan membuat methods sendiri!