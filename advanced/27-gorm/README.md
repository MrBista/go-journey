# Tutorial GORM Golang - Dari Pemula hingga Advanced


### PZN MATERI: https://docs.google.com/presentation/d/11NNdQj-3UkmkdQRPQC0CtzeDfLP-IsSiIfDbDmpfo_4/edit?usp=sharing

## 1. Apa itu GORM?

GORM adalah Object-Relational Mapping (ORM) library untuk Golang yang powerful dan developer-friendly. GORM memudahkan kita berinteraksi dengan database tanpa menulis SQL mentah.

### Keunggulan GORM:
- Full-Featured ORM
- Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism)
- Hooks (Before/After Create/Save/Update/Delete/Find)
- Preloading (Eager Loading)
- Transactions
- Composite Primary Key
- SQL Builder
- Auto Migrations
- Logger
- Write Extendable, write Plugins based on GORM callbacks

## 2. Setup dan Instalasi

### Install GORM dan Database Driver

```bash
# Install GORM
go mod init gorm-tutorial
go get -u gorm.io/gorm

# Install database driver (pilih sesuai database yang digunakan)
go get -u gorm.io/driver/mysql     # MySQL
go get -u gorm.io/driver/postgres  # PostgreSQL
go get -u gorm.io/driver/sqlite    # SQLite
go get -u gorm.io/driver/sqlserver # SQL Server
```

### Koneksi Database

```go
package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "gorm.io/driver/sqlite"
    "log"
)

// Koneksi MySQL
func connectMySQL() *gorm.DB {
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    return db
}

// Koneksi SQLite (untuk testing)
func connectSQLite() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    return db
}
```

## 3. Model Definition

### Basic Model

```go
package models

import (
    "time"
    "gorm.io/gorm"
)

// User model dengan embedded gorm.Model
type User struct {
    gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
    Name     string   `json:"name" gorm:"size:100;not null"`
    Email    string   `json:"email" gorm:"uniqueIndex;size:100;not null"`
    Age      int      `json:"age" gorm:"check:age > 0"`
    Active   bool     `json:"active" gorm:"default:true"`
}

// Custom Model tanpa gorm.Model
type Product struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"size:200;not null"`
    Price       float64   `json:"price" gorm:"type:decimal(10,2);not null"`
    CategoryID  uint      `json:"category_id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

// Category model
type Category struct {
    ID       uint      `json:"id" gorm:"primaryKey"`
    Name     string    `json:"name" gorm:"size:100;not null"`
    Products []Product `json:"products" gorm:"foreignKey:CategoryID"`
}
```

### Field Tags (Penting!)

```go
type User struct {
    ID       uint   `gorm:"primaryKey"`                    // Primary key
    Name     string `gorm:"size:100;not null"`            // VARCHAR(100) NOT NULL
    Email    string `gorm:"uniqueIndex;size:100"`         // Unique index
    Age      int    `gorm:"check:age > 0"`                // Check constraint
    Active   bool   `gorm:"default:true"`                 // Default value
    Bio      string `gorm:"type:text"`                    // Custom type
    Salary   int    `gorm:"column:user_salary"`           // Custom column name
    Ignored  string `gorm:"-"`                            // Ignore field
}
```

## 4. CRUD Operations

### Create (Insert)

```go
func createUser(db *gorm.DB) {
    // Single create
    user := User{
        Name:   "John Doe",
        Email:  "john@example.com",
        Age:    25,
        Active: true,
    }
    
    result := db.Create(&user)
    if result.Error != nil {
        log.Println("Error creating user:", result.Error)
        return
    }
    
    fmt.Printf("User created with ID: %d\n", user.ID)
    fmt.Printf("Rows affected: %d\n", result.RowsAffected)
}

// Batch create
func createUsers(db *gorm.DB) {
    users := []User{
        {Name: "Alice", Email: "alice@example.com", Age: 22},
        {Name: "Bob", Email: "bob@example.com", Age: 28},
        {Name: "Charlie", Email: "charlie@example.com", Age: 35},
    }
    
    // Batch insert dengan batch size
    db.CreateInBatches(users, 100) // Insert 100 records per batch
}
```

### Read (Select)

