# Panduan Lengkap Repository Pattern dengan MySQL di Golang

## 1. Apa itu Repository Pattern?

Repository Pattern adalah sebuah design pattern yang menyediakan lapisan abstraksi antara business logic dan data access logic. Pattern ini bertujuan untuk:

- **Memisahkan concern**: Business logic tidak perlu tahu bagaimana data disimpan
- **Testability**: Mudah untuk membuat mock/fake repository untuk testing
- **Flexibility**: Mudah mengganti database tanpa mengubah business logic
- **Clean Code**: Kode lebih terorganisir dan mudah dipahami

### Analogi Sederhana
Bayangkan Anda adalah seorang chef (business logic) yang ingin mengambil bahan makanan. Anda tidak perlu tahu dimana tepatnya bahan itu disimpan di gudang, Anda cukup meminta kepada keeper gudang (repository). Keeper gudang yang akan mencarikan dan memberikan bahan yang Anda butuhkan.

## 2. Struktur Dasar Repository Pattern

```
├── models/          # Data models/entities
├── repositories/    # Repository interfaces dan implementations  
├── services/        # Business logic
└── handlers/        # HTTP handlers
```

## 3. Implementasi Step by Step

### Step 1: Setup Dependencies

Pertama, kita perlu menginstall dependencies yang diperlukan:

```bash
go mod init repository-pattern-example
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
```

### Step 2: Membuat Model/Entity

```go
// models/user.go
package models

import "time"

type User struct {
    ID        int       `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    Email     string    `json:"email" db:"email"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// UserCreateRequest adalah struct untuk data yang diterima saat membuat user baru
type UserCreateRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// UserUpdateRequest adalah struct untuk data yang diterima saat update user
type UserUpdateRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

**Penjelasan Model:**
- Tag `json` digunakan untuk serialization ke JSON (untuk API response)
- Tag `db` digunakan untuk mapping ke column database
- Kita pisahkan struct untuk request dan response untuk keamanan dan clarity

### Step 3: Membuat Repository Interface

```go
// repositories/user_repository.go
package repositories

import (
    "context"
    "your-project/models"
)

// UserRepository adalah interface yang mendefinisikan kontrak
// untuk semua operasi yang berhubungan dengan User
type UserRepository interface {
    // Create menambahkan user baru ke database
    Create(ctx context.Context, user *models.User) error
    
    // GetByID mengambil user berdasarkan ID
    GetByID(ctx context.Context, id int) (*models.User, error)
    
    // GetByEmail mengambil user berdasarkan email
    GetByEmail(ctx context.Context, email string) (*models.User, error)
    
    // GetAll mengambil semua user dengan pagination
    GetAll(ctx context.Context, limit, offset int) ([]*models.User, error)
    
    // Update mengupdate data user
    Update(ctx context.Context, user *models.User) error
    
    // Delete menghapus user berdasarkan ID
    Delete(ctx context.Context, id int) error
    
    // Count menghitung total user (untuk pagination)
    Count(ctx context.Context) (int, error)
}
```

**Kenapa menggunakan Interface?**
- **Abstraksi**: Business logic tidak perlu tahu implementasi database
- **Testing**: Kita bisa membuat mock repository untuk unit test
- **Flexibility**: Bisa ganti dari MySQL ke PostgreSQL tanpa ubah business logic
- **Dependency Inversion**: High-level modules tidak bergantung pada low-level modules

### Step 4: Implementasi MySQL Repository

```go
// repositories/mysql_user_repository.go
package repositories

import (
    "context"
    "database/sql"
    "fmt"
    "time"
    "your-project/models"
)

// mysqlUserRepository adalah implementasi UserRepository untuk MySQL
type mysqlUserRepository struct {
    db *sql.DB
}

// NewMySQLUserRepository adalah constructor untuk membuat MySQL user repository
func NewMySQLUserRepository(db *sql.DB) UserRepository {
    return &mysqlUserRepository{
        db: db,
    }
}

// Create mengimplementasikan method Create dari interface UserRepository
func (r *mysqlUserRepository) Create(ctx context.Context, user *models.User) error {
    query := `
        INSERT INTO users (name, email, created_at, updated_at) 
        VALUES (?, ?, ?, ?)
    `
    
    now := time.Now()
    user.CreatedAt = now
    user.UpdatedAt = now
    
    result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, now, now)
    if err != nil {
        return fmt.Errorf("failed to create user: %w", err)
    }
    
    // Ambil ID yang baru saja di-generate oleh MySQL
    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("failed to get last insert id: %w", err)
    }
    
    user.ID = int(id)
    return nil
}

