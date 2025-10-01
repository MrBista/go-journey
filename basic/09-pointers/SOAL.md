# Soal Latihan Logic Pointer Golang

## Level 1: Dasar (Beginner)

### Soal 1: Basic Pointer Operations
Buatlah fungsi `swapValues` yang menukar nilai dua integer menggunakan pointer.

**Input:**
- Dua variabel integer a = 10, b = 20

**Expected Output:**
- Setelah pemanggilan fungsi: a = 20, b = 10

**Template:**
```go
func swapValues(a, b *int) {
    // Implementasi di sini
}
```

---

### Soal 2: Nil Pointer Checker
Buatlah fungsi `safeIncrement` yang menerima pointer integer. Jika pointer tidak nil, increment nilainya sebesar 1. Jika nil, return false.

**Signature:**
```go
func safeIncrement(p *int) bool {
    // Return true jika berhasil increment, false jika pointer nil
}
```

**Test Case:**
- Input: pointer ke 5 → Output: true, nilai menjadi 6
- Input: nil → Output: false

---

### Soal 3: Pointer to Struct
Buat struct `Student` dengan field Name (string) dan Score (int). Implementasikan fungsi `updateScore` yang menerima pointer ke Student dan nilai score baru.

**Requirements:**
- Jika pointer nil, tidak melakukan apa-apa
- Update score hanya jika score baru >= 0 dan <= 100

---

## Level 2: Menengah (Intermediate)

### Soal 4: Linked List Node
Implementasikan struktur node untuk linked list dan fungsi untuk menambah node di akhir list.

**Struktur:**
```go
type Node struct {
    Data int
    Next *Node
}
```

**Fungsi yang harus dibuat:**
- `appendNode(head **Node, data int)` - menambah node di akhir
- `printList(head *Node)` - print semua nilai dalam list

**Test Case:**
- Mulai dengan empty list (head = nil)
- Tambah nodes dengan data: 1, 2, 3
- Print hasilnya: 1 -> 2 -> 3 -> nil

---

### Soal 5: Pointer Array Manipulation
Buat fungsi yang menerima array integer dan pointer ke index. Fungsi harus mengembalikan nilai pada index tersebut, dan increment index untuk pemanggilan berikutnya.

**Signature:**
```go
func getNextValue(arr []int, index *int) (int, bool) {
    // Return nilai dan status (true jika valid, false jika index out of bounds)
}
```

**Behavior:**
- Setiap kali dipanggil, kembalikan arr[*index] dan increment *index
- Return false jika index >= len(arr)

---

### Soal 6: Reference Counter
Implementasikan struct `RefCounter` dengan pointer internal dan fungsi-fungsi untuk mengelola reference counting.

**Struktur:**
```go
type RefCounter struct {
    value *int
    count *int
}
```

**Functions:**
- `NewRefCounter(val int) *RefCounter` - constructor
- `(r *RefCounter) AddRef() *RefCounter` - increment count, return new RefCounter dengan same value
- `(r *RefCounter) Release()` - decrement count
- `(r *RefCounter) GetValue() int` - return value jika count > 0, panic jika count <= 0
- `(r *RefCounter) GetCount() int` - return current count

---

## Level 3: Lanjutan (Advanced)

### Soal 7: Circular Buffer dengan Pointers
Implementasikan circular buffer menggunakan pointers untuk head dan tail.

**Struktur:**
```go
type CircularBuffer struct {
    buffer []int
    head   *int  // pointer ke index head
    tail   *int  // pointer ke index tail
    size   *int  // pointer ke current size
    capacity int
}
```

**Methods:**
- `NewCircularBuffer(cap int) *CircularBuffer`
- `(cb *CircularBuffer) Push(val int) bool` - return false jika buffer penuh
- `(cb *CircularBuffer) Pop() (int, bool)` - return value dan status
- `(cb *CircularBuffer) IsFull() bool`
- `(cb *CircularBuffer) IsEmpty() bool`

---

### Soal 8: Pointer Chain Calculator
Buat struktur data yang menggunakan chain of pointers untuk menyimpan operasi matematika dan menghitung hasilnya.