```go
func readUsers(db *gorm.DB) {
    var user User
    var users []User
    
    // Find by primary key
    db.First(&user, 1) // SELECT * FROM users WHERE id = 1 LIMIT 1;
    
    // Find by condition
    db.Where("name = ?", "John Doe").First(&user)
    
    // Find all
    db.Find(&users) // SELECT * FROM users;
    
    // Find with conditions
    db.Where("age > ?", 18).Find(&users)
    db.Where("name LIKE ?", "%John%").Find(&users)
    
    // Find with multiple conditions
    db.Where("age > ? AND active = ?", 18, true).Find(&users)
    
    // Using struct for conditions
    db.Where(&User{Age: 25, Active: true}).Find(&users)
    
    // Using map for conditions
    db.Where(map[string]interface{}{
        "age":    25,
        "active": true,
    }).Find(&users)
    
    // Order and Limit
    db.Where("age > ?", 18).Order("age desc").Limit(10).Offset(5).Find(&users)
    
    // Select specific fields
    db.Select("name", "email").Where("age > ?", 18).Find(&users)
    
    // Count
    var count int64
    db.Model(&User{}).Where("age > ?", 18).Count(&count)
}
```

### Update

```go
func updateUsers(db *gorm.DB) {
    var user User
    
    // Update single record
    db.First(&user, 1)
    user.Name = "John Updated"
    user.Age = 26
    db.Save(&user) // UPDATE users SET name='John Updated', age=26 WHERE id=1;
    
    // Update with Where
    db.Model(&user).Where("id = ?", 1).Update("name", "John Again")
    
    // Update multiple fields
    db.Model(&user).Where("id = ?", 1).Updates(User{
        Name: "John Multiple",
        Age:  27,
    })
    
    // Update with map
    db.Model(&user).Where("id = ?", 1).Updates(map[string]interface{}{
        "name": "John Map",
        "age":  28,
    })
    
    // Update without loading (more efficient)
    db.Where("age > ?", 30).Updates(User{Active: false})
}
```

### Delete

```go
func deleteUsers(db *gorm.DB) {
    var user User
    
    // Soft delete (if using gorm.Model)
    db.Delete(&user, 1) // UPDATE users SET deleted_at=NOW() WHERE id=1;
    
    // Delete with condition
    db.Where("age < ?", 18).Delete(&User{})
    
    // Permanent delete (hard delete)
    db.Unscoped().Delete(&user, 1) // DELETE FROM users WHERE id=1;
    
    // Delete all records (be careful!)
    // db.Where("1 = 1").Delete(&User{})
}
```

## 5. Associations (Relasi)

### One to Many (Has Many / Belongs To)

```go
type User struct {
    gorm.Model
    Name    string
    Email   string
    Posts   []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
    gorm.Model
    Title   string
    Content string
    UserID  uint
    User    User `gorm:"references:ID"`
}

// Usage
func associationsExample(db *gorm.DB) {
    // Create user with posts
    user := User{
        Name:  "John",
        Email: "john@example.com",
        Posts: []Post{
            {Title: "Post 1", Content: "Content 1"},
            {Title: "Post 2", Content: "Content 2"},
        },
    }
    db.Create(&user)
    
    // Preload associations
    var users []User
    db.Preload("Posts").Find(&users)
    
    // Create post for existing user
    post := Post{
        Title:   "New Post",
        Content: "New Content",
        UserID:  1,
    }
    db.Create(&post)
}
```

### Many to Many

```go
type User struct {
    gorm.Model
    Name  string
    Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
    gorm.Model
    Name  string
    Users []User `gorm:"many2many:user_roles;"`
}

// Usage
func manyToManyExample(db *gorm.DB) {
    // Create user with roles
    user := User{
        Name: "John",
        Roles: []Role{
            {Name: "Admin"},
            {Name: "Editor"},
        },
    }
    db.Create(&user)
    
    // Add role to existing user
    var existingUser User
    var adminRole Role
    db.First(&existingUser, 1)
    db.Where("name = ?", "Admin").First(&adminRole)
    
    db.Model(&existingUser).Association("Roles").Append(&adminRole)
}
```

## 6. Advanced Features

### Transactions

```go
func transactionExample(db *gorm.DB) {
    // Manual transaction
    tx := db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()
    
    if err := tx.Error; err != nil {
        return
    }
    
    // Create user
    user := User{Name: "John", Email: "john@example.com"}
    if err := tx.Create(&user).Error; err != nil {
        tx.Rollback()
        return
    }
    
    // Create post
    post := Post{Title: "Post", Content: "Content", UserID: user.ID}
    if err := tx.Create(&post).Error; err != nil {
        tx.Rollback()
        return
    }
    
    tx.Commit()
}

// Transaction with closure (recommended)
func transactionWithClosure(db *gorm.DB) {
    err := db.Transaction(func(tx *gorm.DB) error {
        // Create user
        user := User{Name: "John", Email: "john@example.com"}
        if err := tx.Create(&user).Error; err != nil {
            return err
        }
        
        // Create post
        post := Post{Title: "Post", Content: "Content", UserID: user.ID}
        if err := tx.Create(&post).Error; err != nil {
            return err
        }
        
        return nil
    })
    
    if err != nil {
        log.Println("Transaction failed:", err)
    }
}
```