func (r *mysqlUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    query := `
        SELECT id, name, email, created_at, updated_at 
        FROM users 
        WHERE id = ?
    `
    
    user := &models.User{}
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with id %d not found", id)
        }
        return nil, fmt.Errorf("failed to get user by id: %w", err)
    }
    
    return user, nil
}

func (r *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
    query := `
        SELECT id, name, email, created_at, updated_at 
        FROM users 
        WHERE email = ?
    `
    
    user := &models.User{}
    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with email %s not found", email)
        }
        return nil, fmt.Errorf("failed to get user by email: %w", err)
    }
    
    return user, nil
}

func (r *mysqlUserRepository) GetAll(ctx context.Context, limit, offset int) ([]*models.User, error) {
    query := `
        SELECT id, name, email, created_at, updated_at 
        FROM users 
        ORDER BY created_at DESC
        LIMIT ? OFFSET ?
    `
    
    rows, err := r.db.QueryContext(ctx, query, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("failed to get users: %w", err)
    }
    defer rows.Close() // Penting: selalu close rows untuk mencegah memory leak
    
    var users []*models.User
    
    for rows.Next() {
        user := &models.User{}
        err := rows.Scan(
            &user.ID,
            &user.Name,
            &user.Email,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to scan user row: %w", err)
        }
        users = append(users, user)
    }
    
    // Check untuk error yang terjadi selama iterasi
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("error occurred during row iteration: %w", err)
    }
    
    return users, nil
}

func (r *mysqlUserRepository) Update(ctx context.Context, user *models.User) error {
    query := `
        UPDATE users 
        SET name = ?, email = ?, updated_at = ?
        WHERE id = ?
    `
    
    user.UpdatedAt = time.Now()
    
    result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.UpdatedAt, user.ID)
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }
    
    // Check apakah ada row yang ter-update
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("user with id %d not found", user.ID)
    }
    
    return nil
}

func (r *mysqlUserRepository) Delete(ctx context.Context, id int) error {
    query := `DELETE FROM users WHERE id = ?`
    
    result, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %w", err)
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("user with id %d not found", id)
    }
    
    return nil
}

func (r *mysqlUserRepository) Count(ctx context.Context) (int, error) {
    query := `SELECT COUNT(*) FROM users`
    
    var count int
    err := r.db.QueryRowContext(ctx, query).Scan(&count)
    if err != nil {
        return 0, fmt.Errorf("failed to count users: %w", err)
    }
    
    return count, nil
}
```

**Penjelasan Implementasi MySQL:**

1. **Struct Private**: `mysqlUserRepository` dengan huruf kecil (private) karena hanya akan diakses melalui interface
2. **Constructor Pattern**: `NewMySQLUserRepository` untuk membuat instance
3. **Context**: Semua method menggunakan context untuk timeout dan cancellation
4. **Error Handling**: Consistent error handling dengan wrapping error
5. **Resource Management**: `defer rows.Close()` untuk mencegah memory leak
6. **Validation**: Check `RowsAffected` untuk memastikan operasi berhasil

### Step 5: Service Layer (Business Logic)

```go
// services/user_service.go
package services

import (
    "context"
    "fmt"
    "regexp"
    "your-project/models"
    "your-project/repositories"
)

type UserService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

// CreateUser membuat user baru dengan validasi
func (s *UserService) CreateUser(ctx context.Context, req *models.UserCreateRequest) (*models.User, error) {
    // Validasi input
    if err := s.validateCreateUserRequest(req); err != nil {
        return nil, err
    }
    
    // Check apakah email sudah ada
    existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
    if err == nil && existingUser != nil {
        return nil, fmt.Errorf("user with email %s already exists", req.Email)
    }
    
    // Buat user baru
    user := &models.User{
        Name:  req.Name,
        Email: req.Email,
    }
    
    err = s.userRepo.Create(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    
    return user, nil
}

// GetUser mengambil user berdasarkan ID
func (s *UserService) GetUser(ctx context.Context, id int) (*models.User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user id: %d", id)
    }
    
    user, err := s.userRepo.GetByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return user, nil
}