**Struktur:**
```go
type Operation struct {
    operator rune  // '+', '-', '*', '/'
    operand  int
    next     *Operation
}

type Calculator struct {
    initial int
    ops     *Operation
}
```

**Methods:**
- `NewCalculator(initial int) *Calculator`
- `(c *Calculator) Add(val int) *Calculator` - method chaining
- `(c *Calculator) Subtract(val int) *Calculator`
- `(c *Calculator) Multiply(val int) *Calculator`
- `(c *Calculator) Divide(val int) *Calculator`
- `(c *Calculator) Calculate() int` - execute semua operasi

**Example:**
```go
result := NewCalculator(10).Add(5).Multiply(2).Subtract(3).Calculate()
// (10 + 5) * 2 - 3 = 27
```

---

### Soal 9: Memory Pool Allocator
Implementasikan simple memory pool untuk mengelola alokasi dan dealokasi pointer.

**Struktur:**
```go
type MemoryPool struct {
    pool     []*int
    free     []bool  // true jika slot kosong
    capacity int
}
```

**Methods:**
- `NewMemoryPool(capacity int) *MemoryPool`
- `(mp *MemoryPool) Allocate() *int` - return pointer ke integer, nil jika pool penuh
- `(mp *MemoryPool) Deallocate(ptr *int) bool` - return false jika pointer invalid
- `(mp *MemoryPool) GetStats() (allocated, available int)`

---

### Soal 10: Tree Traversal dengan Pointer Manipulation
Implementasikan binary tree dengan fungsi traversal yang menggunakan pointer ke pointer untuk navigasi.

**Struktur:**
```go
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}
```

**Functions:**
- `insertNode(root **TreeNode, value int)` - insert ke BST
- `findNode(root *TreeNode, value int) **TreeNode` - return pointer ke pointer node
- `deleteNode(root **TreeNode, value int) bool` - delete node, return success status
- `inorderTraversal(root *TreeNode, result *[]int)` - fill slice dengan inorder traversal

---

## Level 4: Expert

### Soal 11: Custom Smart Pointer
Implementasikan smart pointer sederhana yang automatically manage memory dan reference counting.

**Requirements:**
- Auto-cleanup when reference count reaches 0
- Thread-safe reference counting
- Weak reference support
- Custom deleter function

---

### Soal 12: Pointer-based Graph dengan Cycle Detection
Implementasikan directed graph menggunakan pointers dan algoritma untuk mendeteksi cycle.

**Struktur:**
```go
type GraphNode struct {
    Value int
    Edges []*GraphNode
    visited *bool  // untuk traversal algorithms
    inStack *bool  // untuk cycle detection
}
```

**Functions:**
- `NewGraph() *GraphNode`
- `(g *GraphNode) AddEdge(to *GraphNode)`
- `(g *GraphNode) HasCycle() bool` - menggunakan DFS
- `(g *GraphNode) FindPath(target *GraphNode) []*GraphNode`

---

## Bonus Challenge: Real-world Scenario

### Soal 13: Database Connection Pool
Implementasikan connection pool yang mengelola database connections menggunakan pointers.

**Requirements:**
- Pool size yang configurable
- Timeout untuk acquiring connection
- Health check untuk connections
- Statistics tracking
- Graceful shutdown

**Hint Structure:**
```go
type Connection struct {
    id       int
    lastUsed *time.Time
    inUse    *bool
}

type ConnectionPool struct {
    connections []*Connection
    available   chan *Connection
    maxSize     int
    timeout     time.Duration
    stats       *PoolStats
}

type PoolStats struct {
    totalConnections *int
    activeConnections *int
    waitingRequests   *int
}
```

---

## Tips Pengerjaan:

1. **Mulai dari Level 1** - pastikan memahami konsep dasar sebelum lanjut
2. **Fokus pada edge cases** - selalu handle nil pointers
3. **Testing** - buat test cases untuk setiap fungsi
4. **Memory safety** - pastikan tidak ada memory leaks
5. **Documentation** - tulis komentar untuk logic yang kompleks

## Format Submission:
- Buat file terpisah untuk setiap soal
- Sertakan test cases dan benchmark jika memungkinkan
- Gunakan `go fmt` untuk formatting
- Jalankan `go vet` untuk static analysis

Selamat mengerjakan! 🚀