### Hooks

```go
// Hooks pada model
func (u *User) BeforeCreate(tx *gorm.DB) error {
    // Hash password, generate UUID, etc.
    log.Println("Before creating user:", u.Name)
    return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
    // Send welcome email, log activity, etc.
    log.Println("After creating user:", u.Name)
    return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
    log.Println("Before updating user:", u.Name)
    return nil
}

func (u *User) AfterFind(tx *gorm.DB) error {
    // Log access, decrypt sensitive data, etc.
    log.Println("After finding user:", u.Name)
    return nil
}
```

### Scopes

```go
// Define reusable scopes
func ActiveUsers(db *gorm.DB) *gorm.DB {
    return db.Where("active = ?", true)
}

func AdultUsers(db *gorm.DB) *gorm.DB {
    return db.Where("age >= ?", 18)
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        offset := (page - 1) * pageSize
        return db.Offset(offset).Limit(pageSize)
    }
}

// Usage
func scopeExample(db *gorm.DB) {
    var users []User
    
    // Using scopes
    db.Scopes(ActiveUsers, AdultUsers).Find(&users)
    
    // With pagination
    db.Scopes(ActiveUsers, Paginate(1, 10)).Find(&users)
}
```

## 7. Best Practices

### 1. Model Organization

```go
// models/base.go
type BaseModel struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// models/user.go
type User struct {
    BaseModel
    Name     string `json:"name" gorm:"size:100;not null"`
    Email    string `json:"email" gorm:"uniqueIndex;size:100;not null"`
    Password string `json:"-" gorm:"size:255;not null"` // Hidden in JSON
}
```

### 2. Repository Pattern

```go
// repositories/user_repository.go
type UserRepository interface {
    Create(user *User) error
    GetByID(id uint) (*User, error)
    GetByEmail(email string) (*User, error)
    Update(user *User) error
    Delete(id uint) error
    List(page, limit int) ([]User, int64, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*User, error) {
    var user User
    err := r.db.First(&user, id).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*User, error) {
    var user User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) List(page, limit int) ([]User, int64, error) {
    var users []User
    var total int64
    
    offset := (page - 1) * limit
    
    // Count total
    r.db.Model(&User{}).Count(&total)
    
    // Get paginated results
    err := r.db.Offset(offset).Limit(limit).Find(&users).Error
    
    return users, total, err
}
```

### 3. Error Handling

```go
import (
    "errors"
    "gorm.io/gorm"
)

func handleGormErrors(err error) error {
    if err == nil {
        return nil
    }
    
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return errors.New("record not found")
    }
    
    if errors.Is(err, gorm.ErrDuplicatedKey) {
        return errors.New("duplicate key violation")
    }
    
    return err
}

func getUserSafely(db *gorm.DB, id uint) (*User, error) {
    var user User
    err := db.First(&user, id).Error
    if err != nil {
        return nil, handleGormErrors(err)
    }
    return &user, nil
}
```

### 4. Migration Management

```go
// migrations/migrate.go
func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &User{},
        &Post{},
        &Category{},
        &Product{},
        // Add all your models here
    )
}

// Manual migration example
func CreateUserTable(db *gorm.DB) error {
    return db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id BIGINT PRIMARY KEY AUTO_INCREMENT,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            age INT CHECK(age > 0),
            active BOOLEAN DEFAULT TRUE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        )
    `).Error
}
```

### 5. Configuration

```go
// config/database.go
type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

