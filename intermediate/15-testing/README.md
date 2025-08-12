# Panduan Lengkap Testing di Go untuk Pemula

## 1. Konsep Dasar Testing di Go

### Mengapa Testing Penting?
- Memastikan kode berfungsi sesuai ekspektasi
- Mencegah bug di production
- Memudahkan refactoring
- Dokumentasi hidup dari behavior kode
- Meningkatkan confidence dalam deployment

### Struktur File Test di Go
```
project/
├── main.go
├── calculator.go
├── calculator_test.go  // File test untuk calculator.go
├── user/
│   ├── user.go
│   └── user_test.go    // File test untuk user.go
└── go.mod
```

**Aturan Penamaan:**
- File test harus diakhiri dengan `_test.go`
- Function test harus dimulai dengan `Test`
- Function benchmark harus dimulai dengan `Benchmark`

## 2. Unit Testing Dasar

### Struktur Function Test
```go
func TestNamaFunction(t *testing.T) {
    // Arrange - Setup data
    // Act - Jalankan function yang akan ditest
    // Assert - Verifikasi hasil
}
```

### Contoh Sederhana
```go
// calculator.go
package main

func Add(a, b int) int {
    return a + b
}

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

```go
// calculator_test.go
package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    // Test case positive
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestAddNegative(t *testing.T) {
    result := Add(-1, 1)
    expected := 0
    
    if result != expected {
        t.Errorf("Add(-1, 1) = %d; want %d", result, expected)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(10, 2)
    
    if err != nil {
        t.Errorf("Divide(10, 2) returned error: %v", err)
    }
    
    expected := 5.0
    if result != expected {
        t.Errorf("Divide(10, 2) = %f; want %f", result, expected)
    }
}

func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)
    
    if err == nil {
        t.Error("Divide(10, 0) should return error")
    }
}
```

## 3. Table-Driven Tests (Best Practice)

Table-driven tests adalah pattern yang sangat umum di Go untuk menguji multiple scenarios:

```go
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -1, -2},
        {"mixed numbers", -1, 1, 0},
        {"zero", 0, 5, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

## 4. Testing dengan Subtests

Subtests memungkinkan Anda mengelompokkan test cases:

```go
func TestCalculator(t *testing.T) {
    t.Run("Addition", func(t *testing.T) {
        result := Add(2, 3)
        if result != 5 {
            t.Errorf("expected 5, got %d", result)
        }
    })
    
    t.Run("Division", func(t *testing.T) {
        t.Run("Normal", func(t *testing.T) {
            result, err := Divide(10, 2)
            if err != nil || result != 5 {
                t.Errorf("expected 5, got %f with error %v", result, err)
            }
        })
        
        t.Run("ByZero", func(t *testing.T) {
            _, err := Divide(10, 0)
            if err == nil {
                t.Error("expected error for division by zero")
            }
        })
    })
}
```

## 5. Mocking dan Interface Testing

### Menggunakan Interface untuk Testing
```go
// user.go
type UserRepository interface {
    GetUser(id int) (*User, error)
    SaveUser(user *User) error
}

type UserService struct {
    repo UserRepository
}

func (s *UserService) GetUserName(id int) (string, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        return "", err
    }
    return user.Name, nil
}

type User struct {
    ID   int
    Name string
}
```

```go
// user_test.go
type MockUserRepository struct {
    users map[int]*User
    err   error
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    if m.err != nil {
        return nil, m.err
    }
    user, exists := m.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (m *MockUserRepository) SaveUser(user *User) error {
    return m.err
}

func TestUserService_GetUserName(t *testing.T) {
    tests := []struct {
        name          string
        userID        int
        mockUsers     map[int]*User
        mockError     error
        expectedName  string
        expectedError bool
    }{
        {
            name:   "successful get user",
            userID: 1,
            mockUsers: map[int]*User{
                1: {ID: 1, Name: "John Doe"},
            },
            expectedName:  "John Doe",
            expectedError: false,
        },
        {
            name:          "user not found",
            userID:        999,
            mockUsers:     map[int]*User{},
            expectedName:  "",
            expectedError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockRepo := &MockUserRepository{
                users: tt.mockUsers,
                err:   tt.mockError,
            }
            
            service := &UserService{repo: mockRepo}
            
            name, err := service.GetUserName(tt.userID)
            
            if tt.expectedError && err == nil {
                t.Error("expected error but got none")
            }
            
            if !tt.expectedError && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
            
            if name != tt.expectedName {
                t.Errorf("expected name %s, got %s", tt.expectedName, name)
            }
        })
    }
}
```

## 6. Testing HTTP Handlers

```go
// handler.go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

```go
// handler_test.go
func TestHelloHandler(t *testing.T) {
    tests := []struct {
        name           string
        queryParam     string
        expectedStatus int
        expectedBody   string
    }{
        {"default greeting", "", http.StatusOK, "Hello, World!"},
        {"custom name", "name=John", http.StatusOK, "Hello, John!"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/?"+tt.queryParam, nil)
            rr := httptest.NewRecorder()
            
            HelloHandler(rr, req)
            
            if rr.Code != tt.expectedStatus {
                t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
            }
            
            if rr.Body.String() != tt.expectedBody {
                t.Errorf("expected body %s, got %s", tt.expectedBody, rr.Body.String())
            }
        })
    }
}
```

## 7. Benchmark Testing

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkAddLarge(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(999999, 888888)
    }
}
```