// GetUsers mengambil semua user dengan pagination
func (s *UserService) GetUsers(ctx context.Context, page, pageSize int) ([]*models.User, int, error) {
    // Validasi pagination
    if page <= 0 {
        page = 1
    }
    if pageSize <= 0 || pageSize > 100 {
        pageSize = 10
    }
    
    offset := (page - 1) * pageSize
    
    // Ambil data user
    users, err := s.userRepo.GetAll(ctx, pageSize, offset)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to get users: %w", err)
    }
    
    // Hitung total untuk pagination
    total, err := s.userRepo.Count(ctx)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count users: %w", err)
    }
    
    return users, total, nil
}

// UpdateUser mengupdate data user
func (s *UserService) UpdateUser(ctx context.Context, id int, req *models.UserUpdateRequest) (*models.User, error) {
    // Validasi input
    if err := s.validateUpdateUserRequest(req); err != nil {
        return nil, err
    }
    
    // Ambil user yang ada
    user, err := s.userRepo.GetByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }
    
    // Check apakah email baru sudah digunakan user lain
    if req.Email != user.Email {
        existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
        if err == nil && existingUser != nil && existingUser.ID != id {
            return nil, fmt.Errorf("email %s is already used by another user", req.Email)
        }
    }
    
    // Update data
    user.Name = req.Name
    user.Email = req.Email
    
    err = s.userRepo.Update(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to update user: %w", err)
    }
    
    return user, nil
}

// DeleteUser menghapus user
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
    if id <= 0 {
        return fmt.Errorf("invalid user id: %d", id)
    }
    
    // Check apakah user ada
    _, err := s.userRepo.GetByID(ctx, id)
    if err != nil {
        return fmt.Errorf("user not found: %w", err)
    }
    
    err = s.userRepo.Delete(ctx, id)
    if err != nil {
        return fmt.Errorf("failed to delete user: %w", err)
    }
    
    return nil
}

// validateCreateUserRequest melakukan validasi untuk create user request
func (s *UserService) validateCreateUserRequest(req *models.UserCreateRequest) error {
    if req.Name == "" {
        return fmt.Errorf("name is required")
    }
    
    if len(req.Name) < 2 {
        return fmt.Errorf("name must be at least 2 characters")
    }
    
    if req.Email == "" {
        return fmt.Errorf("email is required")
    }
    
    if !s.isValidEmail(req.Email) {
        return fmt.Errorf("invalid email format")
    }
    
    return nil
}

// validateUpdateUserRequest melakukan validasi untuk update user request
func (s *UserService) validateUpdateUserRequest(req *models.UserUpdateRequest) error {
    // Sama dengan create validation
    return s.validateCreateUserRequest(&models.UserCreateRequest{
        Name:  req.Name,
        Email: req.Email,
    })
}

// isValidEmail melakukan validasi format email
func (s *UserService) isValidEmail(email string) bool {
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return emailRegex.MatchString(email)
}
```

**Penjelasan Service Layer:**

1. **Business Logic**: Semua validasi dan business rules ada di service
2. **Input Validation**: Validasi format email, required fields, dll
3. **Business Rules**: Check email duplicate, pagination logic
4. **Error Handling**: Consistent error wrapping
5. **Separation of Concerns**: Service tidak tahu tentang database implementation

### Step 6: HTTP Handler

```go
// handlers/user_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    
    "github.com/gorilla/mux"
    "your-project/models"
    "your-project/services"
)

type UserHandler struct {
    userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

// Response structs untuk API
type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

type PaginatedResponse struct {
    Response
    Page      int `json:"page"`
    PageSize  int `json:"page_size"`
    Total     int `json:"total"`
    TotalPage int `json:"total_page"`
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req models.UserCreateRequest
    
    // Parse JSON request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
        return
    }
    
    // Call service
    user, err := h.userService.CreateUser(r.Context(), &req)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
        return
    }
    
    // Send success response
    h.sendSuccessResponse(w, http.StatusCreated, "User created successfully", user)
}

// GetUser handles GET /users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
        return
    }
    
    user, err := h.userService.GetUser(r.Context(), id)
    if err != nil {
        h.sendErrorResponse(w, http.StatusNotFound, err.Error())
        return
    }
    
    h.sendSuccessResponse(w, http.StatusOK, "User retrieved successfully", user)
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    // Parse query parameters
    pageStr := r.URL.Query().Get("page")
    pageSizeStr := r.URL.Query().Get("page_size")
    
    page, _ := strconv.Atoi(pageStr)
    pageSize, _ := strconv.Atoi(pageSizeStr)
    
    users, total, err := h.userService.GetUsers(r.Context(), page, pageSize)
    if err != nil {
        h.sendErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }
    
    // Calculate pagination
    if page <= 0 {
        page = 1
    }
    if pageSize <= 0 {
        pageSize = 10
    }
    totalPage := (total + pageSize - 1) / pageSize
    
    response := PaginatedResponse{
        Response: Response{
            Success: true,
            Message: "Users retrieved successfully",
            Data:    users,
        },
        Page:      page,
        PageSize:  pageSize,
        Total:     total,
        TotalPage: totalPage,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// UpdateUser handles PUT /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
        return
    }
    
    var req models.UserUpdateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
        return
    }
    
    user, err := h.userService.UpdateUser(r.Context(), id, &req)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
        return
    }
    
    h.sendSuccessResponse(w, http.StatusOK, "User updated successfully", user)
}