func NewDB(config DBConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User,
        config.Password,
        config.Host,
        config.Port,
        config.DBName,
    )
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Enable logging
        NowFunc: func() time.Time {
            return time.Now().Local()
        },
        DryRun: false, // Set to true for SQL preview
    })
    
    if err != nil {
        return nil, err
    }
    
    // Connection pool settings
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    
    sqlDB.SetMaxIdleConns(10)           // Maximum idle connections
    sqlDB.SetMaxOpenConns(100)          // Maximum open connections
    sqlDB.SetConnMaxLifetime(time.Hour) // Maximum connection lifetime
    
    return db, nil
}
```

## 8. Use Cases dan Contoh Praktis

### E-Commerce Example

```go
// models/ecommerce.go
type Customer struct {
    gorm.Model
    Name    string  `gorm:"size:100;not null"`
    Email   string  `gorm:"uniqueIndex;size:100;not null"`
    Phone   string  `gorm:"size:20"`
    Orders  []Order `gorm:"foreignKey:CustomerID"`
}

type Product struct {
    gorm.Model
    Name        string      `gorm:"size:200;not null"`
    Description string      `gorm:"type:text"`
    Price       float64     `gorm:"type:decimal(10,2);not null"`
    Stock       int         `gorm:"not null;check:stock >= 0"`
    CategoryID  uint
    Category    Category    `gorm:"references:ID"`
    OrderItems  []OrderItem `gorm:"foreignKey:ProductID"`
}

type Category struct {
    gorm.Model
    Name     string    `gorm:"size:100;not null"`
    Products []Product `gorm:"foreignKey:CategoryID"`
}

type Order struct {
    gorm.Model
    CustomerID  uint
    Customer    Customer    `gorm:"references:ID"`
    Status      string      `gorm:"size:50;not null;default:'pending'"`
    Total       float64     `gorm:"type:decimal(10,2);not null"`
    OrderItems  []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
    gorm.Model
    OrderID   uint
    Order     Order   `gorm:"references:ID"`
    ProductID uint
    Product   Product `gorm:"references:ID"`
    Quantity  int     `gorm:"not null;check:quantity > 0"`
    Price     float64 `gorm:"type:decimal(10,2);not null"`
}

// Service layer example
type OrderService struct {
    db *gorm.DB
}

func (s *OrderService) CreateOrder(customerID uint, items []OrderItem) (*Order, error) {
    var total float64
    
    // Start transaction
    return nil, s.db.Transaction(func(tx *gorm.DB) error {
        // Validate products and calculate total
        for i, item := range items {
            var product Product
            if err := tx.First(&product, item.ProductID).Error; err != nil {
                return err
            }
            
            // Check stock
            if product.Stock < item.Quantity {
                return fmt.Errorf("insufficient stock for product %s", product.Name)
            }
            
            // Set price and calculate total
            items[i].Price = product.Price
            total += product.Price * float64(item.Quantity)
            
            // Update stock
            product.Stock -= item.Quantity
            if err := tx.Save(&product).Error; err != nil {
                return err
            }
        }
        
        // Create order
        order := Order{
            CustomerID: customerID,
            Status:     "pending",
            Total:      total,
            OrderItems: items,
        }
        
        return tx.Create(&order).Error
    })
}
```

## 9. Performance Tips

### 1. Preloading
```go
// Bad: N+1 query problem
var users []User
db.Find(&users)
for _, user := range users {
    var posts []Post
    db.Where("user_id = ?", user.ID).Find(&posts) // This runs for each user!
}

// Good: Use Preload
var users []User
db.Preload("Posts").Find(&users) // Single query with JOIN
```

### 2. Select Specific Fields
```go
// Bad: Select all fields
var users []User
db.Find(&users)

// Good: Select only needed fields
var users []User
db.Select("id", "name", "email").Find(&users)
```

### 3. Use Indexes
```go
type User struct {
    gorm.Model
    Name  string `gorm:"size:100;not null;index"`        // Simple index
    Email string `gorm:"uniqueIndex;size:100;not null"`  // Unique index
    Age   int    `gorm:"index:idx_age"`                  // Named index
}
```

## 10. Testing dengan GORM

```go
// test_helper.go
func setupTestDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    
    // Auto migrate
    db.AutoMigrate(&User{}, &Post{})
    
    return db
}

// user_test.go
func TestCreateUser(t *testing.T) {
    db := setupTestDB()
    
    user := User{
        Name:  "Test User",
        Email: "test@example.com",
        Age:   25,
    }
    
    err := db.Create(&user).Error
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
    
    // Verify user was created
    var found User
    db.First(&found, user.ID)
    assert.Equal(t, "Test User", found.Name)
}
```

Ini adalah tutorial lengkap GORM untuk pemula hingga advanced. Mulailah dengan basics seperti koneksi database dan CRUD operations, lalu perlahan-lahan pelajari fitur advanced seperti associations, transactions, dan best practices.