Jalankan benchmark dengan:
```bash
go test -bench=.
go test -bench=BenchmarkAdd
```

## 8. Test Coverage

Melihat coverage testing:
```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 9. Best Practices

### 1. Penamaan Test yang Descriptive
```go
// ❌ Bad
func TestAdd(t *testing.T) {}

// ✅ Good
func TestAdd_PositiveNumbers_ReturnsSum(t *testing.T) {}
func TestAdd_WhenBothNegative_ReturnsNegativeSum(t *testing.T) {}
```

### 2. Gunakan Helper Functions
```go
func TestSomething(t *testing.T) {
    result := SomeFunction()
    assertEqual(t, expected, result)
}

func assertEqual(t *testing.T, expected, actual interface{}) {
    t.Helper() // Menandai sebagai helper function
    if expected != actual {
        t.Errorf("expected %v, got %v", expected, actual)
    }
}
```

### 3. Setup dan Cleanup
```go
func TestMain(m *testing.M) {
    // Setup before all tests
    setup()
    
    // Run tests
    code := m.Run()
    
    // Cleanup after all tests
    cleanup()
    
    os.Exit(code)
}

func TestWithCleanup(t *testing.T) {
    // Setup for this specific test
    tempFile := createTempFile()
    
    // Cleanup setelah test selesai
    t.Cleanup(func() {
        os.Remove(tempFile)
    })
    
    // Test logic here
}
```

### 4. Test Isolation
```go
// ❌ Bad - Tests depend on each other
var counter int

func TestIncrement1(t *testing.T) {
    counter++
    if counter != 1 {
        t.Error("expected 1")
    }
}

func TestIncrement2(t *testing.T) {
    counter++
    if counter != 2 {  // Depends on previous test
        t.Error("expected 2")
    }
}

// ✅ Good - Each test is independent
func TestIncrement(t *testing.T) {
    counter := 0
    counter++
    if counter != 1 {
        t.Error("expected 1")
    }
}
```

## 10. Command Line Testing

```bash
# Run all tests
go test

# Run tests in specific package
go test ./user

# Run specific test
go test -run TestAdd

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run benchmarks
go test -bench=.

# Run tests in parallel
go test -parallel 4

# Run tests multiple times
go test -count=10
```

## 11. Testing dengan External Dependencies

### Database Testing
```go
func TestUserRepository_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    
    // Setup test database
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    repo := NewUserRepository(db)
    
    // Test logic here
}
```

Run integration tests:
```bash
go test -short  # Skip integration tests
go test         # Run all tests including integration
```

## 12. Error Testing Patterns

```go
func TestFunction_ErrorCases(t *testing.T) {
    tests := []struct {
        name          string
        input         string
        expectedError string
    }{
        {"empty input", "", "input cannot be empty"},
        {"invalid format", "abc", "invalid format"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := SomeFunction(tt.input)
            
            if err == nil {
                t.Fatal("expected error but got none")
            }
            
            if !strings.Contains(err.Error(), tt.expectedError) {
                t.Errorf("expected error containing %q, got %q", 
                    tt.expectedError, err.Error())
            }
        })
    }
}
```

## Tips Tambahan

1. **Mulai dengan test sederhana** - Jangan langsung membuat test yang kompleks
2. **Test happy path dulu** - Test scenario normal sebelum edge cases
3. **Satu assert per test** - Lebih mudah di-debug
4. **Mock external dependencies** - Database, API calls, file system
5. **Gunakan testify library** - Untuk assertion yang lebih mudah dibaca
6. **Test-driven development** - Tulis test sebelum implementation
7. **Keep tests fast** - Slow tests jarang dijalankan
8. **Test behavior, not implementation** - Test apa yang dilakukan, bukan bagaimana

Dengan mengikuti panduan ini, Anda akan dapat menulis test yang efektif dan maintainable di Go!