// DeleteUser handles DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
        return
    }
    
    err = h.userService.DeleteUser(r.Context(), id)
    if err != nil {
        h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
        return
    }
    
    h.sendSuccessResponse(w, http.StatusOK, "User deleted successfully", nil)
}

// Helper methods untuk response
func (h *UserHandler) sendSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    
    response := Response{
        Success: true,
        Message: message,
        Data:    data,
    }
    
    json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    
    response := Response{
        Success: false,
        Message: message,
    }
    
    json.NewEncoder(w).Encode(response)
}
```

### Step 7: Database Setup dan Main Application

```go
// main.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    
    "your-project/handlers"
    "your-project/repositories"
    "your-project/services"
)

func main() {
    // Database connection
    db, err := connectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()
    
    // Create tables if not exists
    if err := createTables(db); err != nil {
        log.Fatal("Failed to create tables:", err)
    }
    
    // Initialize repository
    userRepo := repositories.NewMySQLUserRepository(db)
    
    // Initialize service
    userService := services.NewUserService(userRepo)
    
    // Initialize handler
    userHandler := handlers.NewUserHandler(userService)
    
    // Setup routes
    router := mux.NewRouter()
    
    // User routes
    router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
    router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
    router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
    
    // Start server
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func connectDB() (*sql.DB, error) {
    // Database configuration - dalam production, gunakan environment variables
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "3306")
    dbUser := getEnv("DB_USER", "root")
    dbPassword := getEnv("DB_PASSWORD", "")
    dbName := getEnv("DB_NAME", "test_db")
    
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
        dbUser, dbPassword, dbHost, dbPort, dbName)
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    // Test connection
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    // Connection pool settings
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    
    return db, nil
}

func createTables(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    )
    `
    
    _, err := db.Exec(query)
    return err
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

## 4. Best Practices & Tips

### Database Best Practices

1. **Connection Pool**: Selalu set `MaxOpenConns` dan `MaxIdleConns`
2. **Prepared Statements**: Gunakan parameter placeholders (`?`) untuk mencegah SQL injection
3. **Context**: Selalu gunakan context untuk timeout dan cancellation
4. **Resource Management**: Selalu close rows, statements, dan connections
5. **Transaction**: Gunakan transaction untuk operasi yang kompleks

### Repository Best Practices

1. **Interface First**: Selalu define interface sebelum implementasi
2. **Single Responsibility**: Satu repository untuk satu entity
3. **Error Wrapping**: Wrap error dengan context yang jelas
4. **Consistent Naming**: Gunakan naming convention yang konsisten
5. **No Business Logic**: Repository hanya untuk data access, no business logic

### Service Best Practices

1. **Input Validation**: Selalu validate input di service layer
2. **Business Rules**: Semua business logic ada di service
3. **Error Handling**: Handle dan wrap error dengan baik
4. **Testable**: Struktur service agar mudah di-test
5. **Dependency Injection**: Gunakan dependency injection pattern

### Security Best Practices

1. **SQL Injection**: Gunakan prepared statements
2. **Input Validation**: Validate semua input dari user
3. **Error Messages**: Jangan expose sensitive information di error message
4. **Rate Limiting**: Implement rate limiting untuk API
5. **Authentication**: Implement proper authentication dan authorization

## 5. Testing Example

```go
// repositories/user_repository_test.go
package repositories_test

import (
    "context"
    "testing"
    "your-project/models"
    "your-project/repositories"
)

// MockUserRepository adalah mock implementation untuk testing
type MockUserRepository struct {
    users []models.User
    nextID int
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users:  make([]models.User, 0),
        nextID: 1,
    }
}

func (m *MockUserRepository) Create(ctx context.Context, user *models.User) error {
    user.ID = m.nextID
    m.nextID++
    m.users = append(m.users, *user)
    return nil
}

func (m *MockUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    for _, user := range m.users {
        if user.ID == id {
            return &user, nil
        }
    }
    return nil, fmt.Errorf("user not found")
}

// Test service dengan mock repository
func TestUserService_CreateUser(t *testing.T) {
    // Setup
    mockRepo := NewMockUserRepository()
    userService := services.NewUserService(mockRepo)
    
    req := &models.UserCreateRequest{
        Name:  "John Doe",
        Email: "john@example.com",
    }
    
    // Test
    user, err := userService.CreateUser(context.Background(), req)
    
    // Assert
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if user.Name != req.Name {
        t.Errorf("Expected name %s, got %s", req.Name, user.Name)
    }
    
    if user.ID == 0 {
        t.Error("Expected user ID to be set")
    }
}
```

## 6. Keuntungan Repository Pattern

### Untuk Development
- **Maintainability**: Kode lebih terorganisir dan mudah dipelihara
- **Testability**: Mudah membuat unit test dengan mock repository
- **Scalability**: Mudah menambah fitur baru tanpa mengubah existing code
- **Team Collaboration**: Tim bisa bekerja parallel pada layer yang berbeda

### Untuk Business
- **Flexibility**: Mudah ganti database (MySQL ke PostgreSQL, MongoDB, etc.)
- **Performance**: Bisa optimize query di repository tanpa ubah business logic
- **Caching**: Mudah implement caching layer di repository
- **Monitoring**: Mudah add logging dan monitoring di satu tempat

## 7. Contoh Penggunaan Advanced

### Transaction Management

```go
// repositories/transaction.go
package repositories

import (
    "context"
    "database/sql"
    "fmt"
)

// TxManager handles database transactions
type TxManager struct {
    db *sql.DB
}

func NewTxManager(db *sql.DB) *TxManager {
    return &TxManager{db: db}
}

// WithTransaction executes function inside a transaction
func (tm *TxManager) WithTransaction(ctx context.Context, fn func(*sql.Tx) error) error {
    tx, err := tm.db.BeginTx(ctx, nil)
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }
    
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p) // Re-throw panic after rollback
        }
    }()
    
    if err := fn(tx); err != nil {
        if rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
        }
        return err
    }
    
    return tx.Commit()
}

// Contoh penggunaan transaction di service
func (s *UserService) CreateUserWithProfile(ctx context.Context, userReq *models.UserCreateRequest, profileReq *models.ProfileCreateRequest) error {
    return s.txManager.WithTransaction(ctx, func(tx *sql.Tx) error {
        // Create user
        user := &models.User{
            Name:  userReq.Name,
            Email: userReq.Email,
        }
        
        if err := s.userRepo.CreateWithTx(ctx, tx, user); err != nil {
            return err
        }
        
        // Create profile
        profile := &models.Profile{
            UserID: user.ID,
            Bio:    profileReq.Bio,
        }
        
        if err := s.profileRepo.CreateWithTx(ctx, tx, profile); err != nil {
            return err
        }
        
        return nil
    })
}
```

### Caching Layer

```go
// repositories/cached_user_repository.go
package repositories

import (
    "context"
    "fmt"
    "time"
    "your-project/models"
)

// CachedUserRepository wraps UserRepository dengan caching
type CachedUserRepository struct {
    repo  UserRepository
    cache map[string]*models.User // Dalam production, gunakan Redis
    ttl   time.Duration
}

func NewCachedUserRepository(repo UserRepository) UserRepository {
    return &CachedUserRepository{
        repo:  repo,
        cache: make(map[string]*models.User),
        ttl:   5 * time.Minute,
    }
}

func (c *CachedUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    cacheKey := fmt.Sprintf("user:%d", id)
    
    // Check cache first
    if user, exists := c.cache[cacheKey]; exists {
        return user, nil
    }
    
    // If not in cache, get from database
    user, err := c.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    c.cache[cacheKey] = user
    
    // Setup TTL (dalam production, gunakan proper cache dengan TTL)
    go func() {
        time.Sleep(c.ttl)
        delete(c.cache, cacheKey)
    }()
    
    return user, nil
}

func (c *CachedUserRepository) Create(ctx context.Context, user *models.User) error {
    err := c.repo.Create(ctx, user)
    if err != nil {
        return err
    }
    
    // Invalidate related cache
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    delete(c.cache, cacheKey)
    
    return nil
}

// Implement other methods...
```

### Repository dengan Filtering dan Sorting

```go
// models/filter.go
package models

type UserFilter struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    SortBy   string `json:"sort_by"`   // "name", "email", "created_at"
    SortDir  string `json:"sort_dir"`  // "asc", "desc"
    Page     int    `json:"page"`
    PageSize int    `json:"page_size"`
}

// repositories/mysql_user_repository.go - tambahkan method ini
func (r *mysqlUserRepository) GetWithFilter(ctx context.Context, filter *models.UserFilter) ([]*models.User, int, error) {
    // Build WHERE clause
    var conditions []string
    var args []interface{}
    
    if filter.Name != "" {
        conditions = append(conditions, "name LIKE ?")
        args = append(args, "%"+filter.Name+"%")
    }
    
    if filter.Email != "" {
        conditions = append(conditions, "email LIKE ?")
        args = append(args, "%"+filter.Email+"%")
    }
    
    whereClause := ""
    if len(conditions) > 0 {
        whereClause = "WHERE " + strings.Join(conditions, " AND ")
    }
    
    // Build ORDER BY clause
    orderClause := "ORDER BY created_at DESC" // default
    if filter.SortBy != "" {
        allowedSorts := map[string]bool{
            "name":       true,
            "email":      true,
            "created_at": true,
        }
        
        if allowedSorts[filter.SortBy] {
            direction := "DESC"
            if filter.SortDir == "asc" {
                direction = "ASC"
            }
            orderClause = fmt.Sprintf("ORDER BY %s %s", filter.SortBy, direction)
        }
    }
    
    // Pagination
    if filter.Page <= 0 {
        filter.Page = 1
    }
    if filter.PageSize <= 0 || filter.PageSize > 100 {
        filter.PageSize = 10
    }
    
    offset := (filter.Page - 1) * filter.PageSize
    
    // Query for data
    query := fmt.Sprintf(`
        SELECT id, name, email, created_at, updated_at 
        FROM users 
        %s 
        %s 
        LIMIT ? OFFSET ?
    `, whereClause, orderClause)
    
    args = append(args, filter.PageSize, offset)
    
    rows, err := r.db.QueryContext(ctx, query, args...)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to get users with filter: %w", err)
    }
    defer rows.Close()
    
    var users []*models.User
    for rows.Next() {
        user := &models.User{}
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return nil, 0, fmt.Errorf("failed to scan user: %w", err)
        }
        users = append(users, user)
    }
    
    // Count total
    countQuery := fmt.Sprintf("SELECT COUNT(*) FROM users %s", whereClause)
    var total int
    err = r.db.QueryRowContext(ctx, countQuery, args[:len(args)-2]...).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to count users: %w", err)
    }
    
    return users, total, nil
}
```

## 8. Error Handling Pattern

### Custom Error Types

```go
// errors/errors.go
package errors

import "fmt"

// Domain errors
type DomainError struct {
    Type    string
    Message string
    Err     error
}

func (e *DomainError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %s: %v", e.Type, e.Message, e.Err)
    }
    return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *DomainError) Unwrap() error {
    return e.Err
}

// Predefined error types
var (
    ErrNotFound     = &DomainError{Type: "NOT_FOUND", Message: "resource not found"}
    ErrDuplicate    = &DomainError{Type: "DUPLICATE", Message: "resource already exists"}
    ErrValidation   = &DomainError{Type: "VALIDATION", Message: "validation error"}
    ErrUnauthorized = &DomainError{Type: "UNAUTHORIZED", Message: "unauthorized"}
    ErrInternal     = &DomainError{Type: "INTERNAL", Message: "internal server error"}
)

// Helper functions
func NotFound(message string, err error) error {
    return &DomainError{Type: "NOT_FOUND", Message: message, Err: err}
}

func Duplicate(message string, err error) error {
    return &DomainError{Type: "DUPLICATE", Message: message, Err: err}
}

func Validation(message string, err error) error {
    return &DomainError{Type: "VALIDATION", Message: message, Err: err}
}

// Usage di repository
func (r *mysqlUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?"
    
    user := &models.User{}
    err := r.db.QueryRowContext(ctx, query, id).Scan(
        &user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.NotFound(fmt.Sprintf("user with id %d not found", id), err)
        }
        return nil, errors.Internal("failed to get user", err)
    }
    
    return user, nil
}

// Usage di handler
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    // ... parse ID ...
    
    user, err := h.userService.GetUser(r.Context(), id)
    if err != nil {
        var domainErr *errors.DomainError
        if errors.As(err, &domainErr) {
            switch domainErr.Type {
            case "NOT_FOUND":
                h.sendErrorResponse(w, http.StatusNotFound, domainErr.Message)
                return
            case "VALIDATION":
                h.sendErrorResponse(w, http.StatusBadRequest, domainErr.Message)
                return
            default:
                h.sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
                return
            }
        }
        
        h.sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
        return
    }
    
    h.sendSuccessResponse(w, http.StatusOK, "User retrieved successfully", user)
}
```

## 9. Monitoring dan Logging

```go
// middleware/logging.go
package middleware

import (
    "context"
    "log"
    "time"
    "your-project/repositories"
    "your-project/models"
)

// LoggingUserRepository wraps repository dengan logging
type LoggingUserRepository struct {
    repo repositories.UserRepository
}

func NewLoggingUserRepository(repo repositories.UserRepository) repositories.UserRepository {
    return &LoggingUserRepository{repo: repo}
}

func (l *LoggingUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    start := time.Now()
    
    user, err := l.repo.GetByID(ctx, id)
    
    duration := time.Since(start)
    
    if err != nil {
        log.Printf("UserRepository.GetByID failed: id=%d, duration=%v, error=%v", 
            id, duration, err)
    } else {
        log.Printf("UserRepository.GetByID success: id=%d, duration=%v", 
            id, duration)
    }
    
    return user, err
}

// Implement other methods dengan logging yang sama...
```

## 10. Performance Tips

### Connection Pooling

```go
// config/database.go
package config

import (
    "database/sql"
    "time"
)

func SetupDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    // Connection pool settings
    db.SetMaxOpenConns(25)        // Maximum number of open connections
    db.SetMaxIdleConns(25)        // Maximum number of idle connections
    db.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime
    db.SetConnMaxIdleTime(5 * time.Minute) // Maximum idle time
    
    return db, nil
}
```

### Prepared Statements untuk Query yang Sering Dipanggil

```go
// repositories/prepared_user_repository.go
package repositories

import (
    "context"
    "database/sql"
    "your-project/models"
)

type PreparedUserRepository struct {
    db *sql.DB
    
    // Prepared statements
    getByIDStmt    *sql.Stmt
    getByEmailStmt *sql.Stmt
    createStmt     *sql.Stmt
    updateStmt     *sql.Stmt
    deleteStmt     *sql.Stmt
}

func NewPreparedUserRepository(db *sql.DB) (repositories.UserRepository, error) {
    repo := &PreparedUserRepository{db: db}
    
    if err := repo.prepareStatements(); err != nil {
        return nil, err
    }
    
    return repo, nil
}

func (r *PreparedUserRepository) prepareStatements() error {
    var err error
    
    // Prepare statements
    r.getByIDStmt, err = r.db.Prepare("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?")
    if err != nil {
        return err
    }
    
    r.getByEmailStmt, err = r.db.Prepare("SELECT id, name, email, created_at, updated_at FROM users WHERE email = ?")
    if err != nil {
        return err
    }
    
    r.createStmt, err = r.db.Prepare("INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, ?, ?)")
    if err != nil {
        return err
    }
    
    r.updateStmt, err = r.db.Prepare("UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?")
    if err != nil {
        return err
    }
    
    r.deleteStmt, err = r.db.Prepare("DELETE FROM users WHERE id = ?")
    if err != nil {
        return err
    }
    
    return nil
}

func (r *PreparedUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    user := &models.User{}
    err := r.getByIDStmt.QueryRowContext(ctx, id).Scan(
        &user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with id %d not found", id)
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return user, nil
}

// Close method untuk cleanup prepared statements
func (r *PreparedUserRepository) Close() error {
    var errs []error
    
    if r.getByIDStmt != nil {
        if err := r.getByIDStmt.Close(); err != nil {
            errs = append(errs, err)
        }
    }
    
    // Close other statements...
    
    if len(errs) > 0 {
        return fmt.Errorf("failed to close prepared statements: %v", errs)
    }
    
    return nil
}
```

## 11. Testing Strategy

### Integration Testing

```go
// tests/integration_test.go
package tests

import (
    "context"
    "database/sql"
    "testing"
    
    _ "github.com/go-sql-driver/mysql"
    "your-project/repositories"
    "your-project/models"
)

func setupTestDB(t *testing.T) *sql.DB {
    // Gunakan test database
    dsn := "root:@tcp(localhost:3306)/test_db?parseTime=true"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        t.Fatal("Failed to connect to test database:", err)
    }
    
    // Clean up test data
    _, err = db.Exec("DELETE FROM users")
    if err != nil {
        t.Fatal("Failed to clean test data:", err)
    }
    
    return db
}

func TestUserRepository_Integration(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    repo := repositories.NewMySQLUserRepository(db)
    ctx := context.Background()
    
    // Test Create
    user := &models.User{
        Name:  "Test User",
        Email: "test@example.com",
    }
    
    err := repo.Create(ctx, user)
    if err != nil {
        t.Errorf("Failed to create user: %v", err)
    }
    
    if user.ID == 0 {
        t.Error("Expected user ID to be set after creation")
    }
    
    // Test GetByID
    retrievedUser, err := repo.GetByID(ctx, user.ID)
    if err != nil {
        t.Errorf("Failed to get user: %v", err)
    }
    
    if retrievedUser.Name != user.Name {
        t.Errorf("Expected name %s, got %s", user.Name, retrievedUser.Name)
    }
    
    // Test Update
    retrievedUser.Name = "Updated Name"
    err = repo.Update(ctx, retrievedUser)
    if err != nil {
        t.Errorf("Failed to update user: %v", err)
    }
    
    // Verify update
    updatedUser, err := repo.GetByID(ctx, user.ID)
    if err != nil {
        t.Errorf("Failed to get updated user: %v", err)
    }
    
    if updatedUser.Name != "Updated Name" {
        t.Errorf("Expected name 'Updated Name', got %s", updatedUser.Name)
    }
    
    // Test Delete
    err = repo.Delete(ctx, user.ID)
    if err != nil {
        t.Errorf("Failed to delete user: %v", err)
    }
    
    // Verify deletion
    _, err = repo.GetByID(ctx, user.ID)
    if err == nil {
        t.Error("Expected error when getting deleted user")
    }
}
```

## 12. Deployment Considerations

### Environment Configuration

```go
// config/config.go
package config

import (
    "fmt"
    "os"
    "strconv"
)

type DatabaseConfig struct {
    Host         string
    Port         int
    User         string
    Password     string
    Name         string
    MaxOpenConns int
    MaxIdleConns int
}

func LoadDatabaseConfig() (*DatabaseConfig, error) {
    port, err := strconv.Atoi(getEnv("DB_PORT", "3306"))
    if err != nil {
        return nil, fmt.Errorf("invalid DB_PORT: %w", err)
    }
    
    maxOpenConns, err := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "25"))
    if err != nil {
        return nil, fmt.Errorf("invalid DB_MAX_OPEN_CONNS: %w", err)
    }
    
    maxIdleConns, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "25"))
    if err != nil {
        return nil, fmt.Errorf("invalid DB_MAX_IDLE_CONNS: %w", err)
    }
    
    return &DatabaseConfig{
        Host:         getEnv("DB_HOST", "localhost"),
        Port:         port,
        User:         getEnv("DB_USER", "root"),
        Password:     getEnv("DB_PASSWORD", ""),
        Name:         getEnv("DB_NAME", "myapp"),
        MaxOpenConns: maxOpenConns,
        MaxIdleConns: maxIdleConns,
    }, nil
}

func (c *DatabaseConfig) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
        c.User, c.Password, c.Host, c.Port, c.Name)
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

### Docker Compose untuk Development

```yaml
# docker-compose.yml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: myapp
      MYSQL_USER: appuser
      MYSQL_PASSWORD: apppassword
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
      - ./scripts/init.sql:/docker-entrypoint-initdb.d/init.sql

  app:
    build: .
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: appuser
      DB_PASSWORD: apppassword
      DB_NAME: myapp
    ports:
      - "8080:8080"
    depends_on:
      - mysql

volumes:
  mysql_data:
```

## Kesimpulan

Repository Pattern adalah salah satu pattern paling penting dalam pengembangan aplikasi Go yang scalable dan maintainable. Key takeaways:

1. **Separation of Concerns**: Pisahkan data access dari business logic
2. **Interface-Driven Design**: Selalu mulai dengan interface
3. **Testability**: Mock repository memudahkan unit testing
4. **Flexibility**: Mudah ganti database atau add caching
5. **Best Practices**: Gunakan context, handle error dengan baik, manage resources

Dengan mengimplementasikan Repository Pattern dengan benar, Anda akan memiliki codebase yang:
- Mudah di-maintain dan di-extend
- Mudah di-test dengan unit testing
- Flexible untuk perubahan requirement
- Performant dengan proper connection pooling
- Secure dengan prepared statements

Pattern ini sangat recommended untuk aplikasi production dan menjadi foundation yang solid untuk aplikasi yang akan grow dan scale.